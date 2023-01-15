package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Aggregate struct {
	Name   string
	Phones []string
}

type AggregateMap map[string]*Aggregate

func (a Aggregate) HasPhone(checking string) bool {
	for _, p := range a.Phones {
		if p == checking {
			return true
		}
	}
	return false
}

func (a *Aggregate) AddPhone(phoneToAdd string) {
	if len(a.Phones) < 5 && !a.HasPhone(phoneToAdd) {
		a.Phones = append(a.Phones, phoneToAdd)
	}
}

func (col *AggregateMap) Add(toAdd Aggregate) {
	var name = toAdd.Name
	if _, ok := (*col)[name]; ok {
		(*col)[name].AddPhone(toAdd.Phones[0])
		return
	}
	(*col)[name] = &toAdd
}

//func (col AggregateMap) Len() int { return len(col) }
//func (col AggregateMap) Less(i, j string) bool {
//	return col[i].Name < col[j].Name
//}
//func (col *AggregateMap) Swap(i, j string) {
//	(*col)[i], (*col)[j] = (*col)[j], (*col)[i]
//}

func main() {
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for testI := 0; testI < testCount; testI++ {
		var col = AggregateMap{}
		var userCount int
		fmt.Fscanf(in, "%d\n", &userCount)

		var inputName, inputPhone string
		var inputAggregates = make([]Aggregate, userCount)
		for i := 0; i < userCount; i++ {
			fmt.Fscanf(in, "%s %s\n", &inputName, &inputPhone)
			var a = Aggregate{
				Name:   inputName,
				Phones: []string{inputPhone}}
			inputAggregates[userCount-i-1] = a
		}

		for _, inputAggr := range inputAggregates {
			col.Add(inputAggr)
		}

		keys := make([]string, 0, len(col))
		for k := range col {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			// fmt.Println(k, col[k])
			var a = col[k]
			fmt.Fprintf(out, "%s: %d %s\n", a.Name, len(a.Phones), strings.Join(a.Phones, " "))
		}
	}
}
