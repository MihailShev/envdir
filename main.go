package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	dir, prog := getArgs()

	vars, err := getEnv(dir)

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Cmd{Path: prog, Env: vars}

	out, err := cmd.Output()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(out))
}

func getArgs() (envDir, prog string) {
	for i, arg := range os.Args {
		switch {
		case i == 1:
			envDir = arg
		case i == 2:
			prog = arg
		}
	}
	return
}

func getEnv(envDir string) ([]string, error) {
	dir, err := ioutil.ReadDir(envDir)
	env := make([]string, 0)
	pathSeparator := string(os.PathSeparator)

	if err != nil {
		return nil, err
	}

	for _, info := range dir {
		name := info.Name()
		bytes, err := ioutil.ReadFile(envDir + string(pathSeparator) + name)

		if err != nil {
			return env, err
		}

		env = append(env, fmt.Sprintf("%s=%s", name, string(bytes)))
	}

	return env, nil
}
