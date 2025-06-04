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

var (
	dbName string
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup is a command to backup databases",
	Long: `backup is a command to backup databases.
It supports various DBMS like Postgres, MongoDB, MySQL, and MariaDB.
You can specify the DBMS using the --dbms flag.
	Example: db-backup backup --dbms postgres`,
	Run: func(cmd *cobra.Command, args []string) {
		if dbName == "" {
			fmt.Println("Please provide a DBMS name using the --dbms flag.")
			return
		}
		dbName = strings.ToLower(dbName)
		switch dbName {
		case "postgres":
			backups.BackupPostgres(dbName)
		case "mongodb":
			backups.BackupMongo(dbName)
		case "mysql":
			backups.BackupMysql(dbName)
		case "mariadb":
			backups.BackupMariaDB(dbName)
		default:
			fmt.Println("Unsupported DBMS. Please use postgres, mongodb, mysql, or mariadb.")
		}

	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVarP(&dbName, "dbms", "d", "", "DBMS name")
	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
