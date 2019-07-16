package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	dir, prog := readArgs()
	vars, err := readVars(dir)

	if err != nil {
		log.Fatalln(err)
	}

	out := start(prog, vars)

	fmt.Println(string(out))
}

func start(prog string, vars []string) []byte {
	cmd := exec.Cmd{Path: prog, Env: vars}

	out, err := cmd.Output()

	if err != nil {
		log.Fatalln(err)
	}

	return out
}

func readArgs() (envDir, prog string) {
	for i, arg := range os.Args {
		switch {
		case i == 1:
			envDir = arg
		case i == 2:
			prog = arg
		case i > 2:
			break
		}
	}

	return
}

func readVars(envDir string) ([]string, error) {
	dir, err := ioutil.ReadDir(envDir)

	if err != nil {
		return nil, err
	}

	vars := make([]string, 0)

	for _, info := range dir {
		name := info.Name()

		bytes, err := ioutil.ReadFile(fmt.Sprintf("%s%c%s", envDir, os.PathSeparator, name))

		if err == nil {
			vars = append(vars, fmt.Sprintf("%s=%s", name, string(bytes)))
		} else {
			log.Println(err)
		}
	}

	return vars, nil
}
