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

package api

import (
	"context"
)

//go:generate gozz run -p "api" ./

// +zz:api:./
type (
	T interface {
		// +zz:api:get:
		Empty()
		// +zz:api:get:
		Ret() (ret int)
		// +zz:api:get:
		Error() (err error)
		// +zz:api:get:
		RetError() (ret int, err error)
		// +zz:api:get:
		Context(ctx context.Context)
		// +zz:api:get:
		ContextRet(ctx context.Context) (ret int)
		// +zz:api:get:
		ContextError(ctx context.Context) (err error)
		// +zz:api:get:
		ContextRetError(ctx context.Context) (ret int, err error)
		// +zz:api:get:
		Param(param int)
		// +zz:api:get:
		ParamRet(param int) (ret error)
		// +zz:api:get:
		ParamError(param int) (err error)
		// +zz:api:get:
		ParamRetError(param int) (ret int, err error)
		// +zz:api:get:
		ContextParam(ctx context.Context, param int)
		// +zz:api:get:
		ContextParamRet(ctx context.Context, param int) (ret int)
		// +zz:api:get:
		ContextParamError(ctx context.Context, param int) (err error)
		// +zz:api:get:
		ContextParamRetError(ctx context.Context, param int) (ret int, err error)
		// +zz:api:get:
		ComplexParam(param map[context.Context][]struct {
			Field []func(context.Context) interface {
				context.Context
			}
		})
		// +zz:api:get:
		PtrParam(*int)
	}
)
