package main

type Submission struct {
    Type string `json:"type"`  // expenditure or income
    Amount float64 `json:"amount"`  // float
    Category string `json:"category"`  // coffee, fun, ...
    Notes string `json:"notes"`
}

type Submissions []Submission

