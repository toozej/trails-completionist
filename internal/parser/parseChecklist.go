package parser

import (
	"os"
	"regexp"
	"strings"

	"github.com/toozej/trails-completionist/internal/types"
)

// Extract trail information from file contents
func extractTrailInfoFromChecklist(file *os.File) ([]types.Trail, error) {
	var trails []types.Trail

	return trails, nil
}

func parseTrailNameFromChecklist(input string) string {
	return strings.TrimSpace(input)
}

func parseTrailTypeFromChecklist(input string) string {
	// Regular expression to parse trail type and length
	re := regexp.MustCompile(`^(\S+)\s+(\d+(\.\d+)?)\s+miles$`)

	// FindStringSubmatch returns a slice of strings containing the text of the leftmost match
	match := re.FindStringSubmatch(input)

	var trailType string
	if len(match) == 4 {
		trailType = match[1]
	} else {
		trailType = ""
	}

	return trailType
}

func parseTrailLengthFromChecklist(input string) string {
	// Regular expression to parse trail type and length
	re := regexp.MustCompile(`^(\S+)\s+(\d+(\.\d+)?)\s+miles$`)

	// FindStringSubmatch returns a slice of strings containing the text of the leftmost match
	match := re.FindStringSubmatch(input)

	var trailLength string
	if len(match) == 4 {
		trailLength = match[2]
	} else {
		trailLength = "0.0 miles"
	}

	return trailLength
}

func parseTrailParkFromChecklist(input string) string {
	// Regular expression to parse trail park
	re := regexp.MustCompile(`>\s*([^>]+)$`)

	// FindStringSubmatch returns a slice of strings containing the text of the leftmost match
	match := re.FindStringSubmatch(input)

	var trailPark string
	if len(match) == 2 {
		trailPark = match[1]
	} else {
		trailPark = ""
	}

	return trailPark
}

func ParseTrailsFromChecklist(filename string) ([]types.Trail, error) {
	f, err := fetchFile(filename)
	if err != nil {
		return []types.Trail{}, err
	}

	trails, err := extractTrailInfoFromChecklist(f)
	if err != nil {
		return []types.Trail{}, err
	}

	return trails, nil
}
