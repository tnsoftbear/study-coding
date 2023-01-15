package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Module struct {
	Dependencies []string
	IsFinal      bool
}

type ModuleMap map[string]Module

func (mm *ModuleMap) buildModule(name string) {
	mm.resolve(name)

	var isResolved = false
	for _, nameResolved := range resolvedList {
		if nameResolved == name {
			isResolved = true
			break
		}
	}

	if !isResolved {
		results = append(results, name)
	}

	if len(results) > 0 {
		var list = strings.Join(results, " ")
		fmt.Fprintf(out, "%d %s\n", len(results), list)

		for _, r := range results {
			resolvedList = append(resolvedList, r)
		}
	} else {
		fmt.Fprintln(out, "0")
	}
	results = nil
}

var resolvedList []string
var results []string

func (mm *ModuleMap) resolve(name string) {
	var dependencies = (*mm)[name].Dependencies
	for _, dependency := range dependencies {
		var isResolved = false
		for _, resolved := range results {
			if dependency == resolved {
				isResolved = true
				break
			}
		}

		for _, resolved := range resolvedList {
			if dependency == resolved {
				isResolved = true
				break
			}
		}

		if isResolved {
			continue
		}

		if (*mm)[dependency].IsFinal {
			results = append(results, dependency)
		} else {
			mm.resolve(dependency)
			results = append(results, dependency)
		}
	}
}

func main() {
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for testI := 0; testI < testCount; testI++ {
		results = nil
		resolvedList = nil
		fmt.Fscanf(in, "\n")

		var total, query int
		fmt.Fscanf(in, "%d\n", &total)

		var modules = make(ModuleMap, total)

		for i := 0; i < total; i++ {
			var line, _ = in.ReadString('\n')
			var parts = strings.Split(line, ":")
			parts[1] = strings.TrimSpace(parts[1])
			var dependencies []string
			if len(parts[1]) > 0 {
				dependencies = strings.Split(parts[1], " ")
			} else {
				dependencies = []string{}
			}
			modules[parts[0]] = Module{Dependencies: dependencies, IsFinal: len(dependencies) == 0}
		}

		fmt.Fscanf(in, "%d\n", &query)
		for i := 0; i < query; i++ {
			var name string
			fmt.Fscanf(in, "%s\n", &name)
			modules.buildModule(name)
		}

		fmt.Fprintln(out)
	}
}
