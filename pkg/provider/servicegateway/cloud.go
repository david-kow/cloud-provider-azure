/*
Copyright 2026 The Kubernetes Authors.

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

package servicegateway

import (
	cloudprovider "k8s.io/cloud-provider"

	"sigs.k8s.io/cloud-provider-azure/pkg/consts"
)

type loadBalancerCloud struct {
	loadBalancer cloudprovider.LoadBalancer
}

var _ cloudprovider.Interface = (*loadBalancerCloud)(nil)

// NewLoadBalancerCloud returns the minimal cloud interface required by the Kubernetes Service controller.
func NewLoadBalancerCloud(loadBalancer cloudprovider.LoadBalancer) cloudprovider.Interface {
	return &loadBalancerCloud{loadBalancer: loadBalancer}
}

func (*loadBalancerCloud) Initialize(cloudprovider.ControllerClientBuilder, <-chan struct{}) {}

func (c *loadBalancerCloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return c.loadBalancer, c.loadBalancer != nil
}

func (*loadBalancerCloud) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

func (*loadBalancerCloud) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return nil, false
}

func (*loadBalancerCloud) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

func (*loadBalancerCloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

func (*loadBalancerCloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

func (*loadBalancerCloud) ProviderName() string {
	return consts.CloudProviderName
}

func (*loadBalancerCloud) HasClusterID() bool {
	return false
}
