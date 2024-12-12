/*
Copyright 2023 The Radius Authors.

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

package dataprovider

import (
	"context"

	"github.com/radius-project/radius/pkg/ucp/store"
)

// StorageProviderType represents types of storage provider.
type StorageProviderType string

const (
	// TypeAPIServer represents the Kubernetes APIServer provider.
	TypeAPIServer StorageProviderType = "apiserver"

	// TypeETCD represents the etcd provider.
	TypeETCD StorageProviderType = "etcd"

	// TypeInMemory represents the in-memory provider.
	TypeInMemory StorageProviderType = "inmemory"

	// TypePostgreSQL represents the PostgreSQL provider.
	TypePostgreSQL StorageProviderType = "postgresql"
)

//go:generate mockgen -typed -destination=./mock_datastorage_provider.go -package=dataprovider -self_package github.com/radius-project/radius/pkg/ucp/dataprovider github.com/radius-project/radius/pkg/ucp/dataprovider DataStorageProvider

// DataStorageProvider is an interfae to provide storage client.
type DataStorageProvider interface {
	// GetStorageClient creates or gets storage client.
	GetStorageClient(context.Context, string) (store.StorageClient, error)
}
