## PowerVS Network Commands

### 1. capibmadm powervs port delete

#### Usage:
Delete PowerVS network port.

#### Environmental Variable:
IBMCLOUD_API_KEY: IBM Cloud API key.

#### Arguments:
--service-instance-id: PowerVS service instance id.

--zone: PowerVS zone.

--port-id: ID of network port.

--network: Network ID or Name(preference will be given to the ID over Name).

#### Example:
```shell
export IBMCLOUD_API_KEY=<api-key>
capibmadm powervs port delete --port-id <port-id> --network <network-name/network-id> --service-instance-id <service-instance-id> --zone <zone>
```