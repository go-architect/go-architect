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
    const pl = localStorage.getItem("projects")
    if(pl != null){
        return JSON.parse(pl)
    }
    return null
}
