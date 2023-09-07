package storage

type ProjectList struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	ID                   string              `json:"id"`
	Name                 string              `json:"name"`
	Path                 string              `json:"path"`
	MainPackage          string              `json:"package"`
	OrganizationPackages []string            `json:"organization_packages"`
	HistoricalMetrics    []HistoricalMetrics `json:"historical_metrics"`
}

type SelectedProject struct {
	Selected bool
	Project  *Project
}

type HistoricalMetrics struct {
	Date    string  `json:"date"`
	Commit  string  `json:"commit"`
	Metrics Metrics `json:"metrics"`
}

type Metrics struct {
	SourceFiles      int `json:"source_files"`      // number of source files
	Structs          int `json:"structs"`           // number of structs
	Interfaces       int `json:"interfaces"`        // number of interfaces
	Functions        int `json:"functions"`         // number of functions
	Methods          int `json:"methods"`           // number of methods
	Variables        int `json:"variables"`         // number of variables
	Constants        int `json:"constants"`         // number of constants
	PublicStructs    int `json:"public_structs"`    // number of public structs
	PublicInterfaces int `json:"public_interfaces"` // number of public interfaces
	PublicFunctions  int `json:"public_functions"`  // number of public functions
	PublicMethods    int `json:"public_methods"`    // number of public methods
	PublicVariables  int `json:"public_variables"`  // number of public variables
	PublicConstants  int `json:"public_constants"`  // number of public constants
	Packages         int `json:"num_packages"`      // number of packages in the module
	TotalLinesOfCode int `json:"total_loc"`         // number of lines of code for the whole project

	InterfaceAverageMethods float64 `json:"interface_avg_size"` // average number of methods declared in all the interfaces
	InterfaceMaxMethods     int     `json:"interface_max_size"` // max number of methods declared in an interfaces
	InterfaceMinMethods     int     `json:"interface_min_size"` // min number of methods declared in an interfaces

	CommentedTotalLines    int     `json:"comments_total_lines"`      // total of comment lines
	CommentsRatio          float64 `json:"comments_lines_ratio"`      // ratio of comment lines (comments_loc / project_loc)
	FilesWithComments      int     `json:"files_with_comments"`       // number of files with comments
	FilesWithCommentsRatio float64 `json:"files_with_comments_ratio"` // ratio of files with comments

	AverageCyclomaticComplexity float64 `json:"cyclomatic_complexity_avg"`
	AverageCognitiveComplexity  float64 `json:"cognitive_complexity_avg"`
}
