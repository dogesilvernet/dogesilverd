package protowire

import (
	"github.com/dogesilvernet/dogesilverd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DogesilverdMessage_DoneHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DogesilverdMessage_DoneHeaders is nil")
	}
	return &appmessage.MsgDoneHeaders{}, nil
}

func (x *DogesilverdMessage_DoneHeaders) fromAppMessage(_ *appmessage.MsgDoneHeaders) error {
	return nil
}
