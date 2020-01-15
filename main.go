// Copyright 2020 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/compute/metadata"
)

type cmd func() error

func main() {
	cmds := map[string]cmd{
		"hostname":      print(metadata.Hostname),
		"external-ip":   print(metadata.ExternalIP),
		"internal-ip":   print(metadata.InternalIP),
		"instance-name": print(metadata.InstanceName),
		"zone":          print(metadata.Zone),
		"project-id":    print(metadata.ProjectID),
		"instance-id":   print(metadata.InstanceID),
		"get":           get,
		"watch":         watch,
	}

	if !metadata.OnGCE() {
		log.Fatal("This tool only works on Google Compute Engine.")
	}

	if len(os.Args) < 2 {
		printUsage()
	}
	cmd, ok := cmds[os.Args[1]]
	if !ok {
		printUsage()
	}
	if err := cmd(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func print(fn func() (string, error)) cmd {
	return func() error {
		v, err := fn()
		if err != nil {
			return err
		}
		fmt.Println(v)
		return nil
	}
}

func get() error {
	v, err := metadata.Get(arg())
	if err != nil {
		return err
	}
	fmt.Println(v)
	return nil
}

func watch() error {
	return metadata.Subscribe(arg(), func(v string, ok bool) error {
		if !ok {
			return errors.New("Not found")
		}
		fmt.Print(v)
		return nil
	})
}

func arg() string {
	if len(os.Args) < 3 {
		printUsage()
		return ""
	}
	return os.Args[2]
}

func printUsage() {
	fmt.Println(usageText)
	os.Exit(1)
}

const usageText = `gce-metadata <cmd> [args...]

Commands:
- hostname
- external-ip
- internal-ip
- instance-name
- zone
- project-id
- instance-id
- get <attr>
- watch <attr>`
