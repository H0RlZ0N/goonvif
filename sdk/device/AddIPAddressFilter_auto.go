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

// Call_AddIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a AddIPAddressFilterResponse.
func Call_AddIPAddressFilter(ctx context.Context, dev *goonvif.Device, request device.AddIPAddressFilter) (device.AddIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddIPAddressFilterResponse device.AddIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddIPAddressFilterResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddIPAddressFilter")
		return reply.Body.AddIPAddressFilterResponse, errors.Annotate(err, "reply")
	}
}
