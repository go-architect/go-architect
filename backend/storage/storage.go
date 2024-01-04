package storage

import (
	"encoding/json"
	"fmt"
	"github.com/go-architect/go-architect/backend/utils"
	"os"
	"path"
	"path/filepath"
)

func GetProjectList() ProjectList {
	var projectList ProjectList
	dbFile := utils.GetHomeDir() + filepath.FromSlash("/.goarchitect/db/projects.json")
	fmt.Printf("DBFile: %s\n", dbFile)
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

func SaveProjectList(projectList ProjectList) {
	dbFile := utils.GetHomeDir() + filepath.FromSlash("/.goarchitect/db/projects.json")
	jsonData, _ := json.Marshal(projectList)
	os.WriteFile(dbFile, jsonData, 0755)
}

func SaveSelectedProject(project Project) {
	dbFile := utils.GetHomeDir() + filepath.FromSlash("/.goarchitect/db/project.json")
	jsonData, _ := json.Marshal(project)
	os.WriteFile(dbFile, jsonData, 0755)
}

func GetSelectedProject() SelectedProject {
	var project Project
	dbFile := utils.GetHomeDir() + filepath.FromSlash("/.goarchitect/db/project.json")
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

func RemoveSelectedProject() {
	dbFile := utils.GetHomeDir() + filepath.FromSlash("/.goarchitect/db/project.json")
	os.Remove(dbFile)
}

func writeEmptyList(filename string) {
	dir := path.Dir(filename)
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
