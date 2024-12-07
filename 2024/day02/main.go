package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		var reports [][]int
		reports = returnSampleReports()
		fmt.Println(reports)
	*/
	var reports [][]int

	//listSourcefiles := []string{"sample_input2.txt"}
	listSourcefiles := []string{"full_input.txt"}
	for _, fileName := range listSourcefiles {
		fmt.Println("Starting to read", fileName)
		reports = buildReportsFromFile((fileName))
	}

	reportSafetyCounts := make(map[string]int)
	reportSafetyCounts["safe"] = 0
	reportSafetyCounts["unsafe"] = 0

	for _, report := range reports {
		var reportIsSafe bool
		reportIsSafe, _ = getReportSafety(report)
		if reportIsSafe {
			reportSafetyCounts["safe"] += 1
		} else {
			reportSafetyCounts["unsafe"] += 1
		}
	}

	for state, stateCount := range reportSafetyCounts {
		fmt.Printf("%s: %d\n", state, stateCount)
	}

}

func getReportSafety(r []int) (bool, string) {
	var reportIsSafe bool
	var reportSafetyReason string
	var reportItemChange string // ux, dx, nx

	reportIsSafe = true
	reportSafetyReason = ""
	reportItemChange = "nx"

	fmt.Println("Testing report", r)

	for rIdx, rItem := range r {
		// skip the last item in the slice
		if rIdx < len(r)-1 && reportIsSafe {
			// nextItemDiff := int(math.Abs(float64(rItem) - float64(r[rIdx+1])))
			nextItemDiff := rItem - r[rIdx+1]
			switch {
			case nextItemDiff > 0:
				if reportItemChange == "dx" {
					reportIsSafe = false
					reportSafetyReason = "report was decreasing and is now increasing"
					break
				} else {
					reportItemChange = "ux"
				}
			case nextItemDiff < 0:
				if reportItemChange == "ux" {
					reportIsSafe = false
					reportSafetyReason = "report was increasing and is now decreasing"
					break
				} else {
					reportItemChange = "dx"
				}
			case nextItemDiff == 0:
				reportItemChange = "nx"
			}

			if reportIsSafe {
				absNextItemDiff := int(math.Abs(float64(nextItemDiff)))
				switch {
				case absNextItemDiff == 0:
					reportIsSafe = false
					reportSafetyReason = "no change between two report items"
					// reportItemChange = "nx" <-- implied, should already be set
					break
				case absNextItemDiff > 3:
					reportIsSafe = false
					reportSafetyReason = "change greater than 3 between two report items"
					break
				case absNextItemDiff <= 3 && absNextItemDiff >= 1:
					reportIsSafe = true
					reportSafetyReason = "within acceptable parameters"
				default:
					reportIsSafe = false
					reportSafetyReason = "received invalid nextItemDiff"
					break
				}
				fmt.Printf("Checking %d and %d: diff %d, %s\n", rItem, r[rIdx+1], nextItemDiff, reportSafetyReason)
			}
		}
	}

	return reportIsSafe, reportSafetyReason
}

func returnSampleReports() [][]int {
	// var reports [][]int
	buildReports := [][]int{}

	buildReports = append(buildReports, []int{7, 6, 4, 2, 1})
	buildReports = append(buildReports, []int{1, 2, 7, 8, 9})
	buildReports = append(buildReports, []int{9, 7, 6, 2, 1})
	buildReports = append(buildReports, []int{1, 3, 2, 4, 5})
	buildReports = append(buildReports, []int{8, 6, 4, 4, 1})
	buildReports = append(buildReports, []int{1, 3, 6, 7, 9})

	return buildReports
}

func buildReportsFromFile(fName string) [][]int {
	fmt.Println(fName)
	//fListOfInts := make([][]int, 8)
	var fListOfInts [][]int

	fileOpener, err := os.Open(fName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(fileOpener)
	readLine, err := Readln(reader)
	for err == nil {
		fmt.Println(readLine)
		fLine := strings.Fields(readLine)
		var newIntList []int
		newIntList, _ = stringsToIntegers(fLine)
		fListOfInts = append(fListOfInts, newIntList)
		readLine, err = Readln(reader)
	}

	return fListOfInts
}

func stringsToIntegers(lines []string) ([]int, error) {
	// Courtesy of https://www.reddit.com/r/golang/comments/r6qwnb/strings_slices_to_integer_slices_the_most_optimal/
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func Readln(rdr *bufio.Reader) (string, error) {
	// Courtesy of https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = rdr.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
