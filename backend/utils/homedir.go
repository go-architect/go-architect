package utils

import (
	"fmt"
	"os"
)

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error trying to detect User HomeDir: %+v\n", err)
		return ""
	}
	return homeDir
}
