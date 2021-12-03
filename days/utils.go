package days

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ArrayToInteger(vs []string) []int {
	vsm := make([]int, len(vs)-1)
	for i, v := range vs {
		if v != "" {
			vsm[i] = ToInteger(v)
		}
	}
	return vsm
}

func ReadLines(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func ToInteger(str string) int {
	var value int
	var err error

	value, err = strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return value
}

func PrintTitle(str string) {
	fmt.Println(str)
	fmt.Println("=============")
	fmt.Println()
}

func PrintSubtitle(str string) {
	fmt.Println(str)
	fmt.Println("-------------")
}
