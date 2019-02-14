// Copyright © 2019 Cornelius Weig <cornelius.weig@tngtech.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

const completionLongDescription = `
	Outputs shell completion for the given shell (bash or zsh)

	OS X:
		$ source $(brew --prefix)/etc/bash_completion
		$ ketall completion bash > ~/.ketall-completion  # for bash users
		$ ketall completion zsh > ~/.ketall-completion   # for zsh users
		$ source ~/.ketall-completion
	Ubuntu:
		$ source /etc/bash-completion
		$ source <(ketall completion bash) # for bash users
		$ source <(ketall completion zsh)  # for zsh users

	Additionally, you may want to output the completion to a file and source in your .bashrc
`

var completionCmd = &cobra.Command{
	Use:       "completion SHELL",
	Short:     "Output shell completion for the given shell (bash or zsh)",
	Long:      completionLongDescription,
	ValidArgs: []string{"bash", "zsh"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("requires 1 arg, found %d", len(args))
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "bash":
			err = rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			err = rootCmd.GenZshCompletion(os.Stdout)
		}
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
