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

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// NoriPartOfSpeechTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96/specification/_types/analysis/token_filters.ts#L273-L276
type NoriPartOfSpeechTokenFilter struct {
	Stoptags []string `json:"stoptags,omitempty"`
	Type     string   `json:"type,omitempty"`
	Version  *string  `json:"version,omitempty"`
}

func (s *NoriPartOfSpeechTokenFilter) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "stoptags":
			if err := dec.Decode(&s.Stoptags); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNoriPartOfSpeechTokenFilter returns a NoriPartOfSpeechTokenFilter.
func NewNoriPartOfSpeechTokenFilter() *NoriPartOfSpeechTokenFilter {
	r := &NoriPartOfSpeechTokenFilter{}

	r.Type = "nori_part_of_speech"

	return r
}