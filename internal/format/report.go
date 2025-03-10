package format

import (
	"fmt"
	"os"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/olekukonko/tablewriter"
)

// PageData stores URL and link count
type pageData struct {
	URL   string
	Count int
}

// Sort pages by count (desc) and alphabetically (asc)
func sortPages(pages map[string]int) []pageData {
	var sortedPages []pageData
	for url, count := range pages {
		sortedPages = append(sortedPages, pageData{URL: url, Count: count})
	}
	sort.Slice(sortedPages, func(i, j int) bool {
		if sortedPages[i].Count == sortedPages[j].Count {
			return sortedPages[i].URL < sortedPages[j].URL
		}
		return sortedPages[i].Count > sortedPages[j].Count
	})
	return sortedPages
}

// Print report using tablewriter for better formatting
func PrintReport(pages map[string]int, baseURL string) {
	// Styling the header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFA500")).
		Underline(true)

	// Render header
	fmt.Println(headerStyle.Render(fmt.Sprintf("\nREPORT for %s\n", baseURL)))

	// Create table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"URL", "Internal Links"})
	table.SetAutoWrapText(false) // Avoid word wrap issues
	table.SetBorder(true)
	table.SetRowSeparator("-")
	table.SetColumnSeparator("|")
	table.SetCenterSeparator("+")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Add rows
	sortedPages := sortPages(pages)
	for _, p := range sortedPages {
		table.Append([]string{p.URL, fmt.Sprintf("%d", p.Count)})
	}

	// Render table
	table.Render()

}
