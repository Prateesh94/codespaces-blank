package backups

import (
	"fmt"
	"os"
	"os/exec"
)

func BackupPostgres(db string) {

	var con DbConfig
	dir := GetDir(db)
	CheckDir(dir)
	fmt.Print("Enter Host: ")
	fmt.Scanln(&con.Host)
	fmt.Print("Enter Port: ")
	fmt.Scanln(&con.Port)
	fmt.Print("Enter User: ")
	fmt.Scanln(&con.User)
	fmt.Print("Enter Database Name: ")
	fmt.Scanln(&con.DBName)
	filename := FileName("postgres") + ".sql"
	cmd := exec.Command("pg_dump", "-h", con.Host, "-p", fmt.Sprint(con.Port), "-U", con.User, "-W", con.DBName, "-f", dir+filename)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		fmt.Fprintln(os.Stderr, err)

		return
	} else {
		fmt.Println("Backup successfully saved to file:", filename)
	}

}

func RestorePostgres(db string) {
	var con DbConfig
	var filename string
	dir := GetDir(db)
	fmt.Print("Enter Host: ")
	fmt.Scanln(&con.Host)
	fmt.Print("Enter Port: ")
	fmt.Scanln(&con.Port)
	fmt.Print("Enter User: ")
	fmt.Scanln(&con.User)
	fmt.Print("Enter Database Name: ")
	fmt.Scanln(&con.DBName)
	fmt.Print("Enter File Name: ")
	fmt.Scanln(&filename)

	if CheckFileNameExist(filename, dir) {
		fmt.Println("File does not exist, please check the file name and try again.")
		return
	}
	cmd := exec.Command("/usr/bin/psql", "-h", con.Host, "-p", fmt.Sprint(con.Port), "-U", con.User, "-d", con.DBName, "-W", "-f", dir+filename)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+con.Password)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Restore successful from file:", filename)

}
