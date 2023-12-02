package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

//go:embed impl.tmpl
var implTmpl string

//go:embed test.tmpl
var testTmpl string

type templateData struct {
	Day int
}

func main() {
	var (
		day     int
		baseDir string
	)

	impl, err := template.New("impl").Parse(implTmpl)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse impl template: %v", err))
	}

	test, err := template.New("test").Parse(testTmpl)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse test template: %v", err))
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	date := time.Now().Local()

	flag.IntVar(&day, "day", date.Day(), "The day to scaffold. By default, the current day is used.")
	flag.StringVar(&baseDir, "dir", wd, "Base directory to scaffold to. By default, the current directory is used.")
	flag.Parse()

	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to get absolute path for base directory: %v", err))
	}

	fmt.Printf("Scaffolding day %d to %s\n", day, absBaseDir)

	err = os.MkdirAll(filepath.Join(absBaseDir, fmt.Sprintf("day%d", day)), 0755)
	if err != nil {
		panic(fmt.Sprintf("Failed to create directory: %v", err))
	}

	implPath := filepath.Join(absBaseDir, fmt.Sprintf("day%d/day%d.go", day, day))
	testPath := filepath.Join(absBaseDir, fmt.Sprintf("day%d/day%d_test.go", day, day))

	scaffoldFile(implPath, day, impl)
	scaffoldFile(testPath, day, test)

}

func scaffoldFile(filePath string, day int, impl *template.Template) {
	implFile, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to create impl file: %v", err))
	}
	defer implFile.Close()

	err = impl.Execute(implFile, templateData{Day: day})
	if err != nil {
		panic(fmt.Sprintf("Failed to execute impl template: %v", err))
	}
}
