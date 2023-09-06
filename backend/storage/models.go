package storage

type ProjectList struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Path                 string   `json:"path"`
	MainPackage          string   `json:"package"`
	OrganizationPackages []string `json:"organization_packages"`
}

type SelectedProject struct {
	Selected bool
	Project  *Project
}
