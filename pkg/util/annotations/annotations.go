/*
Copyright 2024 The Karmada Authors.

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

// Package annotations provides utility functions for working with
// Kubernetes object annotations in the Karmada control plane.
 (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get a specific
// a Kubernetes object's metadata. Returns an empty string if the annotation
// does not exist or if the annotations map is nil.
func GetAnnotationValue(meta metav1 string {
	if meta.Annotations == nil {
n	}
	return meta.Annotations[key]
}

// HasAnnotation checks whether a specific annotation key exists on the object,
// regardless of its value.
func HasAnnotation(meta metav1.ObjectMeta, key string) bool {
	if meta.Annotations == nil {
		return false
	}
	_, ok := meta.Annotations[key]
	return ok
}

// SetAnnotation sets the value of a specific annotation key on the given
// object metadata. If the annotations map is nil, it initializes it.
func SetAnnotation(meta *metav1.ObjectMeta, key, value string) {
	if meta.Annotations == nil {
		meta.Annotations = make(map[string]string)
	}
	meta.Annotations[key] = value
}

// RemoveAnnotation removes the annotation with the given key from the
// object metadata. It is a no-op if the annotation does not exist.
func RemoveAnnotation(meta *metav1.ObjectMeta, key string) {
	if meta.Annotations == nil {
		return
	}
	delete(meta.Annotations, key)
}

// MergeAnnotations merges the src annotations into dst annotations.
// Existing keys in dst will be overwritten by values from src.
// Note: this modifies dst in place and also returns it for convenience.
// If src is nil, dst is returned unchanged.
func MergeAnnotations(dst, src map[string]string) map[string]string {
	if dst == nil {
		dst = make(map[string]string)
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
