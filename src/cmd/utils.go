/*
 * Copyright © 2020 – 2021 Red Hat Inc.
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

package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/akdev1l/toolbox/pkg/utils"
)

// askForConfirmation prints prompt to stdout and waits for response from the
// user
//
// Expected answers are: "yes", "y", "no", "n"
//
// Answers are internally converted to lower case.
//
// The default answer is "no" ([y/N])
func askForConfirmation(prompt string) bool {
	var retVal bool

	for {
		fmt.Printf("%s ", prompt)

		var response string

		fmt.Scanf("%s", &response)
		if response == "" {
			response = "n"
		} else {
			response = strings.ToLower(response)
		}

		if response == "no" || response == "n" {
			break
		} else if response == "yes" || response == "y" {
			retVal = true
			break
		}
	}

	return retVal
}

func createErrorContainerNotFound(container string) error {
	var builder strings.Builder
	fmt.Fprintf(&builder, "container %s not found\n", container)
	fmt.Fprintf(&builder, "Use the 'create' command to create a toolbox.\n")
	fmt.Fprintf(&builder, "Run '%s --help' for usage.", executableBase)

	errMsg := builder.String()
	return errors.New(errMsg)
}

func createErrorInvalidDistro() error {
	var builder strings.Builder
	fmt.Fprintf(&builder, "invalid argument for '--distro'\n")
	fmt.Fprintf(&builder, "Supported values are: %s\n", strings.Join(utils.GetSupportedDistros(), " "))
	fmt.Fprintf(&builder, "Run '%s --help' for usage.", executableBase)

	errMsg := builder.String()
	return errors.New(errMsg)
}

func createErrorInvalidRelease(distro string) error {
	var builder strings.Builder
	fmt.Fprintf(&builder, "invalid argument for '--release'\n")
	fmt.Fprintf(&builder, "Supported values for distribution %s are in format: %s\n", distro, utils.GetReleaseFormat(distro))
	fmt.Fprintf(&builder, "Run '%s --help' for usage.", executableBase)

	errMsg := builder.String()
	return errors.New(errMsg)
}
