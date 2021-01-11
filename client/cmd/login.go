/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/hoeniu/ele-chain/client/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login elemchain server",
	Long: `login elemchain server with username and password,
if user could login success, elemchain can create key in $HOME/.client,if next can direct login.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createLogin(cmd *cobra.Command, args []string) {
	if !login {
		Login()
	}
}

func Login() {
	ui := &pkg.UserInfo{}
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Username:")
	if input.Scan() {
		ui.Name = input.Text()
	}
	fmt.Printf("Password:")
	// t := terminal.NewTerminal(os.Stdin, ui.Password)
	// outgoing, err := t.ReadPassword(ui.Password)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	if input.Scan() {
		ui.Password = input.Text()
	}
	ui.CreateTime = time.Now()
	ui.Expire = time.Now().Add(time.Hour * 1000)
	login = true
	viper.WriteConfig()
}
