package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"
    "io"
    "io/ioutil"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func Submit(w http.ResponseWriter, r *http.Request) {
    var submission Submission
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    if err := json.Unmarshal(body, &submission); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422)  // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    fmt.Printf("Body: %s\n", body)  // TODO replace with database write
    fmt.Println("Submission:", submission)  // TODO replace with database write

    Write(submission)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
}

func Write(submission Submission) {
    db, _ := sql.Open("mysql", "test_submissions:test_password@/test")
    defer db.Close()
    
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    stmt, err := db.Prepare("INSERT INTO test.test (type) VALUES (?)")
    if err != nil {
        panic(err)
    }

    _, err = stmt.Exec(submission.Type)
    if err != nil {
        panic(err)
    }
}
