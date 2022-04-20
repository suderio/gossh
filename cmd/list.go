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
package cmd

import (
	"fmt"
	"strings"

	"github.com/kevinburke/ssh_config"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists all hosts",
	Long:    `Used for a comprehensive list of all hosts, with all details.`,
	Example: `gossh ls
	alias		host		user	port	details
	host1		host1.com	xpto	22
	host2		host2.com	user	2222	StrictHostKeyChecking=no`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hosts")
		for _, host := range cfg.Hosts {
			all := []string(nil)
			for _, node := range host.Nodes {
				switch t := node.(type) {
				case *ssh_config.Empty:
					continue
				case *ssh_config.KV:
					lkey := strings.ToLower(t.Key)
					if lkey == "match" {
						continue // can't handle
					}
					all = append(all, t.Value)
				case *ssh_config.Include:
					continue //won't handle
				default:
					continue //don't know, don't care
				}
			}
			fmt.Printf("%s\t%s\n", host.Patterns[0], all)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
