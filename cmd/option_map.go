/*
Copyright (c) 2020, Arm Limited and affiliates.
SPDX-License-Identifier: Apache-2.0
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"
)

type OptionMap map[string]string

func (m *OptionMap) String() string {
	if m == nil {
		return "{}"
	}

	return fmt.Sprintf("%v", *m)
}

func (m *OptionMap) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return fmt.Errorf("Argument format is <key>=<value>")
	}

	(*m)[parts[0]] = parts[1]

	return nil
}
