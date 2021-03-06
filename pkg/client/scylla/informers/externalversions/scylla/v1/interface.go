// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/scylladb/scylla-operator/pkg/client/scylla/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ScyllaClusters returns a ScyllaClusterInformer.
	ScyllaClusters() ScyllaClusterInformer
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

// ScyllaClusters returns a ScyllaClusterInformer.
func (v *version) ScyllaClusters() ScyllaClusterInformer {
	return &scyllaClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
