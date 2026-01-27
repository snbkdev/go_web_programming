package main

import (
	"database/sql"
	"log"
	"time"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	config, err := LoadConfig()
	if err != nil {
		log.Printf("Warning: %v", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Настройка пула соединений
	Db.SetMaxOpenConns(config.MaxConns)
	Db.SetMaxIdleConns(config.MaxIdle)
	Db.SetConnMaxLifetime(5 * time.Minute)

	if err := Db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Printf("Database connected to %s@%s:%d/%s", 
		config.User, config.Host, config.Port, config.DBName)
}


func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}