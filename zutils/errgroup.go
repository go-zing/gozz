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

package zutils

import (
	"sync"
)

// ErrGroup is a simple version of golang.org/x/sync/errgroup.Group
type ErrGroup struct {
	wait sync.WaitGroup
	once sync.Once
	err  error
}

func (g *ErrGroup) Wait() error { g.wait.Wait(); return g.err }

func (g *ErrGroup) Go(f func() error) {
	g.wait.Add(1)
	go func() {
		defer g.wait.Done()
		if err := f(); err != nil {
			g.once.Do(func() { g.err = err })
		}
	}()
}
