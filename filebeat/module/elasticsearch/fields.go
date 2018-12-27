// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "elasticsearch", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
