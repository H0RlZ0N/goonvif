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

// Call_StopMulticastStreaming forwards the call to dev.CallMethod() then parses the payload of the reply as a StopMulticastStreamingResponse.
func Call_StopMulticastStreaming(ctx context.Context, dev *goonvif.Device, request media.StopMulticastStreaming) (media.StopMulticastStreamingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopMulticastStreamingResponse media.StopMulticastStreamingResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StopMulticastStreamingResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "StopMulticastStreaming")
		return reply.Body.StopMulticastStreamingResponse, errors.Annotate(err, "reply")
	}
}
