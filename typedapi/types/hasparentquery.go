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
	"strconv"
)

// HasParentQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96/specification/_types/query_dsl/joining.ts#L53-L61
type HasParentQuery struct {
	Boost          *float32   `json:"boost,omitempty"`
	IgnoreUnmapped *bool      `json:"ignore_unmapped,omitempty"`
	InnerHits      *InnerHits `json:"inner_hits,omitempty"`
	ParentType     string     `json:"parent_type"`
	Query          *Query     `json:"query,omitempty"`
	QueryName_     *string    `json:"_name,omitempty"`
	Score          *bool      `json:"score,omitempty"`
}

func (s *HasParentQuery) UnmarshalJSON(data []byte) error {

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

		case "ignore_unmapped":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreUnmapped = &value
			case bool:
				s.IgnoreUnmapped = &v
			}

		case "inner_hits":
			if err := dec.Decode(&s.InnerHits); err != nil {
				return err
			}

		case "parent_type":
			if err := dec.Decode(&s.ParentType); err != nil {
				return err
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.QueryName_ = &o

		case "score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Score = &value
			case bool:
				s.Score = &v
			}

		}
	}
	return nil
}

// NewHasParentQuery returns a HasParentQuery.
func NewHasParentQuery() *HasParentQuery {
	r := &HasParentQuery{}

	return r
}