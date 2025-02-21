package protowire

import (
	"github.com/dogesilvernet/dogesilverd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DogesilverdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DogesilverdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *DogesilverdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
