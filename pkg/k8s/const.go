// Copyright 2016-2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

import (
	"github.com/cilium/cilium/common"

	"k8s.io/kubernetes/pkg/kubelet/types"
)

const (
	// AnnotationIsolationNS is the annotation key used in the annotation
	// map for the network isolation on the respective namespace.
	AnnotationIsolationNS = "net.beta.kubernetes.io/network-policy"
	// AnnotationName is an optional annotation to the NetworkPolicy
	// resource which specifies the name of the policy node to which all
	// rules should be applied to.
	AnnotationName = "io.cilium.name"
	// EnvNodeNameSpec is the environment label used by Kubernetes to
	// specify the node's name.
	EnvNodeNameSpec = "K8S_NODE_NAME"
	// LabelSource is the default label source for the labels imported from
	// kubernetes.
	LabelSource = "k8s"
	// LabelSourceKey is the default path to the policy node
	// received from kubernetes.
	LabelSourceKey = common.BaseLabelSourceExtPrefix + LabelSource
	// LabelSourceKeyPrefix is the DefaultPolicyParentPath with the
	// NodePathDelimiter.
	LabelSourceKeyPrefix = LabelSourceKey + common.PathDelimiter

	// PolicyLabelName is the name of the policy label which refers to the
	// k8s policy name
	PolicyLabelName = "io.cilium.k8s-policy-name"
	// PodNamespaceLabel is the label used in kubernetes containers to
	// specify which namespace they belong to.
	PodNamespaceLabel = types.KubernetesPodNamespaceLabel
	// PodNamespaceLabelPrefix is the PodNamespaceLabel prefixed with
	// DefaultPolicyParentPathPrefix.
	PodNamespaceLabelPrefix = LabelSourceKeyPrefix + PodNamespaceLabel
	// PodNamespaceMetaLabels is the label used to store the labels of the
	// kubernetes namespace's labels.
	PodNamespaceMetaLabels = "ns-labels"
	// PodNamespaceMetaLabelsPrefix is the label used to store the labels of
	// the kubernetes namespace's labels.
	PodNamespaceMetaLabelsPrefix = PodNamespaceMetaLabels + common.PathDelimiter
)
