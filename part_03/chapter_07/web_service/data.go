package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "strconv"
    
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)
var DB *sql.DB

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, using system environment variables")
    }
    
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnvAsInt("DB_PORT", 5432)
    dbUser := getEnv("DB_USER", "postgres")
    dbPassword := getEnv("DB_PASSWORD", "")
    dbName := getEnv("DB_NAME", "testdb")
    dbSSLMode := getEnv("DB_SSLMODE", "disable")
    
    connStr := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
    )
    
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
    
    log.Println("Successfully connected to database!")
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if valueStr, exists := os.LookupEnv(key); exists {
        if value, err := strconv.Atoi(valueStr); err == nil {
            return value
        }
    }
    return defaultValue
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	err =DB.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	statement := "insert into posts(content, author) values($1, $2) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = DB.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = DB.Exec("delete from posts where id = $1",post.Id)
	return
}