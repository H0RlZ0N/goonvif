// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/H0RlZ0N/goonvif"
	"github.com/H0RlZ0N/goonvif/sdk"
	"github.com/H0RlZ0N/goonvif/media"
)

// Call_GetOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDResponse.
func Call_GetOSD(ctx context.Context, dev *goonvif.Device, request media.GetOSD) (media.GetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDResponse media.GetOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetOSDResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetOSD")
		return reply.Body.GetOSDResponse, errors.Annotate(err, "reply")
	}
}
