package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

func main() {
	fmt.Println("AoC day 1, puzzle 1")

	// acLists := make([][]int, 2)
	var acLists [][]int

	listSourcefiles := []string{"list0.txt", "list1.txt"}
	for _, fileName := range listSourcefiles {
		acLists = append(acLists, buildListsFromFile(fileName))
	}

	if len(acLists[0]) == len(acLists[1]) {
		for _, l := range acLists {
			sort.Ints(l)
			// fmt.Printf("list: (%d) %d\n", len(l), l)
		}
	}

	listDiffs := getListItemDiffs(acLists[0], acLists[1])
	// fmt.Printf("List differences: %d\n", listDiffs)
	fmt.Printf("Sum of list differences: %d\n", sumList(listDiffs))

	listSims := getListSimilarity(acLists[0], acLists[1])
	fmt.Printf("List similarities: %d\n", listSims)
	fmt.Printf("Sum of list similarity: %d\n", sumList(listSims))

}

func buildListsFromFile(fName string) []int {
	var fListOfInts []int
	var fLine int

	fileOpener, err := os.Open(fName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		_, err := fmt.Fscanf(fileOpener, "%d\n", &fLine)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fListOfInts = append(fListOfInts, fLine)
	}

	return fListOfInts
}

func getListItemDiffs(l0 []int, l1 []int) []int {
	//listItemDiffs := make([]int, len(l0))
	var listItemDiffs []int
	if len(l0) == len(l1) {
		//fmt.Println("Lists match length, now checking differences...")
		// go doesn't have Abs for ints, so either cast to float (and back again) or DIY
		for lidx, _ := range l0 {
			//fmt.Printf("Slice index %d has l0 value %d and l1 value %d,", lidx, l0[lidx], l1[lidx])
			//fmt.Printf("Abs of l0 idx is %d\n", int(math.Abs(float64(l0[lidx]))))
			lDiff := int(math.Abs(float64(l0[lidx] - l1[lidx])))
			//fmt.Printf("which have a difference of %d\n", lDiff)
			listItemDiffs = append(listItemDiffs, lDiff)
		}
	}
	return listItemDiffs
}

func sumList(l0 []int) int {
	listSum := 0
	for _, l := range l0 {
		listSum += l
	}
	return listSum
}

func getItemInListCount(item int, list []int) int {
	// This probably does not scale well
	// Assume we can return 0 as a default, in case a number does not appear in the right-side list
	iilCount := 0

	for _, ic := range list {
		if ic == item {
			iilCount += 1
		}
	}

	return iilCount
}

func getListSimilarity(l0 []int, l1 []int) []int {
	var listItemSimilarities []int
	for _, li := range l0 {
		listItemSimilarities = append(listItemSimilarities, li*getItemInListCount(li, l1))
	}

	return listItemSimilarities
}
