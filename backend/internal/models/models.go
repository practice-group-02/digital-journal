package models

type Program struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Type        int    `json:"type"`
    Country     string `json:"country"`
    Organization string `json:"organization"`
    Deadline    string `json:"deadline"`
    Link        string `json:"link"`
    UserID      int    `json:"user_id"`
}

type ProgramType struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type User struct {
    ID        int    `json:"id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    Role      string `json:"role"`
    CreatedAt string `json:"created_at"`
}