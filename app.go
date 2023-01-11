package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenDirectoryDialog() string {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Go Project Folder",
	})
	if err != nil {
		runtime.LogErrorf(a.ctx, "Error: %+v\n", err)
	}
	return path
}
