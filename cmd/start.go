package cmd

import (
	"errors"
	"fmt"
	"github.com/NodeFactoryIo/vedran/internal/loadbalancer"
	"github.com/NodeFactoryIo/vedran/pkg/util/random"
	"github.com/spf13/cobra"
)

var (
	authSecret string
	name       string
	capacity   int64
	whitelist  []string
	fee        float32
	selection  string
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts vedran load balancer",
	Run:   startCommand,
	Args: func(cmd *cobra.Command, args []string) error {
		// validate flags values
		if selection != "round-robin" && selection != "random" {
			return errors.New("invalid selection option selected")
		}
		return nil
	},
}

func init() {
	startCmd.Flags().StringVar(
		&authSecret,
		"auth-secret",
		"",
		"[REQUIRED] authentication secret used for generating tokens")

	startCmd.Flags().StringVar(
		&name,
		"name",
		fmt.Sprintf("load-balancer-%s", random.String(12, random.Alphabetic)),
		"[OPTIONAL] public name for load balancer, autogenerated name used if omitted")

	startCmd.Flags().Int64Var(
		&capacity,
		"capacity",
		10,
		"[OPTIONAL] maximum number of nodes allowed to connect")

	startCmd.Flags().StringArrayVar(
		&whitelist,
		"whitelist",
		nil,
		"[OPTIONAL] comma separated list of node id-s, if provided only these nodes will be allowed to connect")

	startCmd.Flags().Float32Var(
		&fee,
		"fee",
		0,
		"[REQUIRED] float value representing fee percentage")

	startCmd.Flags().StringVar(
		&selection,
		"selection",
		"round-robin",
		"[OPTIONAL] type of selection used for choosing nodes")

	// mark required flags
	_ = startCmd.MarkFlagRequired("fee")

	RootCmd.AddCommand(startCmd)
}

func startCommand(_ *cobra.Command, _ []string) {
	loadbalancer.StartLoadBalancerServer(loadbalancer.Properties{
		AuthSecret: authSecret,
		Name:       name,
		Capacity:   capacity,
		Whitelist:  whitelist,
		Fee:        fee,
		Selection:  selection,
	}, "4000")
}
