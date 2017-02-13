// Copyright 2017 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"net"

	cnitypes "github.com/containernetworking/cni/pkg/types"
)

type route struct {
	Dst cnitypes.IPNet `json:"dst"`
	GW  net.IP         `json:"gw,omitempty"`
}

type cniNetConfig struct {
	Subnet  cnitypes.IPNet `json:"subnet,omitempty"`
	Gateway net.IP         `json:"gateway,omitempty"`
	Routes  []route        `json:"routes,omitempty"`
}

// Configuration for the host agent
type hostAgentConfig struct {
	// Log level
	LogLevel string `json:"log-level,omitempty"`

	// Absolute path to a kubeconfig file
	KubeConfig string `json:"kubeconfig,omitempty"`

	// Name of Kubernetes node on which this agent is running
	NodeName string `json:"node-name,omitempty"`

	// TCP port to run status server on (or 0 to disable)
	StatusPort int `json:"status-port,omitempty"`

	// Directory containing OpFlex CNI metadata
	CniMetadataDir string `json:"cni-metadata-dir,omitempty"`

	// Name of the CNI network
	CniNetwork string `json:"cni-network,omitempty"`

	// Directory for writing OpFlex endpoint metadata
	OpFlexEndpointDir string `json:"opflex-endpoint-dir,omitempty"`

	// Directory for writing OpFlex service metadata
	OpFlexServiceDir string `json:"opflex-service-dir,omitempty"`

	// Location of the OVS DB socket
	OvsDbSock string `json:"ovs-db-sock,omitempty"`

	// Location of the endpoint RPC socket used for communicating with
	// the CNI plugin
	EpRpcSock string `json:"ep-rpc-sock,omitempty"`

	// Name of the OVS integration bridge
	IntBridgeName string `json:"int-bridge-name,omitempty"`

	// Name of the OVS access bridge
	AccessBridgeName string `json:"access-bridge-name,omitempty"`

	// Interface for external service traffic
	ServiceIface string `json:"service-iface,omitempty"`

	// VLAN for service interface traffic
	ServiceIfaceVlan uint `json:"service-iface-vlan,omitempty"`

	// Interface MTU to use when configuring container interfaces
	InterfaceMtu int `json:"interface-mtu,omitempty"`

	// Configuration for CNI networks
	NetConfig []cniNetConfig `json:"cni-netconfig,omitempty"`

	// ACI VRF for this kubernetes instance
	AciVrf string `json:"aci-vrf,omitempty"`

	// ACI Tenant containing the ACI VRF for this kubernetes instance
	AciVrfTenant string `json:"aci-vrf-tenant,omitempty"`
}

func (config *hostAgentConfig) initFlags() {
	flag.StringVar(&config.LogLevel, "log-level", "info", "Log level")

	flag.StringVar(&config.KubeConfig, "kubeconfig", "", "Absolute path to a kubeconfig file")
	flag.StringVar(&config.NodeName, "node-name", "", "Name of Kubernetes node on which this agent is running")

	flag.IntVar(&config.StatusPort, "status-port", 8090, "TCP port to run status server on (or 0 to disable)")

	flag.StringVar(&config.CniMetadataDir, "cni-metadata-dir", "/usr/local/var/lib/aci-containers/", "Directory for writing OpFlex endpoint metadata")
	flag.StringVar(&config.CniNetwork, "cni-network", "k8s-pod-network", "Name of the CNI network")

	flag.StringVar(&config.OpFlexEndpointDir, "opflex-endpoint-dir", "/usr/local/var/lib/opflex-agent-ovs/endpoints/", "Directory for writing OpFlex endpoint metadata")
	flag.StringVar(&config.OpFlexServiceDir, "opflex-service-dir", "/usr/local/var/lib/opflex-agent-ovs/services/", "Directory for writing OpFlex anycast service metadata")

	flag.StringVar(&config.OvsDbSock, "ovs-db-sock", "/usr/local/var/run/openvswitch/db.sock", "Location of the OVS DB socket")
	flag.StringVar(&config.EpRpcSock, "ep-rpc-sock", "/usr/local/var/run/aci-containers-ep-rpc.sock", "Location of the endpoint RPC socket used for communicating with the CNI plugin")

	flag.StringVar(&config.IntBridgeName, "int-bridge-name", "br-int", "Name of the OVS integration bridge")
	flag.StringVar(&config.AccessBridgeName, "access-bridge-name", "br-access", "Name of the OVS access bridge")

	flag.IntVar(&config.InterfaceMtu, "interface-mtu", 1500, "Interface MTU to use when configuring container interfaces")

	flag.StringVar(&config.ServiceIface, "service-iface", "eth2", "Interface for external service traffic")
	flag.UintVar(&config.ServiceIfaceVlan, "service-iface-vlan", 4003, "VLAN for service interface traffic")

	flag.StringVar(&config.AciVrf, "aci-vrf", "kubernetes-vrf", "ACI VRF for this kubernetes instance")
	flag.StringVar(&config.AciVrfTenant, "aci-vrf-tenant", "common", "ACI Tenant containing the ACI VRF for this kubernetes instance")
}
