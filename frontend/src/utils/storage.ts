export const getSelectedProject = () => {
    const ps = localStorage.getItem("project")
    if(ps != null){
        return JSON.parse(ps)
    }
    return null
}

export const saveSelectedProject = (project: any) => {
    localStorage.setItem("project", JSON.stringify(project))
}

export const removeSelectedProject = () => {
    localStorage.removeItem("project")
}

export const getProjectsList = () => {
    const pl = localStorage.getItem("projects")
    if(pl != null){
        return JSON.parse(pl)
    }
    return null
}
