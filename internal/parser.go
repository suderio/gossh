/*
Copyright © 2021 Paulo Suderio

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kevinburke/ssh_config"
)

func Parse() {
	f, _ := os.OpenFile(filepath.Join(os.Getenv("HOME"), ".ssh", "config"), os.O_RDWR|os.O_CREATE, 0755)
	cfg, _ := ssh_config.Decode(f)
	for _, host := range cfg.Hosts {
		log.Println("hosts:", host.String())
		log.Println("patterns:", host.Patterns)
		for _, node := range host.Nodes {
			// Manipulate the nodes as you see fit, or use a type switch to
			// distinguish between Empty, KV, and Include nodes.
			log.Println("--", node.String())
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func getConfig(file string) *ssh_config.Config {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("File %s not Found, Creating...", file)
		f = createConfig(file)
	}
	// TODO add 'Include .config.d/gossh' in .ssh/config
	cfg, err := ssh_config.Decode(f)
	if err != nil {
		log.Fatalf("File %s is not a valid ssh_config file, cannot continue.", file)
	}
	return cfg
}

func createConfig(file string) *os.File {
	os.Mkdir(filepath.Dir(file), 0755)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		log.Fatalf("Unable to create config file %s.\n%s", file, err)
	}

	return f
}

func SetUpConfig() *ssh_config.Config {
	cfgPath := filepath.Join(os.Getenv("HOME"), ".ssh", "config.d", "gossh")
	cfg := getConfig(cfgPath)
	for _, host := range cfg.Hosts {
		fmt.Printf("patterns: %v\n", host.Patterns)
		for _, node := range host.Nodes {
			// Manipulate the nodes as you see fit, or use a type switch to
			// distinguish between Empty, KV, and Include nodes.
			println("--", node.String())
		}
	}
	return cfg
}
