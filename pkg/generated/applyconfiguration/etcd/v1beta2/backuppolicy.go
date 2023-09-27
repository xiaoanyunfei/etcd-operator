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

// BackupPolicyApplyConfiguration represents an declarative configuration of the BackupPolicy type for use
// with apply.
type BackupPolicyApplyConfiguration struct {
	TimeoutInSecond        *int64 `json:"timeoutInSecond,omitempty"`
	BackupIntervalInSecond *int64 `json:"backupIntervalInSecond,omitempty"`
	MaxBackups             *int   `json:"maxBackups,omitempty"`
}

// BackupPolicyApplyConfiguration constructs an declarative configuration of the BackupPolicy type for use with
// apply.
func BackupPolicy() *BackupPolicyApplyConfiguration {
	return &BackupPolicyApplyConfiguration{}
}

// WithTimeoutInSecond sets the TimeoutInSecond field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TimeoutInSecond field is set to the value of the last call.
func (b *BackupPolicyApplyConfiguration) WithTimeoutInSecond(value int64) *BackupPolicyApplyConfiguration {
	b.TimeoutInSecond = &value
	return b
}

// WithBackupIntervalInSecond sets the BackupIntervalInSecond field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackupIntervalInSecond field is set to the value of the last call.
func (b *BackupPolicyApplyConfiguration) WithBackupIntervalInSecond(value int64) *BackupPolicyApplyConfiguration {
	b.BackupIntervalInSecond = &value
	return b
}

// WithMaxBackups sets the MaxBackups field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxBackups field is set to the value of the last call.
func (b *BackupPolicyApplyConfiguration) WithMaxBackups(value int) *BackupPolicyApplyConfiguration {
	b.MaxBackups = &value
	return b
}