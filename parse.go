// Copyright 2016 Gravitational Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rigging

import (
	"io"

	"github.com/gravitational/trace"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
	batchv1 "k8s.io/client-go/pkg/apis/batch/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/pkg/apis/rbac/v1alpha1"
)

type ResourceHeader struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// ParseResourceHeader parses resource header information
func ParseResourceHeader(reader io.Reader) (*ResourceHeader, error) {
	var out ResourceHeader
	err := yaml.NewYAMLOrJSONDecoder(reader, DefaultBufferSize).Decode(&out)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &out, nil
}

// ParseDaemonSet parses daemon set from reader
func ParseDaemonSet(r io.Reader) (*v1beta1.DaemonSet, error) {
	if r == nil {
		return nil, trace.BadParameter("missing reader")
	}
	ds := v1beta1.DaemonSet{}
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&ds)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &ds, nil
}

// ParseJob parses the specified reader as a Job resource
func ParseJob(r io.Reader) (*batchv1.Job, error) {
	if r == nil {
		return nil, trace.BadParameter("missing reader")
	}

	var job batchv1.Job
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&job)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &job, nil
}

// ParseSerializedReference parses serialized reference object
// used in annotations
func ParseSerializedReference(r io.Reader) (*api.SerializedReference, error) {
	ref := api.SerializedReference{}
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&ref)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &ref, nil
}

// ParseReplicationController parses replication controller
func ParseReplicationController(r io.Reader) (*v1.ReplicationController, error) {
	if r == nil {
		return nil, trace.BadParameter("missing reader")
	}
	rc := v1.ReplicationController{}
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&rc)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &rc, nil
}

// ParseDeployment parses deployment
func ParseDeployment(r io.Reader) (*v1beta1.Deployment, error) {
	if r == nil {
		return nil, trace.BadParameter("missing reader")
	}
	rc := v1beta1.Deployment{}
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&rc)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &rc, nil
}

// ParseService parses service
func ParseService(r io.Reader) (*v1.Service, error) {
	if r == nil {
		return nil, trace.BadParameter("missing reader")
	}
	svc := v1.Service{}
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&svc)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &svc, nil
}

// ParseConfigMap parses Config Map
func ParseConfigMap(r io.Reader) (*v1.ConfigMap, error) {
	if r == nil {
		return nil, trace.BadParameter("missing reader")
	}
	cm := v1.ConfigMap{}
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&cm)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &cm, nil
}

// ParseSecret parses a secret from the specified stream
func ParseSecret(r io.Reader) (*v1.Secret, error) {
	if r == nil {
		return nil, trace.BadParameter("missing reader")
	}
	secret := v1.Secret{}
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&secret)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &secret, nil
}

// ParseServiceAccount parses a service account from the specified stream
func ParseServiceAccount(r io.Reader) (*v1.ServiceAccount, error) {
	var account v1.ServiceAccount
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&account)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &account, nil
}

// ParseRole parses an rbac role from the specified stream
func ParseRole(r io.Reader) (*v1alpha1.Role, error) {
	var role v1alpha1.Role
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&role)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &role, nil
}

// ParseClusterRole parses an rbac cluster role from the specified stream
func ParseClusterRole(r io.Reader) (*v1alpha1.ClusterRole, error) {
	var role v1alpha1.ClusterRole
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&role)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &role, nil
}

// ParseRoleBinding parses an rbac role binding from the specified stream
func ParseRoleBinding(r io.Reader) (*v1alpha1.RoleBinding, error) {
	var binding v1alpha1.RoleBinding
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&binding)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &binding, nil
}

// ParseClusterRoleBinding parses an rbac cluster role binding from the specified stream
func ParseClusterRoleBinding(r io.Reader) (*v1alpha1.ClusterRoleBinding, error) {
	var binding v1alpha1.ClusterRoleBinding
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&binding)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &binding, nil
}

// ParsePodSecurityPolicy parses a pod security policy from the specified stream
func ParsePodSecurityPolicy(r io.Reader) (*v1beta1.PodSecurityPolicy, error) {
	var policy v1beta1.PodSecurityPolicy
	err := yaml.NewYAMLOrJSONDecoder(r, DefaultBufferSize).Decode(&policy)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &policy, nil
}
