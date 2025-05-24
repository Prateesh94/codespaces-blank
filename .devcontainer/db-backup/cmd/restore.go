/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"db-backup/backups"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore is a command to restore databases",
	Long: `restore is a command to restore databases.
It supports various DBMS like Postgres, MongoDB, MySQL, and MariaDB.
You can specify the DBMS using the --dbms flag.
Example: db-backup restore --dbms postgres`,
	Run: func(cmd *cobra.Command, args []string) {
		if dbName == "" {
			fmt.Println("Please provide a DBMS name using the --dbms flag.")
			return
		}
		dbName = strings.ToLower(dbName)
		switch dbName {
		case "postgres":
			backups.RestorePostgres(dbName)
		case "mongodb":
			backups.RestoreMongo(dbName)
		case "mysql":
			backups.RestoreMysql(dbName)
		case "mariadb":
			backups.RestoreMariaDB(dbName)
		default:
			fmt.Println("Unsupported DBMS. Please use 'postgres', 'mongodb', or 'mysql'.")

		}
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().StringVarP(&dbName, "dbms", "d", "", "DBMS name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
