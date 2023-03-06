// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/H0RlZ0N/goonvif"
	"github.com/H0RlZ0N/goonvif/sdk"
	"github.com/H0RlZ0N/goonvif/ptz"
)

// Call_ContinuousMove forwards the call to dev.CallMethod() then parses the payload of the reply as a ContinuousMoveResponse.
func Call_ContinuousMove(ctx context.Context, dev *goonvif.Device, request ptz.ContinuousMove) (ptz.ContinuousMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ContinuousMoveResponse ptz.ContinuousMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ContinuousMoveResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "ContinuousMove")
		return reply.Body.ContinuousMoveResponse, errors.Annotate(err, "reply")
	}
}
