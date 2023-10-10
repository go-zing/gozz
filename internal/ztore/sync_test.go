/*
 * Copyright (c) 2023 Maple Wu <justmaplewu@gmail.com>
 *   National Electronics and Computer Technology Center, Thailand
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ztore

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestInitMap(t *testing.T) {
	n := &initStore{}
	count := int64(0)
	for i := 0; i < 1000; i++ {
		go n.Init("test", func() interface{} {
			if atomic.AddInt64(&count, 1) > 1 {
				t.Failed()
			}
			return new(int)
		})
	}
	time.Sleep(time.Second * 3)
}
