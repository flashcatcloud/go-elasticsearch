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
// https://github.com/elastic/elasticsearch-specification/tree/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96

package types

// LoggingAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96/specification/watcher/_types/Actions.ts#L281-L285
type LoggingAction struct {
	Category *string `json:"category,omitempty"`
	Level    *string `json:"level,omitempty"`
	Text     string  `json:"text"`
}

// NewLoggingAction returns a LoggingAction.
func NewLoggingAction() *LoggingAction {
	r := &LoggingAction{}

	return r
}