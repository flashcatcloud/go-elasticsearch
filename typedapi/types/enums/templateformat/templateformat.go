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
// https://github.com/elastic/elasticsearch-specification/tree/50c316c036cf0c3f567011c2bc24e7d2e1b8c781

// Package templateformat
package templateformat

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/50c316c036cf0c3f567011c2bc24e7d2e1b8c781/specification/security/_types/RoleTemplate.ts#L22-L25
type TemplateFormat struct {
	Name string
}

var (
	String = TemplateFormat{"string"}

	Json = TemplateFormat{"json"}
)

func (t TemplateFormat) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TemplateFormat) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "string":
		*t = String
	case "json":
		*t = Json
	default:
		*t = TemplateFormat{string(text)}
	}

	return nil
}

func (t TemplateFormat) String() string {
	return t.Name
}
