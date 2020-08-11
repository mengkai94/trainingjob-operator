// +build !ignore_autogenerated

/*
Copyright (c) 2020 PaddlePaddle Authors. All Rights Reserved.

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
// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&AITrainingJob{}, func(obj interface{}) { SetObjectDefaults_AITrainingJob(obj.(*AITrainingJob)) })
	scheme.AddTypeDefaultingFunc(&AITrainingJobList{}, func(obj interface{}) { SetObjectDefaults_AITrainingJobList(obj.(*AITrainingJobList)) })
	return nil
}

func SetObjectDefaults_AITrainingJob(in *AITrainingJob) {
	SetDefaults_AITrainingJob(in)
}

func SetObjectDefaults_AITrainingJobList(in *AITrainingJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_AITrainingJob(a)
	}
}
