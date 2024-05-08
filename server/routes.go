package server

import (
	"encoding/json"
	"net/http"
	"sudoku-solver/backend"
)

// StartServer starts the HTTP server and defines routes.
func StartServer(port string) error {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/solve", solveHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return http.ListenAndServe(":"+port, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./static/index.html")
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var puzzle [][]int
	if err := json.NewDecoder(r.Body).Decode(&puzzle); err != nil {
		http.Error(w, "Invalid puzzle data", http.StatusBadRequest)
		return
	}

	if !isValidPuzzle(puzzle) {
		http.Error(w, "Invalid puzzle format", http.StatusBadRequest)
		return
	}

	solution := backend.SolveSudoku(puzzle)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solution)
}

func isValidPuzzle(puzzle [][]int) bool {
	// Check if the puzzle has exactly 9 rows and 9 columns
	if len(puzzle) != 9 {
		return false
	}
	for _, row := range puzzle {
		if len(row) != 9 {
			return false
		}
	}
	return true
}
