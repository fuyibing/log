// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// author: wsfuyibing <websearch@163.com>
// date: 2023-03-01

package conf

import (
	"github.com/fuyibing/log/v5/traces"
)

type (
	// ConfigBucket
	// 数据桶配置.
	ConfigBucket interface {
		GetBucketBatch() int
		GetBucketConcurrency() int32
		GetBucketCapacity() int
		GetBucketFrequency() int
	}
)

func (o *config) GetBucketBatch() int         { return o.BucketBatch }
func (o *config) GetBucketConcurrency() int32 { return o.BucketConcurrency }
func (o *config) GetBucketCapacity() int      { return o.BucketCapacity }
func (o *config) GetBucketFrequency() int     { return o.BucketFrequency }

func (o *config) initBucketDefaults() bool {
	if o.BucketBatch == 0 {
		o.BucketBatch = traces.BucketBatch
	}
	if o.BucketConcurrency == 0 {
		o.BucketConcurrency = traces.BucketConcurrency
	}
	if o.BucketCapacity == 0 {
		o.BucketCapacity = traces.BucketCapacity
	}
	if o.BucketFrequency == 0 {
		o.BucketFrequency = traces.BucketFrequency
	}
	return true
}
