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
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// IntervalsQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/363111664e81786557afe06e68221018847b3676/specification/_types/query_dsl/fulltext.ts#L116-L125
type IntervalsQuery struct {
	AllOf      *IntervalsAllOf    `json:"all_of,omitempty"`
	AnyOf      *IntervalsAnyOf    `json:"any_of,omitempty"`
	Boost      *float32           `json:"boost,omitempty"`
	Fuzzy      *IntervalsFuzzy    `json:"fuzzy,omitempty"`
	Match      *IntervalsMatch    `json:"match,omitempty"`
	Prefix     *IntervalsPrefix   `json:"prefix,omitempty"`
	QueryName_ *string            `json:"_name,omitempty"`
	Wildcard   *IntervalsWildcard `json:"wildcard,omitempty"`
}

func (s *IntervalsQuery) UnmarshalJSON(data []byte) error {

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

		case "all_of":
			if err := dec.Decode(&s.AllOf); err != nil {
				return err
			}

		case "any_of":
			if err := dec.Decode(&s.AnyOf); err != nil {
				return err
			}

		case "boost":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "fuzzy":
			if err := dec.Decode(&s.Fuzzy); err != nil {
				return err
			}

		case "match":
			if err := dec.Decode(&s.Match); err != nil {
				return err
			}

		case "prefix":
			if err := dec.Decode(&s.Prefix); err != nil {
				return err
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.QueryName_ = &o

		case "wildcard":
			if err := dec.Decode(&s.Wildcard); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIntervalsQuery returns a IntervalsQuery.
func NewIntervalsQuery() *IntervalsQuery {
	r := &IntervalsQuery{}

	return r
}
