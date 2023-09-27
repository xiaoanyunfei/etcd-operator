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

package applyconfiguration

import (
	v1beta2 "github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	etcdv1beta2 "github.com/coreos/etcd-operator/pkg/generated/applyconfiguration/etcd/v1beta2"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=etcd.database.coreos.com, Version=v1beta2
	case v1beta2.SchemeGroupVersion.WithKind("ABSBackupSource"):
		return &etcdv1beta2.ABSBackupSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("ABSRestoreSource"):
		return &etcdv1beta2.ABSRestoreSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("BackupPolicy"):
		return &etcdv1beta2.BackupPolicyApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("BackupSource"):
		return &etcdv1beta2.BackupSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("BackupSpec"):
		return &etcdv1beta2.BackupSpecApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("BackupStatus"):
		return &etcdv1beta2.BackupStatusApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("ClusterCondition"):
		return &etcdv1beta2.ClusterConditionApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("ClusterSpec"):
		return &etcdv1beta2.ClusterSpecApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("ClusterStatus"):
		return &etcdv1beta2.ClusterStatusApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("EtcdBackup"):
		return &etcdv1beta2.EtcdBackupApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("EtcdCluster"):
		return &etcdv1beta2.EtcdClusterApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("EtcdClusterRef"):
		return &etcdv1beta2.EtcdClusterRefApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("EtcdRestore"):
		return &etcdv1beta2.EtcdRestoreApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("GCSBackupSource"):
		return &etcdv1beta2.GCSBackupSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("GCSRestoreSource"):
		return &etcdv1beta2.GCSRestoreSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("MemberSecret"):
		return &etcdv1beta2.MemberSecretApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("MembersStatus"):
		return &etcdv1beta2.MembersStatusApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("OSSBackupSource"):
		return &etcdv1beta2.OSSBackupSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("OSSRestoreSource"):
		return &etcdv1beta2.OSSRestoreSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("PodPolicy"):
		return &etcdv1beta2.PodPolicyApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("RestoreSource"):
		return &etcdv1beta2.RestoreSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("RestoreSpec"):
		return &etcdv1beta2.RestoreSpecApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("RestoreStatus"):
		return &etcdv1beta2.RestoreStatusApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("S3BackupSource"):
		return &etcdv1beta2.S3BackupSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("S3RestoreSource"):
		return &etcdv1beta2.S3RestoreSourceApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("StaticTLS"):
		return &etcdv1beta2.StaticTLSApplyConfiguration{}
	case v1beta2.SchemeGroupVersion.WithKind("TLSPolicy"):
		return &etcdv1beta2.TLSPolicyApplyConfiguration{}

	}
	return nil
}
