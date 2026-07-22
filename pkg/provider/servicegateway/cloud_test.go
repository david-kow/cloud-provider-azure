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
	"testing"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/cloud-provider-azure/pkg/consts"
	"sigs.k8s.io/cloud-provider-azure/pkg/provider/servicegateway/difftracker"
)

func TestLoadBalancerCloudExposesOnlyLoadBalancer(t *testing.T) {
	loadBalancer := difftracker.NewLoadBalancer(nil)
	cloud := NewLoadBalancerCloud(loadBalancer)

	got, supported := cloud.LoadBalancer()
	assert.True(t, supported)
	assert.Same(t, loadBalancer, got)
	assert.Equal(t, consts.CloudProviderName, cloud.ProviderName())
	assert.False(t, cloud.HasClusterID())

	instances, supported := cloud.Instances()
	assert.False(t, supported)
	assert.Nil(t, instances)
	instancesV2, supported := cloud.InstancesV2()
	assert.False(t, supported)
	assert.Nil(t, instancesV2)
	zones, supported := cloud.Zones()
	assert.False(t, supported)
	assert.Nil(t, zones)
	clusters, supported := cloud.Clusters()
	assert.False(t, supported)
	assert.Nil(t, clusters)
	routes, supported := cloud.Routes()
	assert.False(t, supported)
	assert.Nil(t, routes)
}

func TestLoadBalancerCloudRejectsNilLoadBalancer(t *testing.T) {
	cloud := NewLoadBalancerCloud(nil)

	loadBalancer, supported := cloud.LoadBalancer()

	assert.False(t, supported)
	assert.Nil(t, loadBalancer)
}
