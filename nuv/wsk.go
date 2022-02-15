// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/apache/openwhisk-cli/commands"
	"github.com/apache/openwhisk-cli/wski18n"
	"github.com/apache/openwhisk-client-go/whisk"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
)

type WskCmd struct {
	Args []string `arg:"" optional:"" name:"args" help:"wsk subcommand args"`
}

var cliDebug = os.Getenv("WSK_CLI_DEBUG") // Useful for tracing init() code

var T goi18n.TranslateFunc

func init() {
	if len(cliDebug) > 0 {
		whisk.SetDebug(true)
	}

	T = wski18n.T

	// Rest of CLI uses the Properties struct, so set the build time there
	commands.Properties.CLIVersion = CLI_VERSION
}

func Wsk(args ...string) error {
	baseArgs := []string{"wsk"}

	addWskProperties(&baseArgs)

	os.Args = append(baseArgs, args...)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println(T("Application exited unexpectedly"))
		}
	}()
	return commands.Execute()
}

func addWskProperties(baseArgs *[]string) error {
	homedir, err := GetHomeDir()

	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(filepath.Join(homedir, ".nuvolaris", ".wskprops"))

	if err != nil {
		return err
	}

	wskEnv, err := godotenv.Unmarshal(string(file))

	if err != nil {
		return err
	}

	apiHost, apiHostExist := wskEnv["APIHOST"]
	auth, authExist := wskEnv["AUTH"]

	if authExist && apiHostExist {
		*baseArgs = append(*baseArgs, "--apihost", apiHost, "--auth", auth)
	}

	return nil
}

func (wsk *WskCmd) Run() error {
	return Wsk(wsk.Args...)
}
