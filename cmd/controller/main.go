/*
Copyright 2019 The Tekton Authors

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
package main

import (
	"k8s.io/klog"
	"knative.dev/pkg/injection/sharedmain"

	"github.com/tektoncd/triggers/pkg/reconciler/v1alpha1/eventlistener"
)

const (
	// ControllerLogKey is the name of the logger for the controller cmd
	ControllerLogKey = "controller"
)

func main() {
	klog.InitFlags(nil)
	sharedmain.Main(ControllerLogKey,
		eventlistener.NewController,
	)
}
