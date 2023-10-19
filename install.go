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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	zcore "github.com/go-zing/gozz-core"
	"github.com/spf13/cobra"
)

const (
	installBuildDir = "/tmp/gozz/build/"
)

var (
	install = &cobra.Command{
		Use:     "install",
		Short:   "install external plugin from provided path or url",
		Example: zcore.ExecName + ` install [repository]`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := Install(args[0]); err != nil {
				_, _ = fmt.Fprint(os.Stderr, err.Error()+"\n")
				os.Exit(2)
			}
		},
	}

	coreDepPath = reflect.TypeOf(zcore.AnnotatedDecl{}).PkgPath()

	installBuildOutput string
	installBuildTarget string
)

func init() {
	flags := install.Flags()
	flags.StringVarP(&installBuildOutput, "output", "o", "", "specify install build output filename")
	flags.StringVarP(&installBuildTarget, "filepath", "f", "", "specify install build target relative filepath")
}

func Install(repository string) (err error) {
	if strings.Contains(repository, "://") {
		return doInstallRemote(repository)
	}

	// local path
	if f, e := os.Stat(repository); e == nil {
		if repository, err = filepath.Abs(repository); err != nil {
			return
		}
		if !f.IsDir() {
			if len(installBuildTarget) == 0 {
				repository = filepath.Dir(repository)
				installBuildTarget = f.Name()
			} else {
				return fmt.Errorf("invalid repository directory %s with filepath: %s ", repository, installBuildTarget)
			}
		}
		return doInstall(repository)
	}

	return doInstallRemote(repository)
}

func doInstallRemote(repository string) (err error) {
	t := time.Now().Format("060102150405")
	dir := filepath.Join(installBuildDir, t)
	_ = os.MkdirAll(installBuildDir, 0o775)
	defer os.RemoveAll(dir)

	buildCmd := exec.Command("sh", "-c", fmt.Sprintf("git clone --depth=1 %s %s", repository, t))
	buildCmd.Dir = installBuildDir
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err = buildCmd.Run(); err != nil {
		return
	}
	return doInstall(dir)
}

func getGoenv(dir string) (env map[string]string, err error) {
	goenv, err := zcore.ExecCommand("go env", dir)
	if err != nil {
		return
	}
	env = make(map[string]string)
	for _, line := range strings.Split(goenv, "\n") {
		if line = strings.TrimSpace(line); len(line) == 0 {
			continue
		}
		if kv := strings.SplitN(line, "=", 2); len(kv) >= 2 {
			env[kv[0]], _ = strconv.Unquote(kv[1])
		}
	}
	return
}

func getCoreVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	for _, m := range bi.Deps {
		if m.Path == coreDepPath {
			return m.Version
		}
	}
	return ""
}

func doInstall(dir string) (err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	buildDst := filepath.Join(wd, "tmp.so")
	args := []string{"build", "--buildmode=plugin", "-o=" + buildDst}

	if len(installBuildTarget) == 0 {
		args = append(args, "./")
	} else if installBuildTarget, err = filepath.Abs(filepath.Join(dir, installBuildTarget)); err != nil {
		return
	} else if info, e := os.Stat(installBuildTarget); e != nil {
		return e
	} else if info.IsDir() {
		dir = installBuildTarget
		args = append(args, "./")
	} else {
		// computed relative directory
		dir, installBuildTarget = filepath.Split(installBuildTarget)
		args = append(args, "./"+installBuildTarget)
	}

	// get env
	goenv, err := getGoenv(dir)
	if err != nil {
		return
	}

	// validate runtime
	runtimeVersion := runtime.Version()
	if v := goenv["GOVERSION"]; runtimeVersion != v {
		return fmt.Errorf("gozz runtime GOVERSION %s mismatch: %s", runtimeVersion, v)
	}

	// replace core mod version
	if cv := getCoreVersion(); len(cv) > 0 {
		if _, err = zcore.ExecCommand(
			fmt.Sprintf("go mod edit -replace=%s=%s@%s && go mod tidy", coreDepPath, coreDepPath, cv),
			filepath.Dir(goenv["GOMOD"]),
		); err != nil {
			return
		}
	}

	// build env
	envs := os.Environ()
	for i := range envs {
		if env := envs[i]; strings.HasPrefix(env, "GOOS=") {
			envs[i] = "GOOS=" + runtime.GOOS
		} else if strings.HasPrefix(env, "GOARCH=") {
			envs[i] = "GOARCH=" + runtime.GOARCH
		}
	}

	// build
	buildCmd := exec.Command("go", args...)
	buildCmd.Dir = dir
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	buildCmd.Env = envs
	if err = buildCmd.Run(); err != nil {
		return
	}
	// validate
	name, err := zcore.LoadExtension(buildDst)
	if err != nil {
		return
	}
	// install
	if len(installBuildOutput) > 0 {
		return os.Rename(buildDst, installBuildOutput)
	}
	if err = os.MkdirAll(pluginDir, 0o755); err != nil {
		return
	}
	return os.Rename(buildDst, filepath.Join(pluginDir, name+".so"))
}
