package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Config struct {
	Folders []string `json:"folders"`
	Files   []string `json:"files"`
	Config  string   `json:"config"`
	Model   string   `json:"model"`
	Handler string   `json:"handler"`
	Route   string   `json:"route"`
	Main    string   `json:"main"`
	Db      string   `json:"db"`
	Module  string   `json:"module"`
}

func WriteDefaultConfig(module string, f *os.File) error {
	configJson := `
{
	"folders": ["cmd/server", "internal/model", "internal/handler", "internal/route", "internal/db", "config"],
	"files": ["cmd/server/main.go", "internal/db/db.go", ".gitignore"],
	"config": "config",
	"model": "internal/model",
	"handler": "internal/handler",
	"route": "internal/route",
	"main": "/cmd/server/main.go",
	"db": "/internal/db/db.go",
	"module": "` + module + `"
}
	`
	if _, err := fmt.Fprintln(f, configJson); err != nil {
		slog.Error("error in writing default json to file", "error", err)
		return err
	}

	if err := f.Close(); err != nil {
		slog.Error("error in closing file", "error", err)
		return err
	}
	return nil
}

func FetchConfig(configFile string) (*Config, error) {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		slog.Error("error in opening file", "file", configFile)
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var result Config
	json.Unmarshal([]byte(byteValue), &result)
	return &result, nil
}

func WriteGoMod(module string, f *os.File) error {
	goMod := `
module ` + module + `

go 1.22.0
	`
	if _, err := fmt.Fprintln(f, goMod); err != nil {
		slog.Error("error in writing default json to file", "error", err)
		return err
	}

	if err := f.Close(); err != nil {
		slog.Error("error in closing file", "error", err)
		return err
	}

	return nil
}

func WriteMainFile(module string, f *os.File) error {
	main := `
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	` + `"` + module + `/internal/db"` + `
	` + `"` + module + `/internal/route"` + `
)

// CORS configuration
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set Headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	// {new_model_routes} do not remove
 
	return r
}

func main() {
	// init database connection
	db.InitDb(DBConfig{})

	r := NewRouter()
	r.Use(CORS)
	http.Handle("/", r)

	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
	`

	if _, err := fmt.Fprintln(f, main); err != nil {
		slog.Error("error in writing default json to file", "error", err)
		return err
	}

	if err := f.Close(); err != nil {
		slog.Error("error in closing file", "error", err)
		return err
	}

	return nil
}

func WriteDbFile(module string, f *os.File) error {
	db := `
package db

import (
	"database/sql"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}	

var Db *sql.DB

func InitDb(config DBConfig) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.GetUsername(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetDatabase())
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	Db = db
	return nil
}	
	`

	if _, err := fmt.Fprintln(f, db); err != nil {
		slog.Error("error in writing default json to file", "error", err)
		return err
	}

	if err := f.Close(); err != nil {
		slog.Error("error in closing file", "error", err)
		return err
	}

	return nil
}
