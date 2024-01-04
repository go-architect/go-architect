package storage

import (
	"context"
	"encoding/json"
	"github.com/go-architect/go-architect/backend/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path/filepath"
)

func GetProjectList(ctx context.Context) ProjectList {
	var projectList ProjectList
	dbFile := utils.GetHomeDir(ctx) + filepath.FromSlash("/.goarchitect/db/projects.json")
	runtime.LogInfof(ctx, "DBFile: %s\n", dbFile)
	fileContent, err := os.ReadFile(dbFile)
	if err != nil {
		writeEmptyList(dbFile)
		return ProjectList{}
	}
	err = json.Unmarshal(fileContent, &projectList)
	if err != nil {
		writeEmptyList(dbFile)
		return ProjectList{}
	}

	return projectList
}

func SaveProjectList(ctx context.Context, projectList ProjectList) {
	dbFile := utils.GetHomeDir(ctx) + filepath.FromSlash("/.goarchitect/db/projects.json")
	jsonData, _ := json.Marshal(projectList)
	os.WriteFile(dbFile, jsonData, 0755)
}

func SaveSelectedProject(ctx context.Context, project Project) {
	dbFile := utils.GetHomeDir(ctx) + filepath.FromSlash("/.goarchitect/db/project.json")
	jsonData, _ := json.Marshal(project)
	os.WriteFile(dbFile, jsonData, 0755)
}

func GetSelectedProject(ctx context.Context) SelectedProject {
	var project Project
	dbFile := utils.GetHomeDir(ctx) + filepath.FromSlash("/.goarchitect/db/project.json")
	fileContent, err := os.ReadFile(dbFile)
	if err != nil {
		return SelectedProject{Selected: false}
	}
	err = json.Unmarshal(fileContent, &project)
	if err != nil {
		return SelectedProject{Selected: false}
	}

	return SelectedProject{Selected: true, Project: &project}
}

func RemoveSelectedProject(ctx context.Context) {
	dbFile := utils.GetHomeDir(ctx) + filepath.FromSlash("/.goarchitect/db/project.json")
	os.Remove(dbFile)
}

func writeEmptyList(filename string) {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	var emptyList ProjectList
	emptyList.Projects = []Project{}
	jsonData, _ := json.Marshal(emptyList)
	os.WriteFile(filename, jsonData, 0755)
}
