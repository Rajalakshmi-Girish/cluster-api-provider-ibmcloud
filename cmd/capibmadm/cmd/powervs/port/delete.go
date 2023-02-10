/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package port

import (
	"context"

	"github.com/spf13/cobra"

	powerClient "github.com/IBM-Cloud/power-go-client/clients/instance"

	logf "sigs.k8s.io/cluster-api/cmd/clusterctl/log"

	"sigs.k8s.io/cluster-api-provider-ibmcloud/cmd/capibmadm/clients/iam"
	"sigs.k8s.io/cluster-api-provider-ibmcloud/cmd/capibmadm/clients/powervs"
	"sigs.k8s.io/cluster-api-provider-ibmcloud/cmd/capibmadm/options"
	"sigs.k8s.io/cluster-api-provider-ibmcloud/cmd/capibmadm/utils"
)

type portDeleteOptions struct {
	network string
	portID  string
}

// DeleteCommand function to delete network's port.
func DeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete PowerVS network port",
		Example: `
# Delete PowerVS network port with ID <port-id> in network "capi-network"
export IBMCLOUD_API_KEY=<api-key>
capibmadm powervs port delete --port-id <port-id> --network <network-name/network-id> --service-instance-id <service-instance-id> --zone <zone>`,
	}

	var portDeleteOption portDeleteOptions
	cmd.Flags().StringVar(&portDeleteOption.portID, "port-id", "", "Port ID to be deleted")
	cmd.Flags().StringVar(&portDeleteOption.network, "network", "", "Network ID or Name(preference will be given to the ID over Name")
	_ = cmd.MarkFlagRequired("port-id")
	_ = cmd.MarkFlagRequired("network")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := deletePort(cmd.Context(), portDeleteOption); err != nil {
			return err
		}
		return nil
	}
	return cmd
}

func deletePort(ctx context.Context, portDeleteOption portDeleteOptions) error {
	log := logf.Log
	log.Info("Deleting PowerVS network port", "network", portDeleteOption.network, "service-instance-id", options.GlobalOptions.ServiceInstanceID, "port-id", portDeleteOption.portID)
	auth := iam.GetIAMAuth()
	accountID, err := utils.GetAccountID(ctx, auth)
	if err != nil {
		return err
	}
	sess, err := powervs.NewPISession(accountID, options.GlobalOptions.PowerVSZone, options.GlobalOptions.Debug)
	if err != nil {
		return err
	}
	networkClient := powerClient.NewIBMPINetworkClient(ctx, sess, options.GlobalOptions.ServiceInstanceID)
	if err := networkClient.DeletePort(portDeleteOption.network, portDeleteOption.portID); err != nil {
		return err
	}
	log.Info("Successfully deleted a port", "port-id", portDeleteOption.portID)
	return nil
}
