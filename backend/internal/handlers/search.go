package handlers

import "strings"

func FuzzySearch(query string, corpus []string) []string {
	var results []string
	for _, c := range corpus {
		if contains(c, query) {
			results = append(results, c)
		}
	}
	return results
}

func contains(data, query string) bool {
	return strings.Contains(data, query)
}
