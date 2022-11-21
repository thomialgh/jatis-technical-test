package main

import (
	"jatis/pkg/config"
	"jatis/pkg/repo"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

func main() {

	config.ReadConfig()

	repo.Init(repo.Option{
		Username: config.Cfg.Mysql.Username,
		Password: config.Cfg.Mysql.Password,
		Host:     config.Cfg.Mysql.Host,
		Port:     config.Cfg.Mysql.Port,
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
