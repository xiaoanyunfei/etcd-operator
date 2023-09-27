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

// MemberSecretApplyConfiguration represents an declarative configuration of the MemberSecret type for use
// with apply.
type MemberSecretApplyConfiguration struct {
	PeerSecret   *string `json:"peerSecret,omitempty"`
	ServerSecret *string `json:"serverSecret,omitempty"`
}

// MemberSecretApplyConfiguration constructs an declarative configuration of the MemberSecret type for use with
// apply.
func MemberSecret() *MemberSecretApplyConfiguration {
	return &MemberSecretApplyConfiguration{}
}

// WithPeerSecret sets the PeerSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PeerSecret field is set to the value of the last call.
func (b *MemberSecretApplyConfiguration) WithPeerSecret(value string) *MemberSecretApplyConfiguration {
	b.PeerSecret = &value
	return b
}

// WithServerSecret sets the ServerSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerSecret field is set to the value of the last call.
func (b *MemberSecretApplyConfiguration) WithServerSecret(value string) *MemberSecretApplyConfiguration {
	b.ServerSecret = &value
	return b
}
