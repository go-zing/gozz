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
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
	"unsafe"
)

var BuffPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

// SplitKV2Map split string into in key-value pairs by separator
func SplitKV(str string, sep string) (key, value string) {
	sp := strings.SplitN(str, sep, 2)
	return sp[0], strings.Join(sp[1:], sep)
}

// SplitKV2Map split string into in key-value pairs by separator and set key-value into dst map
func SplitKV2Map(str string, sep string, dst map[string]string) {
	if len(str) > 0 {
		k, v := SplitKV(str, sep)
		if _, ok := dst[k]; ok {
			dst[k] += "," + v
		} else {
			dst[k] = v
		}
	}
}

// SplitKVSlice2Map split strings into in key-value pairs by separator and set key-value into dst map
func SplitKVSlice2Map(ss []string, sep string, dst map[string]string) {
	for _, str := range ss {
		SplitKV2Map(str, sep, dst)
	}
}

func JoinDocs(docs []string) string {
	return strings.TrimSpace(strings.Join(docs, "\n"))
}

// TrimPrefix check strings has prefix and return trimmed string and check result
func TrimPrefix(str, prefix string) (string, bool) {
	if strings.HasPrefix(str, prefix) {
		return str[len(prefix):], true
	} else {
		return str, false
	}
}

// Bytesf format string and values by fmt.Fprintf and return byte slice
func Bytesf(format string, v ...interface{}) []byte {
	buffer := BuffPool.Get().(*bytes.Buffer)
	buffer.Reset()
	_, _ = fmt.Fprintf(buffer, format, v...)
	return buffer.Bytes()
}

// Appendf format string and values by fmt.Fprintf and append into strings slice
func Appendf(ss *[]string, format string, v ...interface{}) {
	*ss = append(*ss, fmt.Sprintf(format, v...))
}

// KeySet provide a unique key set to deduplicated and sort keys
type KeySet map[string]struct{}

// Keys return all keys sorted in set
func (set KeySet) Keys() []string {
	if len(set) == 0 {
		return nil
	}
	tmp := make([]string, 0, len(set))
	for key := range set {
		tmp = append(tmp, key)
	}
	sort.Strings(tmp)
	return tmp
}

// Add to append multi keys in set
func (set KeySet) Add(keys []string) {
	for _, key := range keys {
		if len(key) > 0 {
			set[key] = struct{}{}
		}
	}
}

func UnsafeBytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func UnsafeString2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
