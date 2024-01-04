package utils

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

func GetHomeDir(ctx context.Context) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logError(ctx, err)
		return ""
	}
	return homeDir
}

func logError(ctx context.Context, err error) {
	if ctx == nil {
		runtime.LogInfof(ctx, "Error trying to detect User HomeDir: %+v", err)
	} else {
		fmt.Printf("Error trying to detect User HomeDir: %+v\n", err)
	}
}
