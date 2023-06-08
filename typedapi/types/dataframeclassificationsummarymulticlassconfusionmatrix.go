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

// DataframeClassificationSummaryMulticlassConfusionMatrix type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96/specification/ml/evaluate_data_frame/types.ts#L79-L82
type DataframeClassificationSummaryMulticlassConfusionMatrix struct {
	ConfusionMatrix       []ConfusionMatrixItem `json:"confusion_matrix"`
	OtherActualClassCount int                   `json:"other_actual_class_count"`
}

func (s *DataframeClassificationSummaryMulticlassConfusionMatrix) UnmarshalJSON(data []byte) error {

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

		case "confusion_matrix":
			if err := dec.Decode(&s.ConfusionMatrix); err != nil {
				return err
			}

		case "other_actual_class_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.OtherActualClassCount = value
			case float64:
				f := int(v)
				s.OtherActualClassCount = f
			}

		}
	}
	return nil
}

// NewDataframeClassificationSummaryMulticlassConfusionMatrix returns a DataframeClassificationSummaryMulticlassConfusionMatrix.
func NewDataframeClassificationSummaryMulticlassConfusionMatrix() *DataframeClassificationSummaryMulticlassConfusionMatrix {
	r := &DataframeClassificationSummaryMulticlassConfusionMatrix{}

	return r
}