// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Caffe2Job{}, func(obj interface{}) { SetObjectDefaults_Caffe2Job(obj.(*Caffe2Job)) })
	scheme.AddTypeDefaultingFunc(&Caffe2JobList{}, func(obj interface{}) { SetObjectDefaults_Caffe2JobList(obj.(*Caffe2JobList)) })
	return nil
}

func SetObjectDefaults_Caffe2Job(in *Caffe2Job) {
	SetDefaults_Caffe2Job(in)
}

func SetObjectDefaults_Caffe2JobList(in *Caffe2JobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Caffe2Job(a)
	}
}