/*
Copyright © 2021 SonicWall, Inc.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SonicwallNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SonicwallNetworkPolicySpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SonicwallNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SonicwallNetworkPolicy `json:"items"`
}

type SonicwallNetworkPolicySpec []SonicwallNetworkPolicySpecItem

type SonicwallNetworkPolicySpecItem struct {
	EndpointSelector SonicwallNetworkPolicyEndpointSelector `json:"endpointSelector"`
	Ingress          []SonicwallNetworkPolicyIngressItem    `json:"ingress,omitempty"`
	Egress           []SonicwallNetworkPolicyEgressItem     `json:"egress,omitempty"`
}

type SonicwallNetworkPolicyEndpointSelector struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

type SonicwallNetworkPolicyIngressItem struct {
	FromEndpoints []SonicwallNetworkPolicyEndpointSelector `json:"fromEndpoints,omitempty"`
	FromCIDR      []string                                 `json:"fromCIDR,omitempty"`
	FromFQDN      []string                                 `json:"fromFQDN,omitempty"`
	FromPorts     []SonicwallNetworkPolicyPorts            `json:"fromPorts,omitempty"`
}

type SonicwallNetworkPolicyEgressItem struct {
	ToEndpoints []SonicwallNetworkPolicyEndpointSelector `json:"toEndpoints,omitempty"`
	ToCIDR      []string                                 `json:"toCIDR,omitempty"`
	ToFQDN      []string                                 `json:"toFQDN,omitempty"`
	ToPorts     []SonicwallNetworkPolicyPorts            `json:"toPorts,omitempty"`
}

type SonicwallNetworkPolicyPorts struct {
	Ports []SonicwallNetworkPolicyPort `json:"ports,omitempty"`
}

type SonicwallNetworkPolicyPort struct {
	Port     string `json:"port,omitempty" validate:"numeric"`
	Protocol string `json:"protocol,omitempty" validate:"oneof=TCP UDP ALL"`
}
