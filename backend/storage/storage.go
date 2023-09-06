package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GetProjectList() ProjectList {
	var projectList ProjectList
	dbFile := os.Getenv("HOME") + filepath.FromSlash("/.goarchitect/db/projects.json")
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

	fmt.Printf("projectList: %+v\n", projectList)
	return projectList
}

func SaveProjectList(projectList ProjectList) {
	dbFile := os.Getenv("HOME") + filepath.FromSlash("/.goarchitect/db/projects.json")
	jsonData, _ := json.Marshal(projectList)
	os.WriteFile(dbFile, jsonData, 0755)
}

func SaveSelectedProject(project Project) {
	dbFile := os.Getenv("HOME") + filepath.FromSlash("/.goarchitect/db/project.json")
	jsonData, _ := json.Marshal(project)
	os.WriteFile(dbFile, jsonData, 0755)
}

func GetSelectedProject() SelectedProject {
	var project Project
	dbFile := os.Getenv("HOME") + filepath.FromSlash("/.goarchitect/db/project.json")
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
	dbFile := os.Getenv("HOME") + filepath.FromSlash("/.goarchitect/db/project.json")
	os.Remove(dbFile)
}

func writeEmptyList(filepath string) {
	var emptyList ProjectList
	jsonData, _ := json.Marshal(emptyList)
	os.WriteFile(filepath, jsonData, 0755)
}

/*
export const getSelectedProject = () => {
const ps = localStorage.getItem("project")
if(ps != null){
return JSON.parse(ps)
}
return null
}

export const saveSelectedProject = (project: any) => {
localStorage.setItem("project", JSON.stringify(project))
const pl = getProjectsList()
if(pl!=null){
for(let i=0;i<pl.length;i++){
if(pl[i].id === project.id){
pl[i].organization_packages = project.organization_packages
}
}
}
localStorage.setItem("projects", JSON.stringify(pl))
}

export const removeSelectedProject = () => {
localStorage.removeItem("project")
}

export const getProjectsList = () => {
openDB()
const pl = localStorage.getItem("projects")
if(pl != null){
return JSON.parse(pl)
}
return []
}
*/
