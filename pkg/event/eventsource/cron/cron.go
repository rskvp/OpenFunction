/*
Copyright 2022 The OpenFunction Authors.

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

package cron

import (
	"sync"

	componentsv1alpha1 "github.com/dapr/dapr/pkg/apis/components/v1alpha1"
	"github.com/go-logr/logr"
	kedav1alpha1 "github.com/kedacore/keda/v2/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	ofcore "github.com/rskvp/openfunction/apis/core/v1beta1"
	ofevent "github.com/rskvp/openfunction/apis/events/v1alpha1"
	"github.com/rskvp/openfunction/pkg/event"
)

const (
	ComponentType    = "bindings.cron"
	ComponentVersion = "v1"
)

type EventSource struct {
	mu       sync.Mutex
	log      logr.Logger
	Spec     *ofevent.CronSpec
	Metadata map[string]interface{}
}

func NewCronEventSource(log logr.Logger, spec *ofevent.CronSpec) event.OpenFunctionEventSource {
	es := &EventSource{}

	es.log = log
	es.log.WithName("CronEventSource")

	es.Spec = spec
	es.init()
	return es
}

func (es *EventSource) init() {
	m := map[string]interface{}{}

	// handle mandatory parameters
	m["schedule"] = es.Spec.Schedule

	es.Metadata = m
}

func (es *EventSource) SetMetadata(key string, value interface{}) {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.Metadata[key] = value
}

func (es *EventSource) GetMetadata() map[string]interface{} {
	es.mu.Lock()
	defer es.mu.Unlock()
	return es.Metadata
}

func (es *EventSource) GenComponent(namespace string, name string) (*componentsv1alpha1.Component, error) {
	component := &componentsv1alpha1.Component{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	component.Spec.Type = ComponentType
	component.Spec.Version = ComponentVersion

	metadataItems, err := event.ConvertMetadata(es.GetMetadata())
	if err != nil {
		es.log.Error(err, "failed to generate component", "namespace", namespace, "name", name)
		return nil, err
	}
	component.Spec.Metadata = metadataItems
	return component, nil
}

func (es *EventSource) GenScaleOptions() (*ofcore.KedaScaledObject, *kedav1alpha1.ScaleTriggers) {
	return nil, nil
}
