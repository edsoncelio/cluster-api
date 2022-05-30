//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.DiscoveryRequest":  schema_runtime_hooks_api_v1alpha1_DiscoveryRequest(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.DiscoveryResponse": schema_runtime_hooks_api_v1alpha1_DiscoveryResponse(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ExtensionHandler":  schema_runtime_hooks_api_v1alpha1_ExtensionHandler(ref),
		"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GroupVersionHook":  schema_runtime_hooks_api_v1alpha1_GroupVersionHook(ref),
	}
}

func schema_runtime_hooks_api_v1alpha1_DiscoveryRequest(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DiscoveryRequest represents the object of a discovery request.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_runtime_hooks_api_v1alpha1_DiscoveryResponse(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DiscoveryResponse represents the object received as a discovery response.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the call. One of \"Success\" or \"Failure\".\n\nPossible enum values:\n - `\"Failure\"` represents a failure response.\n - `\"Success\"` represents the success response.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"Failure", "Success"}},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human-readable description of the status of the call.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"handlers": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Handlers defines the current ExtensionHandlers supported by an Extension.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ExtensionHandler"),
									},
								},
							},
						},
					},
				},
				Required: []string{"status", "message", "handlers"},
			},
		},
		Dependencies: []string{
			"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.ExtensionHandler"},
	}
}

func schema_runtime_hooks_api_v1alpha1_ExtensionHandler(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExtensionHandler represents the discovery information of the extension which includes the hook it supports.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the ExtensionHandler.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"requestHook": {
						SchemaProps: spec.SchemaProps{
							Description: "RequestHook defines the versioned runtime hook which this ExtensionHandler serves.",
							Default:     map[string]interface{}{},
							Ref:         ref("sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GroupVersionHook"),
						},
					},
					"timeoutSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "TimeoutSeconds defines the timeout duration for client calls to the ExtensionHandler.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"failurePolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "FailurePolicy defines how failures in calls to the ExtensionHandler should be handled by a client.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "requestHook"},
			},
		},
		Dependencies: []string{
			"sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1.GroupVersionHook"},
	}
}

func schema_runtime_hooks_api_v1alpha1_GroupVersionHook(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GroupVersionHook defines the runtime hook when the ExtensionHandler is called.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion is the Version of the Hook",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"hook": {
						SchemaProps: spec.SchemaProps{
							Description: "Hook is the name of the hook",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"apiVersion", "hook"},
			},
		},
	}
}