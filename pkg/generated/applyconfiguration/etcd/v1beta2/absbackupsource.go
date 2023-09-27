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

// ABSBackupSourceApplyConfiguration represents an declarative configuration of the ABSBackupSource type for use
// with apply.
type ABSBackupSourceApplyConfiguration struct {
	Path      *string `json:"path,omitempty"`
	ABSSecret *string `json:"absSecret,omitempty"`
}

// ABSBackupSourceApplyConfiguration constructs an declarative configuration of the ABSBackupSource type for use with
// apply.
func ABSBackupSource() *ABSBackupSourceApplyConfiguration {
	return &ABSBackupSourceApplyConfiguration{}
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *ABSBackupSourceApplyConfiguration) WithPath(value string) *ABSBackupSourceApplyConfiguration {
	b.Path = &value
	return b
}

// WithABSSecret sets the ABSSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ABSSecret field is set to the value of the last call.
func (b *ABSBackupSourceApplyConfiguration) WithABSSecret(value string) *ABSBackupSourceApplyConfiguration {
	b.ABSSecret = &value
	return b
}
