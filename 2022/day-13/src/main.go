package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

var floatType = "float64"

func AreEqual(a any, b any) float64 {
	typeAItem := typeof(a)
	typeBItem := typeof(b)
	typeANum := typeAItem == floatType
	typeBNum := typeBItem == floatType
	var aInput any
	var bInput any

	if typeANum && typeBNum {
		return a.(float64) - b.(float64)
	}

	if typeANum {
		aInput = []any{a.(float64)}
	} else {
		aInput = a.([]any)
	}

	if typeBNum {
		bInput = []any{b.(float64)}
	} else {
		bInput = b.([]any)
	}

	for i := range aInput.([]any) {
		if len(bInput.([]any)) <= i {
			return 1
		}

		if res := AreEqual(aInput.([]any)[i], bInput.([]any)[i]); res != 0 {
			return res
		}
	}
	if len(aInput.([]any)) == len(bInput.([]any)) {
		return 0
	}

	return -1
}

func partOne(input string) int {
	data := []any{}
	lines := strings.Split(input, "\n\n")
	for _, line := range lines {
		pairs := strings.Split(line, "\n")
		pair_ := []any{}
		for _, pair := range pairs {
			var unmarshalled []any
			json.Unmarshal([]byte(pair), &unmarshalled)
			pair_ = append(pair_, unmarshalled)
		}
		data = append(data, pair_)
	}

	sum := 0
	rightOrder := []int{}
	for i, pair := range data {
		if r := AreEqual(pair.([]any)[0].([]any), pair.([]any)[1].([]any)); r <= 0 {
			rightOrder = append(rightOrder, i)
		}
	}

	for _, index := range rightOrder {
		sum += index + 1
	}

	return sum
}

func partTwo(input string) int {
	data := []any{}
	lines := strings.Split(input, "\n\n")
	for _, line := range lines {
		pairs := strings.Split(line, "\n")
		for _, pair := range pairs {
			var unmarshalled []any
			json.Unmarshal([]byte(pair), &unmarshalled)
			data = append(data, unmarshalled)
		}
	}

	var add1 any
	var add2 any
	json.Unmarshal([]byte("[[2]]"), &add1)
	json.Unmarshal([]byte("[[6]]"), &add2)
	data = append(data, add1, add2)

	sort.SliceStable(data, func(i, j int) bool {
		return AreEqual(data[i], data[j]) < 0
	})

	product := 1
	for key, val := range data {
		str, _ := json.Marshal(val)
		if string(str) == "[[2]]" || string(str) == "[[6]]" {
			product *= key + 1
		}
	}

	return product
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
