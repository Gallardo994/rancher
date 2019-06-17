/*
Copyright 2015 The Kubernetes Authors.

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

package hyperkube

import (
	"flag"

	"k8s.io/kubernetes/cmd/kube-controller-manager/app"
)

// NewKubeControllerManager creates a new hyperkube Server object that includes the
// description and flags.
func NewKubeControllerManager() *Server {
	command := app.NewControllerManagerCommand()

	hks := Server{
		name:            "controller-manager",
		AlternativeName: "kube-controller-manager",
		SimpleUsage:     "controller-manager",
		Long:            command.Long,
	}

	serverFlags := hks.Flags()
	serverFlags.AddFlagSet(command.Flags())

	command.Flags().AddGoFlagSet(flag.CommandLine)

	hks.Run = func(_ *Server, args []string, stopCh <-chan struct{}) error {
		command.SetArgs(args)
		return command.Execute()
	}

	return &hks
}