package backups

import (
	"fmt"
	"os"
	"os/exec"
)

func BackupMariaDB(db string) {
	var con DbConfig
	dir := GetDir(db)
	CheckDir(dir)
	fmt.Print("Enter Host: ")
	fmt.Scanln(&con.Host)
	fmt.Print("Enter Port: ")
	fmt.Scanln(&con.Port)
	fmt.Print("Enter User: ")
	fmt.Scanln(&con.User)
	fmt.Print("Enter Password: ")
	fmt.Scanln(&con.Password)
	fmt.Print("Enter Database Name: ")
	fmt.Scanln(&con.DBName)
	filename := FileName(con.DBName) + ".sql"

	cmd := exec.Command("mariadb-dump", "--skip-ssl", "--no-tablespaces", "-h", con.Host, "-P", fmt.Sprint(con.Port), "-u", con.User, "-p"+con.Password, con.DBName, fmt.Sprintf("--result-file=%s", dir+filename))

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		fmt.Fprintln(os.Stderr, err)

		return
	} else {
		fmt.Println("Backup successfully saved to file:", filename)
	}
}

func RestoreMariaDB(db string) {
	var con DbConfig
	var filename string
	dir := GetDir(db)
	fmt.Print("Enter Host: ")
	fmt.Scanln(&con.Host)
	fmt.Print("Enter Port: ")
	fmt.Scanln(&con.Port)
	fmt.Print("Enter User: ")
	fmt.Scanln(&con.User)
	fmt.Print("Enter Password: ")
	fmt.Scanln(&con.Password)
	fmt.Print("Enter Database Name: ")
	fmt.Scanln(&con.DBName)
	fmt.Print("Enter File Name: ")
	fmt.Scanln(&filename)
	if CheckFileNameExist(filename, dir) {
		fmt.Println("File does not exist, please check the file name and try again.")
		return
	}

	cmd := exec.Command("mariadb", "--skip-ssl", "-h", con.Host, "-P", fmt.Sprint(con.Port), "-u", con.User, "-p"+con.Password, con.DBName, "-e", "source "+dir+filename)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println("Restore successful from file:", filename)
	}
}
