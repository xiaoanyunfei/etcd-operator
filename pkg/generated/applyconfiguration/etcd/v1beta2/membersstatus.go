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

// MembersStatusApplyConfiguration represents an declarative configuration of the MembersStatus type for use
// with apply.
type MembersStatusApplyConfiguration struct {
	Ready   []string `json:"ready,omitempty"`
	Unready []string `json:"unready,omitempty"`
}

// MembersStatusApplyConfiguration constructs an declarative configuration of the MembersStatus type for use with
// apply.
func MembersStatus() *MembersStatusApplyConfiguration {
	return &MembersStatusApplyConfiguration{}
}

// WithReady adds the given value to the Ready field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ready field.
func (b *MembersStatusApplyConfiguration) WithReady(values ...string) *MembersStatusApplyConfiguration {
	for i := range values {
		b.Ready = append(b.Ready, values[i])
	}
	return b
}

// WithUnready adds the given value to the Unready field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Unready field.
func (b *MembersStatusApplyConfiguration) WithUnready(values ...string) *MembersStatusApplyConfiguration {
	for i := range values {
		b.Unready = append(b.Unready, values[i])
	}
	return b
}
