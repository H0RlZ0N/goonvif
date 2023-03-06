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

// Call_LoadCertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a LoadCertificatesResponse.
func Call_LoadCertificates(ctx context.Context, dev *goonvif.Device, request device.LoadCertificates) (device.LoadCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCertificatesResponse device.LoadCertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.LoadCertificatesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "LoadCertificates")
		return reply.Body.LoadCertificatesResponse, errors.Annotate(err, "reply")
	}
}
