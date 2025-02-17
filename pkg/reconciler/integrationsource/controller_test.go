/*
Copyright 2024 The Knative Authors

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

package integrationsource

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	filteredFactory "knative.dev/pkg/client/injection/kube/informers/factory/filtered"
	"knative.dev/pkg/configmap"
	. "knative.dev/pkg/reconciler/testing"

	// Fake injection informers
	"knative.dev/eventing/pkg/apis/feature"
	_ "knative.dev/eventing/pkg/client/injection/informers/sources/v1/containersource/fake"
	_ "knative.dev/eventing/pkg/client/injection/informers/sources/v1alpha1/integrationsource/fake"
	"knative.dev/eventing/pkg/eventingtls"
	_ "knative.dev/pkg/client/injection/kube/informers/apps/v1/deployment/fake"
	_ "knative.dev/pkg/client/injection/kube/informers/core/v1/configmap/filtered/fake"
	_ "knative.dev/pkg/client/injection/kube/informers/core/v1/serviceaccount/filtered/fake"
	_ "knative.dev/pkg/client/injection/kube/informers/factory/filtered/fake"
	_ "knative.dev/pkg/injection/clients/dynamicclient/fake"
)

func TestNew(t *testing.T) {
	ctx, _ := SetupFakeContext(t, SetUpInformerSelector)

	c := NewController(ctx, configmap.NewStaticWatcher(
		&corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: feature.FlagsConfigName,
			},
		},
	))

	if c == nil {
		t.Fatal("Expected NewController to return a non-nil value")
	}
}

func SetUpInformerSelector(ctx context.Context) context.Context {
	ctx = filteredFactory.WithSelectors(ctx, eventingtls.TrustBundleLabelSelector)
	return ctx
}
