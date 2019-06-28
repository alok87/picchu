// +build !ignore_autogenerated

// Copyright © 2019 A Medium Corporation.
// Licensed under the Apache License, Version 2.0; see the NOTICE file.

// Code generated by openapi. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package openapi

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.Cluster":              schema_pkg_apis_picchu_v1alpha1_Cluster(ref),
		"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecrets":       schema_pkg_apis_picchu_v1alpha1_ClusterSecrets(ref),
		"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretsSpec":   schema_pkg_apis_picchu_v1alpha1_ClusterSecretsSpec(ref),
		"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretsStatus": schema_pkg_apis_picchu_v1alpha1_ClusterSecretsStatus(ref),
		"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManager":       schema_pkg_apis_picchu_v1alpha1_ReleaseManager(ref),
		"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManagerSpec":   schema_pkg_apis_picchu_v1alpha1_ReleaseManagerSpec(ref),
		"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManagerStatus": schema_pkg_apis_picchu_v1alpha1_ReleaseManagerStatus(ref),
		"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.Revision":             schema_pkg_apis_picchu_v1alpha1_Revision(ref),
	}
}

func schema_pkg_apis_picchu_v1alpha1_Cluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Cluster is the Schema for the clusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSpec", "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_picchu_v1alpha1_ClusterSecrets(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterSecrets is the Schema for the clustersecrets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretsSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretsStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretsSpec", "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretsStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_picchu_v1alpha1_ClusterSecretsSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterSecretsSpec defines the desired state of ClusterSecrets",
				Properties: map[string]spec.Schema{
					"source": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretSource"),
						},
					},
					"target": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretTarget"),
						},
					},
				},
				Required: []string{"source", "target"},
			},
		},
		Dependencies: []string{
			"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretSource", "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ClusterSecretTarget"},
	}
}

func schema_pkg_apis_picchu_v1alpha1_ClusterSecretsStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterSecretsStatus defines the observed state of ClusterSecrets",
				Properties: map[string]spec.Schema{
					"secrets": {
						SchemaProps: spec.SchemaProps{
							Description: "Names of secrets copied to targets",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_picchu_v1alpha1_ReleaseManager(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReleaseManager is the Schema for the releasemanagers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManagerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManagerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManagerSpec", "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManagerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_picchu_v1alpha1_ReleaseManagerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReleaseManagerSpec defines the desired state of ReleaseManager",
				Properties: map[string]spec.Schema{
					"fleet": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"app": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"target": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"fleet", "app", "target"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_picchu_v1alpha1_ReleaseManagerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReleaseManagerStatus defines the observed state of ReleaseManager",
				Properties: map[string]spec.Schema{
					"revisions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManagerRevisionStatus"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.ReleaseManagerRevisionStatus"},
	}
}

func schema_pkg_apis_picchu_v1alpha1_Revision(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Revision is the Schema for the revisions API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.RevisionSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.RevisionStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.RevisionSpec", "go.medium.engineering/picchu/pkg/apis/picchu/v1alpha1.RevisionStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}
