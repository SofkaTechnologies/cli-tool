/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"../utils"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// javaCmd represents the initJava command
var nosqlrepoCmd = &cobra.Command{
	Use:   "nosqlrepo",
	Short: "Add a non-sql repository to existent java project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nosqlrepo called")
		//
		//file, ex := filepath.Abs("./templates/java/gradle/settings.gradle")
		//if ex == nil {
		//	fmt.Println("Absolute:", file)
		//}

		_ = viper.BindPFlag("entity", cmd.PersistentFlags().Lookup("entity"))
		_ = viper.BindPFlag("path", cmd.PersistentFlags().Lookup("path"))
		_ = viper.BindPFlag("name", cmd.PersistentFlags().Lookup("name"))
		_ = viper.BindPFlag("group", cmd.PersistentFlags().Lookup("group"))

		projectName := viper.GetString("name")
		entity := viper.GetString("entity")
		group := viper.GetString("group")
		destinationPath := viper.GetString("path")
		tp := utils.ProjectSettings{projectName, group, entity}
		utils.CreateDirectoryStructure(destinationPath, "java", cmd.Name(), tp)
	},
}

func init() {
	rootCmd.AddCommand(nosqlrepoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// javaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	nosqlrepoCmd.PersistentFlags().StringP("entity", "e", "", "Entity's name")
	nosqlrepoCmd.PersistentFlags().StringP("path", "p", "", "Project's export path")
	nosqlrepoCmd.PersistentFlags().StringP("name", "n", "", "Project's name")
	nosqlrepoCmd.PersistentFlags().StringP("group", "g", "", "Project's group")
}
