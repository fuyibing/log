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
// date: 2023-03-03

package configurer

type (
	// ConfigBucket
	// expose bucket (memory queue) configuration methods.
	ConfigBucket interface {
		GetBucketBatch() int
		GetBucketCapacity() int
		GetBucketConcurrency() int32
		GetBucketFrequency() int
	}
)

// Getter

func (o *config) GetBucketBatch() int         { return o.BucketBatch }
func (o *config) GetBucketCapacity() int      { return o.BucketCapacity }
func (o *config) GetBucketConcurrency() int32 { return o.BucketConcurrency }
func (o *config) GetBucketFrequency() int     { return o.BucketFrequency }

// Setter

func (o *Setter) SetBucketBatch(n int) *Setter         { o.config.BucketBatch = n; return o }
func (o *Setter) SetBucketCapacity(n int) *Setter      { o.config.BucketCapacity = n; return o }
func (o *Setter) SetBucketConcurrency(n int32) *Setter { o.config.BucketConcurrency = n; return o }
func (o *Setter) SetBucketFrequency(n int) *Setter     { o.config.BucketFrequency = n; return o }

// Access

func (o *config) defaultBucket() {
	if o.BucketBatch == 0 {
		o.BucketBatch = defaultBucketBatch
	}
	if o.BucketCapacity == 0 {
		o.BucketCapacity = defaultBucketCapacity
	}
	if o.BucketConcurrency == 0 {
		o.BucketConcurrency = defaultBucketConcurrency
	}
	if o.BucketFrequency == 0 {
		o.BucketFrequency = defaultBucketFrequency
	}
}
