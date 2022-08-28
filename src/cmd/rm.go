/*
 * Copyright © 2019 – 2021 Red Hat Inc.
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
	"os"
	"strings"

	"github.com/akdev1l/toolbox/pkg/podman"
	"github.com/spf13/cobra"
)

var (
	rmFlags struct {
		deleteAll   bool
		forceDelete bool
	}
)

var rmCmd = &cobra.Command{
	Use:               "rm",
	Short:             "Remove one or more toolbox containers",
	RunE:              rm,
	ValidArgsFunction: completionContainerNamesFiltered,
}

func init() {
	flags := rmCmd.Flags()

	flags.BoolVarP(&rmFlags.deleteAll, "all", "a", false, "Remove all toolbox containers")

	flags.BoolVarP(&rmFlags.forceDelete,
		"force",
		"f",
		false,
		"Force the removal of running and paused toolbox containers")

	rootCmd.AddCommand(rmCmd)
}

func rm(cmd *cobra.Command, args []string) error {
	if rmFlags.deleteAll {
		toolboxContainers, err := getContainers()
		if err != nil {
			return err
		}

		for _, container := range toolboxContainers {
			containerID := container.ID
			if err := podman.RemoveContainer(containerID, rmFlags.forceDelete); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				continue
			}
		}
	} else {
		if len(args) == 0 {
			var builder strings.Builder
			fmt.Fprintf(&builder, "missing argument for \"rm\"\n")
			fmt.Fprintf(&builder, "Run '%s --help' for usage.", executableBase)

			errMsg := builder.String()
			return errors.New(errMsg)
		}

		for _, container := range args {
			if _, err := podman.IsToolboxContainer(container); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				continue
			}

			if err := podman.RemoveContainer(container, rmFlags.forceDelete); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				continue
			}
		}
	}

	return nil
}
