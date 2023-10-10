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
	"sync"
)

// initStore provide key mapping registry to init and store keyed object
type initStore struct {
	mu sync.RWMutex
	m  map[interface{}]interface{}
}

func (s *initStore) Init(key interface{}, fn func() interface{}) (r interface{}) {
	s.mu.RLock()
	r, ok := s.m[key]
	if ok {
		s.mu.RUnlock()
		return
	}
	s.mu.RUnlock()

	s.mu.Lock()
	// double check key
	if r, ok = s.m[key]; !ok {
		if r = fn(); s.m == nil {
			s.m = make(map[interface{}]interface{})
		}
		s.m[key] = r
	}
	s.mu.Unlock()
	return
}

// VersionStore provide a store with source load like single-flights and versioned cache store
type VersionStore struct {
	m initStore
}

type versionEntity struct {
	sync.Mutex
	version string
	value   interface{}
}

func (s *VersionStore) Load(key interface{}, version string, fn func() (interface{}, error)) (r interface{}, err error) {
	entity := s.m.Init(key, func() interface{} { return new(versionEntity) }).(*versionEntity)
	// lock
	entity.Lock()

	// version match
	if version == entity.version {
		entity.Unlock()
		return entity.value, nil
	}

	// get new value error
	if r, err = fn(); err != nil {
		entity.Unlock()
		return
	}

	// update store value and version
	entity.version = version
	entity.value = r

	// unlock
	entity.Unlock()
	return
}

func (s *VersionStore) Update(key, version string, v interface{}) {
	entity := s.m.Init(key, func() interface{} { return new(versionEntity) }).(*versionEntity)
	entity.Lock()
	entity.value = v
	entity.version = version
	entity.Unlock()
}
