// Package solutions provides ...
package solutions

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	_ "reflect"
	"regexp"
	"strconv"
	"strings"
)

// Have the function BinarySearchTreeLCA(strArr) take the array of strings
// stored in strArr, which will contain 3 elements: the first element will
// be a binary search tree with all unique values in a preorder traversal
// array, the second and third elements will be two different values, and
// your goal is to find the lowest common ancestor of these two values. For
// example: if strArr is ["[10, 5, 1, 7, 40, 50]", "1", "7"] then this tree
// looks like the following:

//      10

//   5       40

// 1   7       50

// For the input above, your program should return 5 because that is the
// value of the node that is the LCA of the two nodes with values 1 and 7.
// You can assume the two nodes you are searching for in the tree will
// exist somewhere in the tree.

func BinarySearchTreeLCA(strArr string) string {
	splits := strings.Split(strArr, ", ")
	// PARSE NODE TREE - nasty stringified array
	// remove quotes
	s0, err := strconv.Unquote(splits[0])
	fmt.Println(s0)
	if err != nil {
		log.Fatal(err)
	}
	// remove brackets, use 2nd subgroup
	smatch := regexp.MustCompile(`(?:\[)(.*)(?:\])`).FindStringSubmatch(s0)
	// fmt.Println(smatch[1])
	// parse CSV
	r := csv.NewReader(strings.NewReader(smatch[1]))
	var sArr []string
	for {
		r, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		sArr = r
	}
	// fmt.Printf("sArr %v\n", sArr[0])

	// convert string array values to ints
	nArray := make([]int, len(sArr))
	for i := 0; i < len(sArr); i++ {
		num, err := strconv.Atoi(sArr[i])
		if err != nil {
			log.Fatal(err)
		}
		nArray[i] = num
	}

	// PARSE SEARCH ARG 1 TO INT
	s1, err := strconv.Unquote(splits[1])
	if err != nil {
		log.Fatal(err)
	}
	n1, err := strconv.Atoi(s1)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("n1 %v\n", n1)

	// PARSE SEARCH ARG 2 TO INT
	s2, err := strconv.Unquote(splits[2])
	if err != nil {
		log.Fatal(err)
	}
	n2, err := strconv.Atoi(s2)
	// fmt.Printf("n2 %v\n", n2)

	// FIND INDEXES OF ARGS IN NODE TREE
	var idx1 int
	var idx2 int
	for i, v := range nArray {
		if v == n1 {
			idx1 = i
		}
		if v == n2 {
			idx2 = i
		}
	}
	fmt.Printf("idx1 %v, idx2 %v\n", idx1, idx2)

	// determine farthest of two index positions
	// set to min max of nodetree index value for clarity
	var rEdgeIdx int
	var min int
	var max int
	if idx1 > idx2 {
		rEdgeIdx = idx1
		min = nArray[idx2]
		max = nArray[idx1]
	} else {
		rEdgeIdx = idx2
		min = nArray[idx1]
		max = nArray[idx2]
	}
	fmt.Printf("rEdgeIdx %v, min %v, max %v\n", rEdgeIdx, min, max)

	// iterate over tree to the left of rEdgeIdx
	// check to see if any two tree values that split n1 and n2
	// if yes, thats the answer
	var split int
	// set to -1 since split could be at array[0]
	split = -1
	for i := 0; i <= rEdgeIdx; i++ {
		val := nArray[i]
		fmt.Printf("val %v\n", val)
		if val >= min && val <= max {
			split = val
			fmt.Printf("split %v\n", split)
			break
		}
	}

	// if no split, return num farthest to left
	var result int
	if split == -1 {
		result = min
	} else {
		result = split
	}

	// return node tree value at index as string
	fmt.Printf("nArray result %v\n", result)
	ans := strconv.Itoa(result)

	// fmt.Printf("nArray result %v, ans %v\n", result, ans)
	return ans
}
