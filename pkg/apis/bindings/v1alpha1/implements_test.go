/*
Copyright 2019 The Knative Authors

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

package v1alpha1

import (
	"testing"

	rduckv1alpha1 "github.com/projectriff/bindings/pkg/apis/duck/v1alpha1"
	"knative.dev/pkg/apis/duck"
	duckv1alpha1 "knative.dev/pkg/apis/duck/v1alpha1"
)

func TestImplementsBinding(t *testing.T) {
	if err := duck.VerifyType(&ImageBinding{}, &duckv1alpha1.Binding{}); err != nil {
		t.Fatal(err)
	}
	if err := duck.VerifyType(&ProvisionedService{}, &rduckv1alpha1.Serviceable{}); err != nil {
		t.Fatal(err)
	}
}
