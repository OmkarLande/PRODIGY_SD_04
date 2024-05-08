document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("sudokuForm");
    const solutionDiv = document.getElementById("solution");

    form.addEventListener("submit", function(e) {
        e.preventDefault();
        
        const puzzleCells = form.querySelectorAll(".sudoku-cell");
        let puzzle = [];
        let row = [];
        puzzleCells.forEach((cell, index) => {
            const value = cell.value.trim() === "" ? 0 : parseInt(cell.value.trim(), 10); 
            row.push(value);
            if ((index + 1) % 9 === 0) {
                puzzle.push(row);
                row = [];
            }
        });
        console.log("Puzzle: ", JSON.stringify(puzzle));
        fetch("/solve", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(puzzle)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error("Network response was not ok", response.json());
            }
            return response.json();
        })
        .then(solution => {
            displaySolution(solution);
        })
        .catch(error => {
            console.error("Error:", error.message);
            solutionDiv.innerHTML = "An error occurred. Please try again.";
        });
    });

    function displaySolution(solution) {
        solutionDiv.innerHTML = "";
    
        const table = document.createElement("table");
        table.classList.add("solution-table");
    
        for (let i = 0; i < solution.length; i++) {
            const row = document.createElement("tr");
            for (let j = 0; j < solution[i].length; j++) {
                const cell = document.createElement("td");
                cell.textContent = solution[i][j];
                row.appendChild(cell);
            }
            table.appendChild(row);
        }
    
        solutionDiv.appendChild(table);
    }
    

    const fillZerosButton = document.getElementById("fillZerosButton");
    fillZerosButton.addEventListener("click", function() {
        const puzzleCells = form.querySelectorAll(".sudoku-cell");
        puzzleCells.forEach(cell => {
            if (cell.value.trim() === "") {
                cell.value = "0";
            }
        });
    });
});
