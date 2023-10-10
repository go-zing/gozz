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
	"os"

	"github.com/spf13/cobra"

	_ "github.com/Just-maple/gozz/internal/plugins"
	"github.com/Just-maple/gozz/zcore"
)

var (
	cmd = cobra.Command{
		Use:          zcore.ExecName,
		SilenceUsage: true,
	}

	registry = zcore.PluginRegistry()
)

func init() {
	cmd.AddCommand(
		run,
		listCmd,
	)
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
