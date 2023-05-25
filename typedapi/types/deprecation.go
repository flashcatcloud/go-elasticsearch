// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/363111664e81786557afe06e68221018847b3676

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deprecationlevel"
)

// Deprecation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/363111664e81786557afe06e68221018847b3676/specification/migration/deprecations/types.ts#L29-L35
type Deprecation struct {
	Details string `json:"details"`
	// Level The level property describes the significance of the issue.
	Level   deprecationlevel.DeprecationLevel `json:"level"`
	Message string                            `json:"message"`
	Url     string                            `json:"url"`
}

// NewDeprecation returns a Deprecation.
func NewDeprecation() *Deprecation {
	r := &Deprecation{}

	return r
}
