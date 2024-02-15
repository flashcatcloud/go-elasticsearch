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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// MatrixStatsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/50c316c036cf0c3f567011c2bc24e7d2e1b8c781/specification/_types/aggregations/Aggregate.ts#L757-L761
type MatrixStatsAggregate struct {
	DocCount int64               `json:"doc_count"`
	Fields   []MatrixStatsFields `json:"fields,omitempty"`
	Meta     Metadata            `json:"meta,omitempty"`
}

func (s *MatrixStatsAggregate) UnmarshalJSON(data []byte) error {

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

		case "doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.DocCount = value
			case float64:
				f := int64(v)
				s.DocCount = f
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMatrixStatsAggregate returns a MatrixStatsAggregate.
func NewMatrixStatsAggregate() *MatrixStatsAggregate {
	r := &MatrixStatsAggregate{}

	return r
}
