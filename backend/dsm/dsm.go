package dsm

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/dsm"
)

func RearrangeDSM(result *dsm.DependencyStructureMatrix) *dsm.DependencyStructureMatrix {
	result, currentRow := rearrangeZeroDepenciesPackages(result)
	result, lastRow := rearrangeZeroDependantsPackages(result, currentRow)
	fmt.Printf("Pending to ReArrange from [%d]->[%d]", currentRow, lastRow)
	return result
}

func rearrangeZeroDependantsPackages(result *dsm.DependencyStructureMatrix, currentRow int) (*dsm.DependencyStructureMatrix, int) {
	lastColumn := len(result.Packages) - 1
	var zeroDependants []int
	dependantsCount := make([]int, len(result.Packages), len(result.Packages))

	for _, r := range result.Dependencies {
		for i, v := range r {
			dependantsCount[i] += v
		}
	}
	for i, d := range dependantsCount {
		if d == 0 && i > currentRow {
			fmt.Printf("Col[%d][%s] has zeroDependants\n", i, result.Packages[i])
			zeroDependants = append(zeroDependants, i)
		}
	}
	currentColumn := len(result.Packages) - len(zeroDependants) - 1
	fmt.Printf("There are [%d] packages with zeroDependants, so then changes will start at %d\n", len(zeroDependants), currentColumn)
	for i, d := range zeroDependants {
		fmt.Printf("Col[%d][%s] should be moved to current column position (%d)\n", d, result.Packages[d-i], lastColumn-len(zeroDependants))
		result = changePositions(result, d-i, lastColumn)
		currentColumn++
		//		lastColumn--
	}

	return result, len(result.Packages) - len(zeroDependants) - 1
}

func rearrangeZeroDepenciesPackages(result *dsm.DependencyStructureMatrix) (*dsm.DependencyStructureMatrix, int) {
	var currentRow int
	var zeroDependencies []int

	for idx, row := range result.Dependencies {
		fmt.Printf("[%d] -> [%+v]\n", idx, row)
		if containsOnlyZeros(row) {
			zeroDependencies = append(zeroDependencies, idx)
		}
	}
	for idx, i := range zeroDependencies {
		if i != currentRow {
			fmt.Printf("[%d] - Change Row/Column (%d) to(%d)\n", idx, i, currentRow)
			result = changePositions(result, i, currentRow)
		}
		currentRow++
	}
	return result, currentRow
}
func changeDependencies(dependencies [][]int, source, destination int) [][]int {
	dependencies = changeDependenciesRows(dependencies, source, destination)
	dependencies = changeDependenciesColumns(dependencies, source, destination)
	return dependencies
}

func changeDependenciesColumns(dependencies [][]int, source int, destination int) [][]int {
	var newDependencies [][]int

	for _, row := range dependencies {
		deps := changeDependants(row, source, destination)
		newDependencies = append(newDependencies, deps)
	}

	return newDependencies
}

func changeDependenciesRows(dependencies [][]int, source int, destination int) [][]int {
	rowToMove := dependencies[source]
	deps := append(dependencies[:source], dependencies[source+1:]...)
	deps = append(deps, []int{})
	copy(deps[destination+1:], deps[destination:])
	deps[destination] = rowToMove

	return deps
}
func changePositions(result *dsm.DependencyStructureMatrix, source, destination int) *dsm.DependencyStructureMatrix {
	result.Packages = changePackages(result.Packages, source, destination)
	result.Dependencies = changeDependencies(result.Dependencies, source, destination)
	return result
}
func changePackages(packages []string, source, destination int) []string {
	packageToMove := packages[source]
	pkgs := append(packages[:source], packages[source+1:]...)
	pkgs = append(pkgs, "")
	copy(pkgs[destination+1:], pkgs[destination:])
	pkgs[destination] = packageToMove

	return pkgs
}
func changeDependants(dependants []int, source, destination int) []int {
	depToMove := dependants[source]
	deps := append(dependants[:source], dependants[source+1:]...)
	deps = append(deps, 0)
	copy(deps[destination+1:], deps[destination:])
	deps[destination] = depToMove

	return deps
}
func containsOnlyZeros(row []int) bool {
	for _, i := range row {
		if i != 0 {
			return false
		}
	}
	return true
}
