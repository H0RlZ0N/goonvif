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

// Call_SetVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoAnalyticsConfigurationResponse.
func Call_SetVideoAnalyticsConfiguration(ctx context.Context, dev *goonvif.Device, request media.SetVideoAnalyticsConfiguration) (media.SetVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoAnalyticsConfigurationResponse media.SetVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoAnalyticsConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetVideoAnalyticsConfiguration")
		return reply.Body.SetVideoAnalyticsConfigurationResponse, errors.Annotate(err, "reply")
	}
}
