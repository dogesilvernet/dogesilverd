package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dogesilvernet/dogesilverd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.DogesilverdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.DogesilverdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.DogesilverdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.DogesilverdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.DogesilverdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.DogesilverdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.DogesilverdMessage_BanRequest{}),
	reflect.TypeOf(protowire.DogesilverdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
