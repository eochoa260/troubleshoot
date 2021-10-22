/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/replicatedhq/troubleshoot/pkg/longhorn/client/clientset/versioned/typed/longhorn/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLonghornV1beta1 struct {
	*testing.Fake
}

func (c *FakeLonghornV1beta1) BackingImages(namespace string) v1beta1.BackingImageInterface {
	return &FakeBackingImages{c, namespace}
}

func (c *FakeLonghornV1beta1) BackingImageDataSources(namespace string) v1beta1.BackingImageDataSourceInterface {
	return &FakeBackingImageDataSources{c, namespace}
}

func (c *FakeLonghornV1beta1) BackingImageManagers(namespace string) v1beta1.BackingImageManagerInterface {
	return &FakeBackingImageManagers{c, namespace}
}

func (c *FakeLonghornV1beta1) Backups(namespace string) v1beta1.BackupInterface {
	return &FakeBackups{c, namespace}
}

func (c *FakeLonghornV1beta1) BackupTargets(namespace string) v1beta1.BackupTargetInterface {
	return &FakeBackupTargets{c, namespace}
}

func (c *FakeLonghornV1beta1) BackupVolumes(namespace string) v1beta1.BackupVolumeInterface {
	return &FakeBackupVolumes{c, namespace}
}

func (c *FakeLonghornV1beta1) Engines(namespace string) v1beta1.EngineInterface {
	return &FakeEngines{c, namespace}
}

func (c *FakeLonghornV1beta1) EngineImages(namespace string) v1beta1.EngineImageInterface {
	return &FakeEngineImages{c, namespace}
}

func (c *FakeLonghornV1beta1) InstanceManagers(namespace string) v1beta1.InstanceManagerInterface {
	return &FakeInstanceManagers{c, namespace}
}

func (c *FakeLonghornV1beta1) Nodes(namespace string) v1beta1.NodeInterface {
	return &FakeNodes{c, namespace}
}

func (c *FakeLonghornV1beta1) RecurringJobs(namespace string) v1beta1.RecurringJobInterface {
	return &FakeRecurringJobs{c, namespace}
}

func (c *FakeLonghornV1beta1) Replicas(namespace string) v1beta1.ReplicaInterface {
	return &FakeReplicas{c, namespace}
}

func (c *FakeLonghornV1beta1) Settings(namespace string) v1beta1.SettingInterface {
	return &FakeSettings{c, namespace}
}

func (c *FakeLonghornV1beta1) ShareManagers(namespace string) v1beta1.ShareManagerInterface {
	return &FakeShareManagers{c, namespace}
}

func (c *FakeLonghornV1beta1) Volumes(namespace string) v1beta1.VolumeInterface {
	return &FakeVolumes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLonghornV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
