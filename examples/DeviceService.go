package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/H0RlZ0N/goonvif"
	"github.com/H0RlZ0N/goonvif/device"
	"github.com/H0RlZ0N/goonvif/media"
	devsdk "github.com/H0RlZ0N/goonvif/sdk/device"
	medsdk "github.com/H0RlZ0N/goonvif/sdk/media"
)

const (
	login    = "admin"
	password = "besovideo235"
)

func main() {
	ctx := context.Background()

	netinfos, _ := net.Interfaces()
	for _, netinfo := range netinfos {
		if strings.Contains(netinfo.Flags.String(), "up") && !strings.Contains(netinfo.Flags.String(), "loopback") && !strings.Contains(netinfo.Name, "VMware") {
			log.Printf("------> net Name: %v Flags: %v", netinfo.Name, netinfo.Flags)
			devlist, err := goonvif.GetAvailableDevicesAtSpecificEthernetInterface(netinfo.Name)
			if err != nil {
				log.Println(err)
			}
			log.Printf("------> devlist num: %v", len(devlist))
			for _, dev := range devlist {
				log.Printf("------> DeviceInfo: %v", dev.GetDeviceParams().Xaddr)
			}
		}
	}

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.6.235:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	//Preparing commands
	systemDateAndTyme := device.GetSystemDateAndTime{}
	getCapabilities := device.GetCapabilities{Category: "All"}
	getDeviceInfo := device.GetDeviceInformation{}
	getProfiles := media.GetProfiles{}
	getStreamUri := media.GetStreamUri{}

	//Commands execution
	systemDateAndTymeResponse, err := devsdk.Call_GetSystemDateAndTime(ctx, dev, systemDateAndTyme)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("---> SystemDateAndTime DateTimeType: %v\n", systemDateAndTymeResponse.SystemDateAndTime.DateTimeType)
		log.Printf("---> SystemDateAndTime TimeZone: %v\n", systemDateAndTymeResponse.SystemDateAndTime.TimeZone)
	}
	getCapabilitiesResponse, err := devsdk.Call_GetCapabilities(ctx, dev, getCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("---> Capabilities Analytics: %v\n", getCapabilitiesResponse.Capabilities.Analytics.XAddr)
		log.Printf("---> Capabilities Device: %v\n", getCapabilitiesResponse.Capabilities.Device.XAddr)
		log.Printf("---> Capabilities Events: %v\n", getCapabilitiesResponse.Capabilities.Events.XAddr)
		log.Printf("---> Capabilities Imaging: %v\n", getCapabilitiesResponse.Capabilities.Imaging.XAddr)
		log.Printf("---> Capabilities Media: %v\n", getCapabilitiesResponse.Capabilities.Media.XAddr)
		log.Printf("---> Capabilities DeviceIO: %v\n", getCapabilitiesResponse.Capabilities.Extension.DeviceIO.XAddr)

	}

	getDeviceInfoResponse, err := devsdk.Call_GetDeviceInformation(ctx, dev, getDeviceInfo)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("---> DeviceInfo FirmwareVersion: %v\n", getDeviceInfoResponse.FirmwareVersion)
		log.Printf("---> DeviceInfo Manufacturer: %v\n", getDeviceInfoResponse.Manufacturer)
		log.Printf("---> DeviceInfo Model: %v\n", getDeviceInfoResponse.Model)
		log.Printf("---> DeviceInfo SerialNumber: %v\n", getDeviceInfoResponse.SerialNumber)
	}

	getProfilesResponse, err := medsdk.Call_GetProfiles(ctx, dev, getProfiles)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("---> Profiles num: %v\n", len(getProfilesResponse.Profiles))
		for _, Profile := range getProfilesResponse.Profiles {
			log.Printf("---> Profile Token: %v\n", Profile.Token)
			log.Printf("---> Profile Name: %v\n", Profile.Name)
			log.Printf("---> Profile VideoEncoder Encoding: %v\n", Profile.VideoEncoderConfiguration.Encoding)
			log.Printf("---> Profile VideoEncoder Height: %v\n", Profile.VideoEncoderConfiguration.Resolution.Height)
			log.Printf("---> Profile VideoEncoder Width: %v\n", Profile.VideoEncoderConfiguration.Resolution.Width)
			log.Printf("---> Profile AudioEncoder Encoding: %v\n", Profile.AudioEncoderConfiguration.Encoding)
			log.Printf("---> Profile AudioEncoder Bitrate: %v\n", Profile.AudioEncoderConfiguration.Bitrate)
			log.Printf("---> Profile AudioEncoder SampleRate: %v\n", Profile.AudioEncoderConfiguration.SampleRate)
			getStreamUri.ProfileToken = Profile.Token
			getStreamUriResponse, err := medsdk.Call_GetStreamUri(ctx, dev, getStreamUri)
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("---> Profile %v rtsp url: %v\n", getStreamUri.ProfileToken, getStreamUriResponse.MediaUri.Uri)
			}
			//break
		}
	}

}
