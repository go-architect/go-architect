package main

import (
	"context"
	"fmt"
	"github.com/go-architect/go-architect/backend/storage"
	"github.com/go-architect/go-architect/backend/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	exec "golang.org/x/sys/execabs"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	ctx       context.Context
	goVersion string
}

func NewApp() *App {
	return &App{
		goVersion: "undefined",
	}
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

func (a *App) CheckEnvironmentPath() bool {
	envFile := utils.GetHomeDir() + filepath.FromSlash("/.goarchitect/environment")
	runtime.LogInfof(a.ctx, "CheckEnvironmentPath - EnvFile: %s", envFile)

	if doesFileExist(envFile) {
		runtime.LogInfof(a.ctx, "CheckEnvironmentPath - EnvFile exists")
		b, err := os.ReadFile(envFile)
		if err != nil {
			fmt.Print(err)
		} else {
			customPath := strings.TrimSuffix(string(b), "\n")
			os.Setenv("PATH", customPath+":"+os.Getenv("PATH"))
			runtime.LogInfof(a.ctx, "CheckEnvironmentPath - NewPath: %s", os.Getenv("PATH"))
		}
	}

	sb := new(strings.Builder)
	cmd := exec.Command("go", "version")
	cmd.Stdout = sb

	err := cmd.Run()
	if err != nil {
		runtime.LogErrorf(a.ctx, "getPackages: CmdError %+v", err)
		return false
	}

	a.goVersion = sb.String()
	runtime.LogInfof(a.ctx, "CheckEnvironmentPath - GoVersion: %s", a.goVersion)

	return true
}

func (a *App) GetProjectList() storage.ProjectList {
	return storage.GetProjectList()
}

func (a *App) SaveProjectList(projects storage.ProjectList) {
	storage.SaveProjectList(projects)
}

func (a *App) SaveSelectedProject(project storage.Project) {
	storage.SaveSelectedProject(project)
}

func (a *App) RemoveSelectedProject() {
	storage.RemoveSelectedProject()
}

func (a *App) GetSelectedProject() storage.SelectedProject {
	return storage.GetSelectedProject()
}

func (a *App) GetGoVersion() string {
	return a.goVersion
}

func doesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)

	if os.IsNotExist(error) {
		return false
	}
	return true
}
