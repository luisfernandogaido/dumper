package mysql

import (
	"os/exec"
	"fmt"
	"time"
	"path/filepath"
)

func Dump(host, user, pass, db, dir string) (string, error) {
	fileName := fmt.Sprintf("%v-%v.sql", db, time.Now().Format("20060102-150405"))
	fullPath := filepath.Join(dir, fileName)
	cmd := exec.Command(
		"mysqldump",
		fmt.Sprintf("-h%v", host),
		fmt.Sprintf("-u%v", user),
		fmt.Sprintf("-p%v", pass),
		"--routines",
		"--databases",
		db,
		fmt.Sprintf("--result-file=%v", fullPath),
	)
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return fileName, nil
}
