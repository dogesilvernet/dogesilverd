package protowire

import (
	"github.com/dogesilvernet/dogesilverd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DogesilverdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DogesilverdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *DogesilverdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
