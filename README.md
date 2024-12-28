# trails-completionist
A simple Golang application to parse a list of trails, then display that in a searchable HTML table for ease of tracking completion of trails.

## Features
- Responsive design with automatic light/dark mode
- Fuzzy search across all columns
- Advanced filtering with specific column searches
- Easy-to-use interface

## Search Examples
- `completed: yes` - Show only completed trails
- `park name: Forest Park` - Trails in Forest Park
- `Moderate yes` - Moderate trails that are completed
- `5 miles` - Trails with "5" in their length

## Technology Stack
- Go
- HTML5
- CSS3 (with CSS Variables for theming)
- Vanilla JavaScript

## Usage
- Type in the search bar to filter trails
- Use specific column searches like "completed: yes"
- Combine multiple search criteria

## Development
Operations on the trails-completionist application are driven by `make`. See `make help` for more details.
