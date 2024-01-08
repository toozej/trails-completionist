package generator

import (
	"embed"
	"fmt"
	"html/template"
	"os"

	"github.com/toozej/trails-completionist/internal/types"
)

// generate easy-to-use Markdown-based checklist of trails given parsed list of trails from raw input text file

// format:
// # PDX Trails Completionist
// ## Park 1
// - Trail A
//     - Trail
//     - 7.3 miles
//     - Completed 10/10/2023
// - Trail B
//     - Connector
//     - 0.2 miles
//     - Completed 05/01/2022
// - Trail C
//     - Connector
//     - 0.4 miles
// ## Park 2

// Create a Markdown file
func createMDOutputFile(filename string) (*os.File, error) {
	fp, err := os.Create(filename) // #nosec G304
	if err != nil {
		return nil, err
	}

	return fp, nil
}

func executeMDTemplate(fp *os.File, tmpl *embed.FS, trailsByPark map[string][]types.Trail) error {
	// Create and execute the Markdown template
	t := template.Must(template.ParseFS(tmpl, "*/*.md.tmpl"))
	err := t.Execute(fp, trailsByPark)
	if err != nil {
		return err
	}

	defer fp.Close()
	return nil
}

func GenerateChecklist(filename string, trails []types.Trail) error {
	trailsByPark, err := organizeTrails(trails)
	if err != nil {
		return err
	}

	f, err := createMDOutputFile(filename)
	if err != nil {
		return err
	}

	err = executeMDTemplate(f, &Templates, trailsByPark)
	if err != nil {
		return err
	} else {
		fmt.Println("Markdown checklist file generated successfully.")
	}

	return nil
}
