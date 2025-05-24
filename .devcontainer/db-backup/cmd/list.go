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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list is a command to list backups for a specific DBMS",
	Long: `list is a command to list backups for a specific DBMS.
It supports various DBMS like Postgres, MongoDB, MySQL, and MariaDB.
You can specify the DBMS using the --dbms flag.
Example: db-backup list --dbms postgres`,
	Run: func(cmd *cobra.Command, args []string) {
		dbName = strings.ToLower(dbName)
		if dbName == "" {
			fmt.Println("Please provide a DBMS name using the --dbms flag.")
			return
		}
		if dbName == "postgres" || dbName == "mongodb" || dbName == "mysql" || dbName == "mariadb" {
			backups.ListBackups(dbName)
		} else {
			fmt.Println("Unsupported DBMS. Please use 'postgres', 'mongodb', 'mysql', or 'mariadb'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&dbName, "dbms", "d", "", "DBMS name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
