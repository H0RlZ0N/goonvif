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

// Call_CreateUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateUsersResponse.
func Call_CreateUsers(ctx context.Context, dev *goonvif.Device, request device.CreateUsers) (device.CreateUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateUsersResponse device.CreateUsersResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateUsersResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateUsers")
		return reply.Body.CreateUsersResponse, errors.Annotate(err, "reply")
	}
}
