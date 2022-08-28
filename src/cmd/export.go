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
	"path"

	"github.com/akdev1l/toolbox/pkg/podman"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	exportFlags struct {
		container   string
		appName     string
		serviceName string
		binaryPath  string
		extraFlags  string
		deleteMode  bool
		exportPath  string
	}
)

var exportCmd = &cobra.Command{
	Use:               "export",
	Short:             "Exports an application, binary or service to the host",
	RunE:              export,
	ValidArgsFunction: completionEmpty,
}

func init() {
	flags := exportCmd.Flags()

	flags.StringVarP(&exportFlags.container,
		"container",
		"c",
		"",
		"container name",
	)
	flags.StringVarP(&exportFlags.appName,
		"app",
		"a",
		"",
		"applicatio to export",
	)
	flags.StringVarP(&exportFlags.serviceName,
		"service",
		"s",
		"",
		"systemd service name to export",
	)
	flags.StringVarP(&exportFlags.binaryPath,
		"bin",
		"b",
		"",
		"path to binary to export",
	)
	flags.StringVarP(&exportFlags.exportPath,
		"export-path",
		"p",
		"$HOME/.local/bin",
		"Path to export binary applications",
	)
	flags.StringVarP(&exportFlags.extraFlags,
		"extra-flags",
		"e",
		"",
		"Set this option to pass extra flags to the application",
	)
	flags.BoolVarP(&exportFlags.deleteMode,
		"delete",
		"d",
		false,
		"Set this flag to export a command line application",
	)

	exportCmd.MarkFlagRequired("container")
	rootCmd.AddCommand(exportCmd)
}

func export(cmd *cobra.Command, args []string) error {
	if exportFlags.appName != "" {
		fmt.Println("exporting application")
	} else if exportFlags.binaryPath != "" {
		exportBinary(exportFlags.container, exportFlags.binaryPath, exportFlags.exportPath, exportFlags.deleteMode)
	} else if exportFlags.serviceName != "" {
		fmt.Println("export service")
	} else {
		return errors.New("please pass --app, --bin or --service")
	}

	return nil
}

func exportBinary(container string, binaryPath string, exportPath string, deleteMode bool) {
	var binaryName = path.Base(binaryPath)
	var shimBinaryPath = path.Join(os.ExpandEnv(exportPath), binaryName)
	var shellTemplate = fmt.Sprintf(`#!/bin/sh
toolbox run -c '%s' '%s' %s
`, container, binaryPath, "")

	if deleteMode {
		if _, err := os.Stat(shimBinaryPath); err == nil {
			logrus.Infof("removing binary shim located at %s", shimBinaryPath)
			os.Remove(shimBinaryPath)
		} else if errors.Is(err, os.ErrNotExist) {
			logrus.Warnf("shim binary not found at %s - skipping.", shimBinaryPath)
		}

		return
	}

	if _, err := os.Stat(shimBinaryPath); err == nil {
		logrus.Errorf("file %s already exists", shimBinaryPath)
	} else if errors.Is(err, os.ErrNotExist) {
		if containerExists, err := podman.ContainerExists(container); containerExists && err == nil {
			f, err := os.Create(shimBinaryPath)
			if err != nil {
				logrus.Error(err)
				logrus.Errorf("error creating file '%s'", shimBinaryPath)
			}
			f.WriteString(shellTemplate)
			f.Chmod(0755)
		} else {
			logrus.Errorf("container '%s' does not exist", container)
		}
	}
}
