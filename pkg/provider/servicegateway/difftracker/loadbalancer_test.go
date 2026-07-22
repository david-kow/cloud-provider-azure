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

package difftracker

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/cloud-provider-azure/pkg/consts"
)

func TestLoadBalancerPublishesValidationWarningEvent(t *testing.T) {
	recorder := record.NewFakeRecorder(1)
	tracker := &DiffTracker{eventRecorder: recorder}
	loadBalancer := NewLoadBalancer(tracker)
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "service",
			UID:       types.UID("service-uid"),
			Annotations: map[string]string{
				consts.ServiceAnnotationLoadBalancerInternal: consts.TrueAnnotationValue,
			},
		},
	}

	status, err := loadBalancer.EnsureLoadBalancer(context.Background(), "cluster", service, nil)

	assert.Nil(t, status)
	assert.Error(t, err)
	select {
	case event := <-recorder.Events:
		assert.Contains(t, event, "Warning UnsupportedInternalLoadBalancer")
		assert.Contains(t, event, consts.ServiceAnnotationLoadBalancerInternal)
	case <-time.After(time.Second):
		t.Fatal("expected a ServiceGateway warning event")
	}
}

func TestLoadBalancerRequiresTrackerInitialization(t *testing.T) {
	loadBalancer := NewLoadBalancer(nil)
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "service",
			UID:       types.UID("service-uid"),
		},
	}

	status, err := loadBalancer.EnsureLoadBalancer(context.Background(), "cluster", service, nil)
	assert.EqualError(t, err, "ServiceGateway LoadBalancer is not initialized")
	assert.Nil(t, status)
	assert.EqualError(t, loadBalancer.UpdateLoadBalancer(context.Background(), "cluster", service, nil), "ServiceGateway LoadBalancer is not initialized")
	assert.EqualError(t, loadBalancer.EnsureLoadBalancerDeleted(context.Background(), "cluster", service), "ServiceGateway LoadBalancer is not initialized")
	_, _, err = loadBalancer.GetLoadBalancer(context.Background(), "cluster", service)
	assert.EqualError(t, err, "ServiceGateway LoadBalancer is not initialized")
}
