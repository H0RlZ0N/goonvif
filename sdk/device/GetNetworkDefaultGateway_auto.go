// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/H0RlZ0N/goonvif"
	"github.com/H0RlZ0N/goonvif/sdk"
	"github.com/H0RlZ0N/goonvif/device"
)

// Call_GetNetworkDefaultGateway forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkDefaultGatewayResponse.
func Call_GetNetworkDefaultGateway(ctx context.Context, dev *goonvif.Device, request device.GetNetworkDefaultGateway) (device.GetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkDefaultGatewayResponse device.GetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetNetworkDefaultGateway")
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Annotate(err, "reply")
	}
}
