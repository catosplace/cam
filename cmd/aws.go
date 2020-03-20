/* Copyright © 2020 Catosplace <engineering@catosplace.co.nz>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package cmd

import (
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/catosplace/cam/internal/box"
)

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Initialise AWS Terraform state storage",
	Long:  `Initialise an AWS S3 Bucket to store Terraform state`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("aws called")

		if viper.IsSet("aws.account_id") {
			log.Println(viper.GetString("aws.account_id"))
		}

		tfFiles := []string{"/terragrunt.hcl", "/main.tf"}
		createTerraformFiles(tfFiles)

		osCmd :=
			exec.Command("terragrunt", "init",
				"--terragrunt-config", "/home/cato/Code/cam/.cam/terragrunt.hcl",
				"--terragrunt-working-dir", "/home/cato/Code/cam/.cam")
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
		err := osCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func createTerraformFiles(tfFiles []string) {
	if _, err := os.Stat(".cam"); os.IsNotExist(err) {
		err = os.Mkdir(".cam", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, file := range tfFiles {
		tpl := template.Must(template.New("").Parse(string(box.Get(file))))
		f, err := os.Create(".cam" + file)
		if err != nil {
			log.Fatal(err)
		}
		err = tpl.Execute(f, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	rootCmd.AddCommand(awsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// awsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// awsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
