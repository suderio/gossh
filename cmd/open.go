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
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens a connection to some host",
	Long: `This is the main posh command, the equivalent of
'ssh somehost.somedomain.com', with some niceties:

* if you tagged your server, you may use them to connect
> posh apache prd
* if more than one host is tagged, a menu is shown
> posh apache
| There are two hosts tagged apache:
| 1) apache.domain.com prd
| 2) apache.dsv.domain.com dsv
| Which one? 
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkSshExists()
		goshHost := args[0]
		viper.SetDefault("hosts", map[string]interface{}{
			goshHost: map[string]interface{}{
				"port": goshPort, "user": goshUser, "tags": [0]string{}}})
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.Flags().Int16VarP(&goshPort, "port", "p", 22, "Connection port")
	openCmd.Flags().StringArrayVarP(&goshFlags, "flags", "f", nil, "Connection flags")
}

func checkSshExists() {
	_, err := exec.LookPath("ssh")
	if err != nil {
		fmt.Printf("Cannot find 'ssh' executable.\nGosh will work but no ssh connection will be available.\n")
	}
}
