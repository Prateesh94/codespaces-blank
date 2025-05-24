package backups

import (
	"fmt"
	"os"
	"os/exec"

	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

func BackupMongo(db string) {
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
	filename := FileName(con.DBName)
	filename = filename + ".gz"
	cmd := exec.Command("mongodump", "--host", con.Host, "--port", fmt.Sprint(con.Port), "--username", con.User, "--password", con.Password, "--db", con.DBName, "--authenticationDatabase", "admin", "--gzip", "--archive="+dir+filename)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println("Backup successfully saved to file:", filename)
	}
}

func RestoreMongo(db string) {
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
	cmd := exec.Command("mongorestore", "--host", con.Host, "--port", fmt.Sprint(con.Port), "--username", con.User, "--password", con.Password, "--db", con.DBName, "--authenticationDatabase", "admin", "--gzip", "--archive="+dir+filename)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println("Restore successful from file:", filename)
	}
}
