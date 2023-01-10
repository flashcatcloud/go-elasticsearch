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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Gets configuration and usage information about inference trained models.
package mltrainedmodels

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/bytes"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlTrainedModels struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	modelid string
}

// NewMlTrainedModels type alias for index.
type NewMlTrainedModels func() *MlTrainedModels

// NewMlTrainedModelsFunc returns a new instance of MlTrainedModels with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMlTrainedModelsFunc(tp elastictransport.Interface) NewMlTrainedModels {
	return func() *MlTrainedModels {
		n := New(tp)

		return n
	}
}

// Gets configuration and usage information about inference trained models.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cat-trained-model.html
func New(tp elastictransport.Interface) *MlTrainedModels {
	r := &MlTrainedModels{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *MlTrainedModels) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("trained_models")

		method = http.MethodGet
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		path.WriteString(r.modelid)

		method = http.MethodGet
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r MlTrainedModels) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the MlTrainedModels query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r MlTrainedModels) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the MlTrainedModels headers map.
func (r *MlTrainedModels) Header(key, value string) *MlTrainedModels {
	r.headers.Set(key, value)

	return r
}

// ModelId The ID of the trained models stats to fetch
// API Name: modelid
func (r *MlTrainedModels) ModelId(v string) *MlTrainedModels {
	r.paramSet |= modelidMask
	r.modelid = v

	return r
}

// AllowNoMatch Whether to ignore if a wildcard expression matches no trained models. (This
// includes `_all` string or when no trained models have been specified)
// API name: allow_no_match
func (r *MlTrainedModels) AllowNoMatch(b bool) *MlTrainedModels {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// Bytes The unit in which to display byte values
// API name: bytes
func (r *MlTrainedModels) Bytes(enum bytes.Bytes) *MlTrainedModels {
	r.values.Set("bytes", enum.String())

	return r
}

// H Comma-separated list of column names to display
// API name: h
func (r *MlTrainedModels) H(value string) *MlTrainedModels {
	r.values.Set("h", value)

	return r
}

// S Comma-separated list of column names or column aliases to sort by
// API name: s
func (r *MlTrainedModels) S(value string) *MlTrainedModels {
	r.values.Set("s", value)

	return r
}

// From skips a number of trained models
// API name: from
func (r *MlTrainedModels) From(i int) *MlTrainedModels {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// Size specifies a max number of trained models to get
// API name: size
func (r *MlTrainedModels) Size(i int) *MlTrainedModels {
	r.values.Set("size", strconv.Itoa(i))

	return r
}
