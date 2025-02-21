package main

import (
	"context"
	"fmt"

	"github.com/dogesilvernet/dogesilverd/cmd/dogesilverwallet/daemon/client"
	"github.com/dogesilvernet/dogesilverd/cmd/dogesilverwallet/daemon/pb"
	"github.com/dogesilvernet/dogesilverd/cmd/dogesilverwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatDag(addressBalance.Available), utils.FormatDag(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, DAG %s %s%s\n", utils.FormatDag(response.Available), utils.FormatDag(response.Pending), pendingSuffix)

	return nil
}
