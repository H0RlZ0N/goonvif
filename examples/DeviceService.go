package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/H0RlZ0N/onvif"
	"github.com/H0RlZ0N/onvif/device"
	devtypes "github.com/H0RlZ0N/onvif/device"
	"github.com/H0RlZ0N/onvif/xsd/onvif"
)

const (
	login    = "test"
	password = "besovideo215"
)

func main() {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.13.14:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	devlist, err := onvif.GetAvailableDevicesAtSpecificEthernetInterface("以太网 2")
	if err != nil {
		log.Println(err)
	}
	log.Printf("------> devlist num: %v", len(devlist))
	for _, dev := range devlist {
		log.Printf("------> DeviceInfo: %v", dev)
	}

	//Preparing commands
	systemDateAndTyme := device.GetSystemDateAndTime{}
	getCapabilities := device.GetCapabilities{Category: "All"}

	//Commands execution
	systemDateAndTymeResponse, err := devtypes.Call_GetSystemDateAndTime(ctx, dev, systemDateAndTyme)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(systemDateAndTymeResponse)
	}
	getCapabilitiesResponse, err := devtypes.Call_GetCapabilities(ctx, dev, getCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}

}
