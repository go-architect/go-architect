import {GetProjectList, GetSelectedProject, RemoveSelectedProject, SaveProjectList, SaveSelectedProject} from "../../wailsjs/go/main/App";
import {project, storage} from "../../wailsjs/go/models";
import ProjectList = storage.ProjectList;
import ProjectInfo = project.ProjectInfo;

export const getSelectedProject = async (): Promise<ProjectInfo | null> => {
    const sp = await GetSelectedProject()
    if (!sp.Selected) {
        return null
    }
    return sp.Project
}

export const saveSelectedProject = async (project: any) => {
    await SaveSelectedProject(project)
    const pl = await getProjectsList()
    if (pl != null) {
        for (let i = 0; i < pl.projects.length; i++) {
            if (pl.projects[i].id === project.id) {
                pl.projects[i].organization_packages = project.organization_packages
            }
        }
    }
    await SaveProjectList(pl)
}

export const removeSelectedProject = async () => {
    await RemoveSelectedProject()
}

export const saveProjectsList = async (projects: any) => {
    console.log("Save Project List: ", projects)
    const projectList = new ProjectList({"projects": projects})
    await SaveProjectList(projectList)
}

export const getProjectsList = async () => {
    const projectList = await GetProjectList()
    console.log("Projects: ", projectList)
    return projectList
}
