// Copyright 2020 The Operator-SDK Authors
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

package testing

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
)

var _ machinery.Template = &Kustomization{}

// Kustomization scaffolds the kustomization file for use
// during Ansible testing
type Kustomization struct {
	machinery.TemplateMixin
}

// SetTemplateDefaults implements machinery.Template
func (f *Kustomization) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("config", "testing", "kustomization.yaml")
	}

	f.TemplateBody = KustomizationTemplate

	f.IfExistsAction = machinery.Error

	return nil
}

const KustomizationTemplate = `# Adds namespace to all resources.
namespace: osdk-test

namePrefix: osdk-

# Labels to add to all resources and selectors.
#commonLabels:
#  someName: someValue

patches:
- path: manager_image.yaml
- path: debug_logs_patch.yaml
- path: ../default/manager_metrics_patch.yaml
  target:
    kind: Deployment

apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
- ../crd
- ../rbac
- ../manager
images:
- name: testing
  newName: testing-operator
`
