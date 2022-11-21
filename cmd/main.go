package main

import (
	"jatis/pkg/repo"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

func main() {
	repo.Init(repo.Option{
		Username: "Jatis",
		Password: "123456",
		Host:     "localhost",
		Port:     "3306",
	})

	cmd := cobra.Command{
		Use:   "APP",
		Short: "Jatis App",
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:   "run-csv",
			Short: "App to Populate csv data to database",
			Run: func(cmd *cobra.Command, args []string) {
				log.Println("Running Populate data App")
				runCsv()
			},
		},
	)

	cmd.AddCommand(
		&cobra.Command{
			Use:   "run-api",
			Short: "Jatis API",
			Run: func(cmd *cobra.Command, args []string) {
				log.Println("Running Jatis API Apps")
				runAPI(option{
					Addr: ":2000",
				})
			},
		},
	)

	cmd.Execute()
}
