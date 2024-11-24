package utilities

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

func CheckNil(err error, fail_message string, success_message string) {
	if err != nil {
		message := fail_message
		if fail_message != "" {
			message = err.Error()
		}

		panic(message)

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
