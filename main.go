package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/graphql-go/graphql"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var cfg *ini.File
var db *sqlx.DB

func init() {
	var err error
	cfg, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment:         true,
		UnescapeValueCommentSymbols: true,
	}, "config.ini")
	if err != nil {
		fmt.Printf("Main Init: Failed to read file %v", err)
	}

	fmt.Println(">> Loaded Config File Successfully")
}

func main() {
	dbConnString := fmt.Sprintf("%v:%v@tcp(%v)/%v",
		cfg.Section("Database").Key("Username"),
		cfg.Section("Database").Key("Password"),
		cfg.Section("Database").Key("Hostname"),
		cfg.Section("Database").Key("Database"),
	)
	var err error
	db, err = sqlx.Connect("mysql", dbConnString)
	defer db.Close()
	if err != nil {
		fmt.Println("SQL Connect Error: Main - " + err.Error())
	}
	fmt.Println(">> Database Connection opened successfully")

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	fmt.Println(">> Server Started.\nWaiting for connections on:\n---http://localhost:8000/graphql---")
	http.Handle("/graphql", h)
	http.ListenAndServe(":8000", nil)
}
