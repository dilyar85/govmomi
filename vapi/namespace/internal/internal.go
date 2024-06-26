/*
Copyright (c) 2020-2024 VMware, Inc. All Rights Reserved.

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

package internal

const (
	// NamespaceClusterPath is the rest endpoint for the namespace cluster management API
	NamespaceClusterPath                    = "/api/vcenter/namespace-management/clusters"
	NamespaceDistributedSwitchCompatibility = "/api/vcenter/namespace-management/distributed-switch-compatibility"
	NamespaceEdgeClusterCompatibility       = "/api/vcenter/namespace-management/edge-cluster-compatibility"
	SupervisorServicesPath                  = "/api/vcenter/namespace-management/supervisor-services"

	NamespacesPath = "/api/vcenter/namespaces/instances"
	VmClassesPath  = "/api/vcenter/namespace-management/virtual-machine-classes"
)

type SupportBundleToken struct {
	Value string `json:"wcp-support-bundle-token"`
}
