/*
Copyright © 2021 Paulo Suderio <paulo.suderio@gmail.com>

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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var l, r, s, t bool

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists known connections",
	Long:  `Lists all known hosts, `,
	Run: func(cmd *cobra.Command, args []string) {
		hosts := viper.GetStringMap("hosts")
		for k := range hosts {
			fmt.Print(k)
			if l {
				fmt.Print("\t")
				details := viper.GetStringMap("hosts." + k)
				for _, k := range details {
					fmt.Print(k)
					fmt.Print("\t")
				}
			}
			fmt.Println()
		}

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVarP(&l, "long", "l", false, "Long listing")
	lsCmd.Flags().BoolVarP(&r, "reverse", "r", false, "Reverse order")
	lsCmd.Flags().BoolVarP(&s, "size", "s", false, "Orders by number of connections")
	lsCmd.Flags().BoolVarP(&t, "time", "t", false, "Orders by last connection")
}
