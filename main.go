package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type platform struct {
	os, arch string
}

func main() {
	log.SetPrefix("buildall: ")
	log.SetFlags(0)

	if len(os.Args[1:]) != 1 {
		log.Fatal("supply a .go source file to build binaries from")
	}
	pkg := os.Args[1]

	os.Mkdir("binaries", os.FileMode(0750))
	platforms, err := getPlatforms()
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan string)
	for _, p := range platforms {
		go compile(pkg, p.os, p.arch, ch)
	}
	for range platforms {
		fmt.Println(<-ch)
	}
}

func getPlatforms() ([]platform, error) {
	cmd := exec.Command("go", "tool", "dist", "list")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var binaries []platform
	for _, line := range strings.Split(string(out), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "/")
		binaries = append(binaries, platform{
			os:   parts[0],
			arch: parts[1],
		})
	}
	return binaries, nil
}

func compile(pkg, opsys, arch string, ch chan string) {
	bin := fmt.Sprintf("%s-%s-%s",
		strings.TrimSuffix(pkg, filepath.Ext(pkg)), opsys, arch)
	output := filepath.Join("binaries/", bin)
	args := []string{"go", "build", "-o", output, pkg}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = append(cmd.Environ(), "GOOS="+opsys, "GOARCH="+arch)

	err := cmd.Run()
	if err != nil {
		ch <- fmt.Sprintf("%-4s %s", "ERR:", strings.Join(cmd.Args, " "))
	} else {
		ch <- fmt.Sprintf("%-4s %s", "OK:", strings.Join(cmd.Args, " "))
	}
}
