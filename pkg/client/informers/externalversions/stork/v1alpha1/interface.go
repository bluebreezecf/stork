/*
Copyright 2018 Openstorage.org

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/libopenstorage/stork/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterPairs returns a ClusterPairInformer.
	ClusterPairs() ClusterPairInformer
	// Migrations returns a MigrationInformer.
	Migrations() MigrationInformer
	// Rules returns a RuleInformer.
	Rules() RuleInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterPairs returns a ClusterPairInformer.
func (v *version) ClusterPairs() ClusterPairInformer {
	return &clusterPairInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Migrations returns a MigrationInformer.
func (v *version) Migrations() MigrationInformer {
	return &migrationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Rules returns a RuleInformer.
func (v *version) Rules() RuleInformer {
	return &ruleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}