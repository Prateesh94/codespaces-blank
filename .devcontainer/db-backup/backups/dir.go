package backups

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"
)

func CheckDir(dir string) {
	// Check if the directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
}

func GetDir(db string) string {
	switch db {
	case "postgres":
		return "/data/postgres/"
	case "mongodb":
		return "/data/mongodb/"
	case "mysql":
		return "/data/mysql/"
	case "mariadb":
		return "/data/mariadb/"
	}
	return ""
}

func ListBackups(db string) {
	dir := GetDir(db)
	list := [][]string{{"Name", "Size", "Modified Time"}}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {

			return err
		}
		fileInfo, err := d.Info()
		if err != nil {

			return err
		}
		if fileInfo.IsDir() {
			return nil // Skip directories
		}
		list = append(list, []string{fileInfo.Name(), fmt.Sprintf("%d bytes", fileInfo.Size()), fileInfo.ModTime().Format("2006/01/02 15:04:05")})
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, item := range list {
		fmt.Fprintf(w, "%s\t%s\t%s\n", item[0], item[1], item[2])
	}
	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}
}
