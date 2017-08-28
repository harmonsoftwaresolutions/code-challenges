// Package solutions provides ...
package solutions

import (
	"encoding/csv"
	_ "fmt"
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

func BinarySearchTreeLCA2(strArr []string) string {

	smatch := regexp.MustCompile(`(?:\[)(.*)(?:\])`).FindStringSubmatch(strArr[0])

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

	nArray := make([]int, len(sArr))
	for i := 0; i < len(sArr); i++ {
		num, err := strconv.Atoi(sArr[i])
		if err != nil {
			log.Fatal(err)
		}
		nArray[i] = num
	}

	n1, err := strconv.Atoi(strArr[1])
	if err != nil {
		log.Fatal(err)
	}

	n2, err := strconv.Atoi(strArr[2])
	if err != nil {
		log.Fatal(err)
	}

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

	var split int
	split = -1
	for i := 0; i <= rEdgeIdx; i++ {
		val := nArray[i]
		if val >= min && val <= max {
			split = val
			break
		}
	}

	var result int
	if split == -1 {
		result = min
	} else {
		result = split
	}

	ans := strconv.Itoa(result)
	return ans
}
