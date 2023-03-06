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

// Call_GetMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetadataConfigurationResponse.
func Call_GetMetadataConfiguration(ctx context.Context, dev *goonvif.Device, request media.GetMetadataConfiguration) (media.GetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationResponse media.GetMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetMetadataConfiguration")
		return reply.Body.GetMetadataConfigurationResponse, errors.Annotate(err, "reply")
	}
}
