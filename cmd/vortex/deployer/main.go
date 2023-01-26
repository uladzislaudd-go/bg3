package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func p(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	format = "C:\\Users\\%s\\AppData\\Local\\Larian Studios\\Baldur's Gate 3\\Mods"
)

func main() {
	// dir1, err := os.Getwd()
	// p(err)
	dir1 := "G:\\vortex\\staging\\baldursgate3"

	// u, err := user.Current()
	// p(err)
	// dir2 := fmt.Sprintf(format, u.Name)
	dir2 := fmt.Sprintf(format, "user")

	mods := map[string]interface{}{}
	modFixerFound := false
	entries2, err := os.ReadDir(dir2)
	p(err)
	for _, entry := range entries2 {
		name := entry.Name()
		if entry.IsDir() {
			fmt.Printf("found a dir \"%s\" inside \"%s\"\n", name, dir2)
			continue
		}
		if name == "Patch3ModFixer.pak" {
			modFixerFound = true
			continue
		}

		file := filepath.Join(dir2, name)

		path, err := os.Readlink(file)
		p(err)

		_, err = os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				os.Remove(file)
				continue
			}
			panic(err)
		}

		mods[path] = nil
	}

	if !modFixerFound {
		panic("Patch3ModFixer.pak not found")
	}

	fmt.Printf("%#v\n", mods)

	entries1, err := os.ReadDir(dir1)
	p(err)
	for _, entry := range entries1 {
		name := entry.Name()
		if !entry.IsDir() {
			fmt.Printf("found a file \"%s\" inside \"%s\"\n", name, dir1)
			continue
		}

		dir := filepath.Join(dir1, name)
		mod := ""
		entries3, err := os.ReadDir(dir)
		p(err)
		for _, entry := range entries3 {
			name := entry.Name()
			if strings.HasSuffix(name, ".pak") {
				mod = name
				break
			}
		}
		if mod == "" {
			fmt.Printf("dir \"%s\" don't have any .pak files\n", dir)
			continue
		}
		mod2 := filepath.Join(dir, mod)

		if _, ok := mods[mod2]; ok {
			fmt.Printf("mod \"%s\" already deployed\n", mod)
			continue
		}

		p(os.Symlink(mod2, filepath.Join(dir2, mod)))
	}

}
