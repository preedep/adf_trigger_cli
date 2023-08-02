package cmd

import (
	"adf_trigger_cli/azure"
	"adf_trigger_cli/config"
	"fmt"
	"github.com/spf13/cobra"
)

func Execute() {

	var subscriptionId string
	var resourceGroup string
	var factoryName string
	var pipelineName string
	var recovery bool
	var parameters string

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
			var p *config.Parameters = nil
			var err error = nil
			if len(parameters) > 0 {
				p, err = config.ReadParametersFile(parameters)
				if err != nil {
					panic(err)
				}
			}
			adf := azure.CreateDataFactories(subscriptionId, resourceGroup, factoryName)
			err = adf.RunPipeLine(pipelineName, recovery, p, func(adfStatus azure.ADFStatus, s string) {
				fmt.Printf("Run pipeline status : %v , message : %v\r\n", adfStatus, s)
			})
			if err != nil {
				panic(err)
			}
		},
	}
	runCmd.Flags().StringVarP(&subscriptionId, "subscription_id", "s", "", "Azure Subscription ID [*required]")
	runCmd.Flags().StringVarP(&resourceGroup, "resource_group", "r", "", "Azure Resource Group [*required")
	runCmd.Flags().StringVarP(&factoryName, "factory_name", "f", "", "Azure ADF Factory Name [*required]")
	runCmd.Flags().StringVarP(&pipelineName, "pipeline_name", "p", "", "Azure ADF Pipeline Name [*required]")
	runCmd.Flags().BoolVarP(&recovery, "recovery", "c", false, "Azure ADF Pipeline try support recovery")
	runCmd.Flags().StringVarP(&parameters, "parameter_file", "v", "", "Azure ADF Pipeline parameters")

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
