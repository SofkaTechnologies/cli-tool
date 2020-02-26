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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initJavaCmd represents the initJava command
var initJavaCmd = &cobra.Command{
	Use:   "initJava",
	Short: "Java base project template",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initJava called")
		fmt.Println(viper.Get("type"))
		//
		//file, ex := filepath.Abs("./templates/java/gradle/settings.gradle")
		//if ex == nil {
		//	fmt.Println("Absolute:", file)
		//}
		projectName := viper.GetString("name")
		projectType := viper.GetString("type")
		group := viper.GetString("group")
		destinationPath := viper.GetString("path")
		CreateDestinationStructure(projectName, destinationPath, group, "java", projectType)
	},
}

func init() {
	rootCmd.AddCommand(initJavaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initJavaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initJavaCmd.PersistentFlags().StringP("type", "t", "gradle", "Select between Gradle or Maven")
	initJavaCmd.PersistentFlags().StringP("path", "p", "", "Project's export path")
	initJavaCmd.PersistentFlags().StringP("name", "n", "", "Project's name")
	initJavaCmd.PersistentFlags().StringP("group", "g", "", "Project's group")

	_ = viper.BindPFlag("type", initJavaCmd.PersistentFlags().Lookup("type"))
	_ = viper.BindPFlag("path", initJavaCmd.PersistentFlags().Lookup("path"))
	_ = viper.BindPFlag("name", initJavaCmd.PersistentFlags().Lookup("name"))
	_ = viper.BindPFlag("group", initJavaCmd.PersistentFlags().Lookup("group"))
}
