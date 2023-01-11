package gocyclo

import (
	"bufio"
	"strconv"
	"strings"
)

func MapToGoCycloModel(cmdOutput string) *GoCycloOutput {
	var details []*ComplexityDetail
	var average float64

	scanner := bufio.NewScanner(strings.NewReader(cmdOutput))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Average") {
			average = mapAverageLine(line)
		} else {
			details = append(details, mapLine(line))
		}
	}

	return &GoCycloOutput{
		AverageComplexity: average,
		ComplexityDetails: details,
	}
}

func mapAverageLine(line string) float64 {
	data := strings.Split(line, " ")
	cx, _ := strconv.ParseFloat(data[1], 64)
	return cx
}

func mapLine(line string) *ComplexityDetail {
	data := strings.Split(line, " ")
	cx, _ := strconv.Atoi(data[0])

	return &ComplexityDetail{
		Complexity: cx,
		Package:    data[1],
		File:       resolveFileDetails(data[3]),
		Function:   data[2],
	}
}

func resolveFileDetails(input string) string {
	slashIndex := strings.LastIndex(input, "/") + 1
	return input[slashIndex:]
}
