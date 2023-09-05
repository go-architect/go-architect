package main

import (
	"context"
	"fmt"
	"github.com/fdaines/go-architect-lib/coupling"
	"github.com/fdaines/go-architect-lib/dependency"
	"github.com/fdaines/go-architect-lib/dsm"
	"github.com/fdaines/go-architect-lib/instability"
	"github.com/fdaines/go-architect-lib/metrics/comments"
	"github.com/fdaines/go-architect-lib/metrics/interfaces"
	"github.com/fdaines/go-architect-lib/metrics/loc"
	"github.com/fdaines/go-architect-lib/metrics/types"
	"github.com/fdaines/go-architect-lib/project"
	"github.com/fdaines/go-architect-lib/repository"
	"github.com/fdaines/go-architect/backend/git"
	"github.com/fdaines/go-architect/backend/gocyclo"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	exec "golang.org/x/sys/execabs"
	"time"
)

type Api struct {
	ctx context.Context
}

func NewApi() *Api {
	return &Api{}
}

func (a *Api) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *Api) GetProjectInfo(path string) *project.ProjectInfo {
	projectInfo, err := project.LoadProjectInfo(path)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetProjectInfo - Error: %s\n", err.Error())
		return nil
	}
	return projectInfo
}

func (a *Api) GetRepositoryInfo(path string) *repository.RepositoryInfo {
	repositoryInfo, err := repository.LoadRepositoryInfo(path)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetRepositoryInfo - Error: %s\n", err.Error())
		return nil
	}
	return repositoryInfo
}

func (a *Api) GetMetricsLOC(project *project.ProjectInfo) *loc.ProjectLOC {
	projectLOC, err := loc.CountLoc(project)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetMetricsLOC - Error: %s\n", err.Error())
		return nil
	}
	return projectLOC
}

func (a *Api) GetMetricsInterfaces(project *project.ProjectInfo) *interfaces.InterfaceMetrics {
	projectInterfaces, err := interfaces.ResolveInterfaceMetrics(project)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetMetricsInterfaces - Error: %s\n", err.Error())
		return nil
	}
	return projectInterfaces
}

func (a *Api) GetMetricsComments(project *project.ProjectInfo) *comments.CommentsMetrics {
	projectComments, err := comments.ResolveCommentsMetrics(project)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetMetricsComments - Error: %s\n", err.Error())
		return nil
	}
	return projectComments
}

func (a *Api) GetMetricsTypes(project *project.ProjectInfo) *types.ProjectTypes {
	projectTypes, err := types.ResolveProjectTypes(project)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetMetricsTypes - Error: %s\n", err.Error())
		return nil
	}
	return projectTypes
}

func (a *Api) GetInstability(project *project.ProjectInfo) []*instability.PackageInstability {
	result, err := instability.GetInstability(project)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetInstability - Error: %s\n", err.Error())
		return nil
	}
	return result
}

func (a *Api) GetDSM(project *project.ProjectInfo) *dsm.DependencyStructureMatrix {
	result, err := dsm.GetDependencyStructureMatrix(project)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetDSM - Error: %s\n", err.Error())
		return nil
	}
	return result
}

func (a *Api) GetDependencyGraph(project *project.ProjectInfo) *dependency.ModuleDependencyGraph {
	result, err := dependency.GetDependencyGraph(project, "ALL")
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetDependencyGraph - Error: %s\n", err.Error())
		return nil
	}
	return result
}

func (a *Api) GetDependencyCoupling(project *project.ProjectInfo, dependency string) *coupling.DependencyCoupling {
	result, err := coupling.CalculateCoupling(project, dependency)
	if err != nil {
		runtime.LogErrorf(a.ctx, "GetDependencyCoupling - Error: %s\n", err.Error())
		return nil
	}
	return result
}

func (a *Api) GetMetricsComplexity(project *project.ProjectInfo) (*gocyclo.GoCycloOutput, error) {
	bytes, err := exec.Command("gocyclo", "-top", "5", "-avg", project.Path).Output()
	if err != nil {
		errorMessage := "cannot use GoCyclo Tool, please check module documentation at 'https://github.com/fzipp/gocyclo'"
		runtime.LogErrorf(a.ctx, "%s. Details: %s\n", errorMessage, err.Error())
		return nil, fmt.Errorf(errorMessage)
	}
	return gocyclo.MapToGoCycloModel(string(bytes)), nil
}

func (a *Api) GetMetricsCognitiveComplexity(project *project.ProjectInfo) (*gocyclo.GoCycloOutput, error) {
	bytes, err := exec.Command("gocognit", "-top", "5", "-avg", project.Path).Output()
	if err != nil {
		errorMessage := "cannot use GoCognit Tool, please check module documentation at 'https://github.com/uudashr/gocognit'"
		runtime.LogErrorf(a.ctx, "%s. details: %s\n", errorMessage, err.Error())
		return nil, fmt.Errorf(errorMessage)
	}
	return gocyclo.MapToGoCycloModel(string(bytes)), nil
}

func (a *Api) GetVCSAnalysisInfo(project *project.ProjectInfo, monthsFromNow int) *git.VCSAnalysisInfo {
	since := time.Now().AddDate(0, -1*monthsFromNow, 0)
	r, _ := git.LoadRepositoryTeamCohesion(project.Path, &since)
	return r
}
