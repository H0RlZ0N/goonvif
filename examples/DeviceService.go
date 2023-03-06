package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/H0RlZ0N/goonvif"
	"github.com/H0RlZ0N/goonvif/device"
	devsdk "github.com/H0RlZ0N/goonvif/sdk/device"
)

const (
	login    = "test"
	password = "besovideo215"
)

func main() {
	ctx := context.Background()

	devlist, err := goonvif.GetAvailableDevicesAtSpecificEthernetInterface("以太网 2")
	if err != nil {
		log.Println(err)
	}
	log.Printf("------> devlist num: %v", len(devlist))
	for _, dev := range devlist {
		log.Printf("------> DeviceInfo: %v", dev.GetDeviceParams().Xaddr)
	}

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.6.215:80",
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

	//Commands execution
	systemDateAndTymeResponse, err := devsdk.Call_GetSystemDateAndTime(ctx, dev, systemDateAndTyme)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(systemDateAndTymeResponse)
	}
	getCapabilitiesResponse, err := devsdk.Call_GetCapabilities(ctx, dev, getCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}

}
