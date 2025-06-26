package handlers

import (
	"digital-journal/internal/dal"
	"digital-journal/internal/models"
	"encoding/json"
	"log"
	"net/http"
)


func PostProgram(w http.ResponseWriter, r *http.Request) {
	 op := "POST /program"
    log.Println(op)

    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    var prog models.Program
    if err := json.NewDecoder(r.Body).Decode(&prog); err != nil {
        http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
        return
    }
 
    result, err := dal.DB.Exec(
        `INSERT INTO programs 
            (title, description, type, country, organization, deadline, link, user_id) 
         VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
        prog.Title,
        prog.Description,
        prog.Type,
        prog.Country,
        prog.Organization,
        prog.Deadline,
        prog.Link,
        prog.UserID,
    )
    if err != nil {
        log.Printf("%s error inserting program: %v", op, err)
        http.Error(w, "failed to save program", http.StatusInternalServerError)
        return
    }

    lastID, err := result.LastInsertId()
    if err != nil {
        log.Printf("%s error getting last insert id: %v", op, err)
        http.Error(w, "failed to retrieve program ID", http.StatusInternalServerError)
        return
    }
    prog.ID = int(lastID)

    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(prog); err != nil {
        log.Printf("%s error encoding response: %v", op, err)
    }
}