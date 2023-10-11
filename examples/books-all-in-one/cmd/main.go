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

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Just-maple/gozz/examples/books-all-in-one/inject"
)

func main() {
	// initialize application
	app, cf, err := inject.Initialize_booksallinone_Application()
	if err != nil {
		panic(err)
	}
	defer cf()

	controller := make(map[string]map[string]http.HandlerFunc)

	// range apis to bind http controller
	app.Apis.Range(func(i interface{}, apis []map[string]interface{}) {
		for _, api := range apis {
			method := strings.ToUpper(api["method"].(string))
			pattern := strings.ToUpper(api["resource"].(string))
			options := api["options"].(map[string]string)

			if prefix := options["prefix"]; len(prefix) > 0 {
				// pattern prefix
				pattern = prefix + "/" + strings.TrimPrefix(pattern, "/")
			}

			invoke, ok := api["invoke"].(func(ctx context.Context, dec func(interface{}) error) (interface{}, error))
			if !ok {
				continue
			}

			// alloc pattern
			v, ok := controller[pattern]
			if !ok {
				controller[pattern] = make(map[string]http.HandlerFunc)
				v = controller[pattern]
			}

			// register method handler
			v[method] = func(writer http.ResponseWriter, request *http.Request) {
				if r, e := invoke(request.Context(), func(interface{}) error {
					// TODO: encode params from request
					return nil
				}); e != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					_, _ = fmt.Fprintf(writer, "error: %v", e)
				} else {
					writer.WriteHeader(http.StatusOK)
					_ = json.NewEncoder(writer).Encode(r)
				}
			}
		}
	})

	// register http mux
	for pattern, methods := range controller {
		http.Handle(pattern, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if handle, ok := methods[request.Method]; ok {
				handle(writer, request)
			} else {
				writer.WriteHeader(http.StatusMethodNotAllowed)
			}
		}))
	}

	// listen and serve
	_ = http.ListenAndServe(":0", nil)
}
