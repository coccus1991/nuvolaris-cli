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
	"os"
	"path/filepath"
)

const wsk_auth = "23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP"
const wsk_apihost = "http://localhost:3233"

func writeWskPropertiesFile() error {
	homedir, err := GetHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(homedir, ".nuvolaris", ".wskprops")
	if _, err := os.Stat(path); err == nil {
		os.Remove(path)
	}
	content := []byte("AUTH=" + wsk_auth + "\nAPIHOST=" + wsk_apihost)
	if err := os.WriteFile(path, content, 0600); err != nil {
		return err
	}
	fmt.Println(".wskprops file written")
	return nil
}
