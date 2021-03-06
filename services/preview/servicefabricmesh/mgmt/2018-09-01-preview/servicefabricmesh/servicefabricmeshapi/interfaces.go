package servicefabricmeshapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/servicefabricmesh/mgmt/2018-09-01-preview/servicefabricmesh"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result servicefabricmesh.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result servicefabricmesh.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*servicefabricmesh.OperationsClient)(nil)

// SecretClientAPI contains the set of methods on the SecretClient type.
type SecretClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, secretResourceName string, secretResourceDescription servicefabricmesh.SecretResourceDescription) (result servicefabricmesh.SecretResourceDescription, err error)
	Delete(ctx context.Context, resourceGroupName string, secretResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, secretResourceName string) (result servicefabricmesh.SecretResourceDescription, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicefabricmesh.SecretResourceDescriptionListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result servicefabricmesh.SecretResourceDescriptionListIterator, err error)
	ListBySubscription(ctx context.Context) (result servicefabricmesh.SecretResourceDescriptionListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result servicefabricmesh.SecretResourceDescriptionListIterator, err error)
}

var _ SecretClientAPI = (*servicefabricmesh.SecretClient)(nil)

// SecretValueClientAPI contains the set of methods on the SecretValueClient type.
type SecretValueClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, secretValueResourceDescription servicefabricmesh.SecretValueResourceDescription) (result servicefabricmesh.SecretValueResourceDescription, err error)
	Delete(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (result servicefabricmesh.SecretValueResourceDescription, err error)
	List(ctx context.Context, resourceGroupName string, secretResourceName string) (result servicefabricmesh.SecretValueResourceDescriptionListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, secretResourceName string) (result servicefabricmesh.SecretValueResourceDescriptionListIterator, err error)
	ListValue(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string) (result servicefabricmesh.SecretValue, err error)
}

var _ SecretValueClientAPI = (*servicefabricmesh.SecretValueClient)(nil)

// VolumeClientAPI contains the set of methods on the VolumeClient type.
type VolumeClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, volumeResourceName string, volumeResourceDescription servicefabricmesh.VolumeResourceDescription) (result servicefabricmesh.VolumeResourceDescription, err error)
	Delete(ctx context.Context, resourceGroupName string, volumeResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, volumeResourceName string) (result servicefabricmesh.VolumeResourceDescription, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicefabricmesh.VolumeResourceDescriptionListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result servicefabricmesh.VolumeResourceDescriptionListIterator, err error)
	ListBySubscription(ctx context.Context) (result servicefabricmesh.VolumeResourceDescriptionListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result servicefabricmesh.VolumeResourceDescriptionListIterator, err error)
}

var _ VolumeClientAPI = (*servicefabricmesh.VolumeClient)(nil)

// NetworkClientAPI contains the set of methods on the NetworkClient type.
type NetworkClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, networkResourceName string, networkResourceDescription servicefabricmesh.NetworkResourceDescription) (result servicefabricmesh.NetworkResourceDescription, err error)
	Delete(ctx context.Context, resourceGroupName string, networkResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, networkResourceName string) (result servicefabricmesh.NetworkResourceDescription, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicefabricmesh.NetworkResourceDescriptionListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result servicefabricmesh.NetworkResourceDescriptionListIterator, err error)
	ListBySubscription(ctx context.Context) (result servicefabricmesh.NetworkResourceDescriptionListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result servicefabricmesh.NetworkResourceDescriptionListIterator, err error)
}

var _ NetworkClientAPI = (*servicefabricmesh.NetworkClient)(nil)

// GatewayClientAPI contains the set of methods on the GatewayClient type.
type GatewayClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, gatewayResourceName string, gatewayResourceDescription servicefabricmesh.GatewayResourceDescription) (result servicefabricmesh.GatewayResourceDescription, err error)
	Delete(ctx context.Context, resourceGroupName string, gatewayResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, gatewayResourceName string) (result servicefabricmesh.GatewayResourceDescription, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicefabricmesh.GatewayResourceDescriptionListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result servicefabricmesh.GatewayResourceDescriptionListIterator, err error)
	ListBySubscription(ctx context.Context) (result servicefabricmesh.GatewayResourceDescriptionListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result servicefabricmesh.GatewayResourceDescriptionListIterator, err error)
}

var _ GatewayClientAPI = (*servicefabricmesh.GatewayClient)(nil)

// ApplicationClientAPI contains the set of methods on the ApplicationClient type.
type ApplicationClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, applicationResourceName string, applicationResourceDescription servicefabricmesh.ApplicationResourceDescription) (result servicefabricmesh.ApplicationResourceDescription, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationResourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, applicationResourceName string) (result servicefabricmesh.ApplicationResourceDescription, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicefabricmesh.ApplicationResourceDescriptionListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result servicefabricmesh.ApplicationResourceDescriptionListIterator, err error)
	ListBySubscription(ctx context.Context) (result servicefabricmesh.ApplicationResourceDescriptionListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result servicefabricmesh.ApplicationResourceDescriptionListIterator, err error)
}

var _ ApplicationClientAPI = (*servicefabricmesh.ApplicationClient)(nil)

// ServiceClientAPI contains the set of methods on the ServiceClient type.
type ServiceClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string) (result servicefabricmesh.ServiceResourceDescription, err error)
	List(ctx context.Context, resourceGroupName string, applicationResourceName string) (result servicefabricmesh.ServiceResourceDescriptionListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, applicationResourceName string) (result servicefabricmesh.ServiceResourceDescriptionListIterator, err error)
}

var _ ServiceClientAPI = (*servicefabricmesh.ServiceClient)(nil)

// ServiceReplicaClientAPI contains the set of methods on the ServiceReplicaClient type.
type ServiceReplicaClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, replicaName string) (result servicefabricmesh.ServiceReplicaDescription, err error)
	List(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string) (result servicefabricmesh.ServiceReplicaDescriptionListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string) (result servicefabricmesh.ServiceReplicaDescriptionListIterator, err error)
}

var _ ServiceReplicaClientAPI = (*servicefabricmesh.ServiceReplicaClient)(nil)

// CodePackageClientAPI contains the set of methods on the CodePackageClient type.
type CodePackageClientAPI interface {
	GetContainerLogs(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, replicaName string, codePackageName string, tail *int32) (result servicefabricmesh.ContainerLogs, err error)
}

var _ CodePackageClientAPI = (*servicefabricmesh.CodePackageClient)(nil)
