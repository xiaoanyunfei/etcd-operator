/*
Copyright 2023 The etcd-operator Authors

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

// ClusterSpecApplyConfiguration represents an declarative configuration of the ClusterSpec type for use
// with apply.
type ClusterSpecApplyConfiguration struct {
	Size       *int                         `json:"size,omitempty"`
	Repository *string                      `json:"repository,omitempty"`
	Version    *string                      `json:"version,omitempty"`
	Paused     *bool                        `json:"paused,omitempty"`
	Pod        *PodPolicyApplyConfiguration `json:"pod,omitempty"`
	TLS        *TLSPolicyApplyConfiguration `json:"TLS,omitempty"`
}

// ClusterSpecApplyConfiguration constructs an declarative configuration of the ClusterSpec type for use with
// apply.
func ClusterSpec() *ClusterSpecApplyConfiguration {
	return &ClusterSpecApplyConfiguration{}
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *ClusterSpecApplyConfiguration) WithSize(value int) *ClusterSpecApplyConfiguration {
	b.Size = &value
	return b
}

// WithRepository sets the Repository field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Repository field is set to the value of the last call.
func (b *ClusterSpecApplyConfiguration) WithRepository(value string) *ClusterSpecApplyConfiguration {
	b.Repository = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ClusterSpecApplyConfiguration) WithVersion(value string) *ClusterSpecApplyConfiguration {
	b.Version = &value
	return b
}

// WithPaused sets the Paused field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Paused field is set to the value of the last call.
func (b *ClusterSpecApplyConfiguration) WithPaused(value bool) *ClusterSpecApplyConfiguration {
	b.Paused = &value
	return b
}

// WithPod sets the Pod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pod field is set to the value of the last call.
func (b *ClusterSpecApplyConfiguration) WithPod(value *PodPolicyApplyConfiguration) *ClusterSpecApplyConfiguration {
	b.Pod = value
	return b
}

// WithTLS sets the TLS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLS field is set to the value of the last call.
func (b *ClusterSpecApplyConfiguration) WithTLS(value *TLSPolicyApplyConfiguration) *ClusterSpecApplyConfiguration {
	b.TLS = value
	return b
}
