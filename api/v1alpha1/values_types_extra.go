// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/istio-ecosystem/sail-operator/pkg/helm"
)

type SDSConfigToken struct {
	Aud string `json:"aud,omitempty"`
}

func (x *Values) ToHelmValues() helm.HelmValues {
	var obj helm.HelmValues
	data, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(data, &obj); err != nil {
		panic(err)
	}
	return obj
}

func ValuesFromHelmValues(helmValues helm.HelmValues) (*Values, error) {
	data, err := json.Marshal(helmValues)
	if err != nil {
		return nil, err
	}

	values := Values{}
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&values)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal into Values struct: %v:\n%v", err, string(data))
	}
	return &values, nil
}