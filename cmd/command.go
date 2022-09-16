package cmd

import (
	"adf_trigger_cli/azure"
	"fmt"
	"github.com/spf13/cobra"
)

func Execute() {

	var subscription_id string
	var resource_group string
	var factory_name string
	var pipeline_name string

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ADF Trigger CLI",
		Long:  `All software has versions. This is ADF Trigger CLI's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ADF Trigger CLI Tools v1.0 -- HEAD")
		},
	}
	var rootCmd = &cobra.Command{Use: "app",
		Short: "ADF Trigger CLI Tools for remote run ADF pipeline",
		Long: `ADF Trigger CLI Tools for remote run ADF pipeline.
Developed by Mr.Preedee Ponchevin copyright 2022`,
	}

	var runCmd = &cobra.Command{Use: "run [run ADF pipeline]",
		Short: "Run ADF with specific pipeline",
		Long:  `Run ADF with specific pipeline`,
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println(subscription_id)
			//fmt.Println(resource_group)
			//fmt.Println(factory_name)
			//fmt.Println(pipeline_name)
			datafactories := azure.CreateDataFactories(subscription_id, resource_group, factory_name)
			err := datafactories.RunPipeLine(pipeline_name, func(adfStatus azure.ADFStatus, s string) {
				fmt.Printf("Run pipeline status : %v , message : %v\r\n", adfStatus, s)
			})
			if err != nil {
				panic(err)
			}
		},
	}

	runCmd.Flags().StringVarP(&subscription_id, "subscription_id", "s", "", "Azure Subscription ID [*required]")
	runCmd.Flags().StringVarP(&resource_group, "resource_group", "r", "", "Azure Resource Group [*required")
	runCmd.Flags().StringVarP(&factory_name, "factory_name", "f", "", "Azure ADF Factory Name [*required]")
	runCmd.Flags().StringVarP(&pipeline_name, "pipeline_name", "p", "", "Azure ADF Pipeline Name [*required]")

	err := runCmd.MarkFlagRequired("subscription_id")
	if err != nil {
		panic(err)
	}
	err = runCmd.MarkFlagRequired("resource_group")
	if err != nil {
		panic(err)
	}
	err = runCmd.MarkFlagRequired("factory_name")
	if err != nil {
		panic(err)
	}
	err = runCmd.MarkFlagRequired("pipeline_name")
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runCmd)

	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
