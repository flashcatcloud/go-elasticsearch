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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Retrieves configuration information for datafeeds.
package getdatafeeds

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDatafeeds struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	datafeedid string
}

// NewGetDatafeeds type alias for index.
type NewGetDatafeeds func() *GetDatafeeds

// NewGetDatafeedsFunc returns a new instance of GetDatafeeds with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetDatafeedsFunc(tp elastictransport.Interface) NewGetDatafeeds {
	return func() *GetDatafeeds {
		n := New(tp)

		return n
	}
}

// Retrieves configuration information for datafeeds.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed.html
func New(tp elastictransport.Interface) *GetDatafeeds {
	r := &GetDatafeeds{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetDatafeeds) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		path.WriteString(r.datafeedid)

		method = http.MethodGet
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")

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

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetDatafeeds) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetDatafeeds query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getdatafeeds.Response
func (r GetDatafeeds) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil

	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetDatafeeds) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

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

// Header set a key, value pair in the GetDatafeeds headers map.
func (r *GetDatafeeds) Header(key, value string) *GetDatafeeds {
	r.headers.Set(key, value)

	return r
}

// DatafeedId Identifier for the datafeed. It can be a datafeed identifier or a
// wildcard expression. If you do not specify one of these options, the API
// returns information about all datafeeds.
// API Name: datafeedid
func (r *GetDatafeeds) DatafeedId(v string) *GetDatafeeds {
	r.paramSet |= datafeedidMask
	r.datafeedid = v

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// 1. Contains wildcard expressions and there are no datafeeds that match.
// 2. Contains the `_all` string or no identifiers and there are no matches.
// 3. Contains wildcard expressions and there are only partial matches.
//
// The default value is `true`, which returns an empty `datafeeds` array
// when there are no matches and the subset of results when there are
// partial matches. If this parameter is `false`, the request returns a
// `404` status code when there are no matches or only partial matches.
// API name: allow_no_match
func (r *GetDatafeeds) AllowNoMatch(b bool) *GetDatafeeds {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// ExcludeGenerated Indicates if certain fields should be removed from the configuration on
// retrieval. This allows the configuration to be in an acceptable format to
// be retrieved and then added to another cluster.
// API name: exclude_generated
func (r *GetDatafeeds) ExcludeGenerated(b bool) *GetDatafeeds {
	r.values.Set("exclude_generated", strconv.FormatBool(b))

	return r
}