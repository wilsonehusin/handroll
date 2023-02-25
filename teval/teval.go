package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/imdario/mergo"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

var (
	logger *log.Logger

	debug = flag.Bool("debug", false, "turn on debug logs")

	valuesFiles = flag.StringSliceP("values", "f", []string{}, "path to values file")
	baseOutDir  = flag.StringP("dist", "o", "dist", "output processed template")
)

func loadValues() (map[string]any, error) {
	allValues := map[string]any{}

	for _, valuesPath := range *valuesFiles {
		logger.Printf("processing %s", valuesPath)

		data, err := os.ReadFile(valuesPath)
		if err != nil {
			return nil, fmt.Errorf("reading file '%s': %w", valuesPath, err)
		}

		v := map[string]any{}
		if err := yaml.Unmarshal(data, &v); err != nil {
			return nil, fmt.Errorf("parsing yaml '%s': %w", valuesPath, err)
		}

		if err := mergo.Merge(&allValues, v, mergo.WithOverride); err != nil {
			return nil, fmt.Errorf("merging data '%s': %w", valuesPath, err)
		}
	}

	return allValues, nil
}

func processTemplate(templatePath string, values map[string]any) error {
	outDir := filepath.Join(*baseOutDir, filepath.Dir(templatePath))
	if err := os.MkdirAll(outDir, 0777); err != nil {
		return fmt.Errorf("creating output directory '%s': %w", outDir, err)
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("parsing template '%s': %w", templatePath, err)
	}

	outFile := filepath.Join(outDir, filepath.Base(templatePath))
	// If a file already exists, remove it.
	// This prevents messy output where old file is longer than new file,
	// which results in merged output.
	_ = os.RemoveAll(outFile)
	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return fmt.Errorf("opening destination '%s': %w", outFile, err)
	}
	defer f.Close()

	if err := t.Execute(f, values); err != nil {
		return fmt.Errorf("writing result '%s': %w", templatePath, err)
	}

	return nil
}

func processTemplates(templatePaths []string, values map[string]any) error {
	for _, templatePath := range templatePaths {
		logger.Printf("processing: %s", templatePath)
		if err := processTemplate(templatePath, values); err != nil {
			return fmt.Errorf("processing '%s': %w", templatePath, err)
		}
	}
	return nil
}

func printQuery(query string, values map[string]any) error {
	tmplStr := fmt.Sprintf("{{ %s }}", query)
	logger.Printf("querying: %s", tmplStr)

	t, err := template.New("").Parse(tmplStr)
	if err != nil {
		return fmt.Errorf("templating query '%s': %w", query, err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, values); err != nil {
		return fmt.Errorf("interpolating query: %w", err)
	}

	fmt.Println(buf.String())
	return nil
}

func main() {
	flag.Parse()

	logW := io.Discard
	if *debug {
		logW = os.Stderr
	}
	logger = log.New(logW, "", log.Ldate|log.Ltime|log.Lshortfile)

	allValues, err := loadValues()
	if err != nil {
		panic(err)
	}

	switch cmd := flag.Args()[0]; cmd {
	case "query":
		if err := printQuery(flag.Args()[1], allValues); err != nil {
			panic(err)
		}
	case "run":
		if err := processTemplates(flag.Args()[1:], allValues); err != nil {
			panic(err)
		}
	default:
		fmt.Printf("unknown command: %s", cmd)
	}
}
