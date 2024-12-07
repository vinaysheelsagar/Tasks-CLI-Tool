package utilities

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

func CheckNil(err error, fail_message string, success_message string) {
	if err != nil {
		if fail_message != "" {
			fmt.Println(fail_message)
		} else {
			panic(err)
		}
		os.Exit(0)
	} else {
		if success_message != "" {
			fmt.Println(success_message)
		}
	}
}

func GetTaclDir() string {
	homeDir, err := os.UserHomeDir()
	CheckNil(err, "Could not find user's home dir", "")

	taclDir := path.Join(homeDir, ".tacl")
	os.MkdirAll(taclDir, fs.ModePerm)

	return taclDir
}
