package rpchandlers

import (
	"github.com/dogesilvernet/dogesilverd/app/appmessage"
	"github.com/dogesilvernet/dogesilverd/app/rpc/rpccontext"
	"github.com/dogesilvernet/dogesilverd/infrastructure/network/netadapter/router"
)

// HandleGetBlockCount handles the respectively named RPC command
func HandleGetBlockCount(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	syncInfo, err := context.Domain.Consensus().GetSyncInfo()
	if err != nil {
		return nil, err
	}
	response := appmessage.NewGetBlockCountResponseMessage(syncInfo)
	return response, nil
}
