package generator

import (
	"embed"
	"fmt"
	"html/template"
	"os"

	"github.com/toozej/trails-completionist/internal/types"
)

// Create a map to organize trails by park
func organizeTrails(trails []types.Trail) (map[string][]types.Trail, error) {
	trailsByPark := make(map[string][]types.Trail)
	for _, trail := range trails {
		trailsByPark[trail.Park] = append(trailsByPark[trail.Park], trail)
	}

	return trailsByPark, nil
}

// Create an HTML file
func createHTMLOutputFile(filename string) (*os.File, error) {
	fp, err := os.Create(filename) // #nosec G304
	if err != nil {
		return nil, err
	}

	return fp, nil
}

// Create and execute the template
func executeHTMLTemplate(fp *os.File, tmpl *embed.FS, trailsByPark map[string][]types.Trail) error {
	t := template.Must(template.ParseFS(tmpl, "*/*.html.tmpl"))
	err := t.Execute(fp, trailsByPark)
	if err != nil {
		return err
	}

	defer fp.Close()
	return nil
}

// Create HTML page using template
func GenerateHTMLOutput(filename string, trails []types.Trail) error {
	trailsByPark, err := organizeTrails(trails)
	if err != nil {
		return err
	}

	file, err := createHTMLOutputFile(filename)
	if err != nil {
		return err
	}

	err = executeHTMLTemplate(file, &Templates, trailsByPark)
	if err != nil {
		return err
	} else {
		fmt.Println("HTML file generated successfully.")
	}

	return nil
}
