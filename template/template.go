package template

func GetDefaultConfigTemplate() *ReplaceMap {
	data := `
{
	"folders": ["cmd/server", "internal/model", "internal/handler", "internal/route", "internal/db", "config"],
	"files": ["cmd/server/main.go", "internal/db/db.go", ".gitignore"],
	"config": "config",
	"model": "internal/model",
	"handler": "internal/handler",
	"route": "internal/route",
	"main": "/cmd/server/main.go",
	"db": "/internal/db/db.go",
	"module": "{module}"
}	
	`
	return &ReplaceMap{
		Input: data,
	}
}

func GetGoModTemplate() *ReplaceMap {
	data := `
	module {module}

	go 1.22.0	
	`
	return &ReplaceMap{
		Input: data,
	}
}

func GetDbTemplate() *ReplaceMap {
	data := `
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
	return &ReplaceMap{
		Input: data,
	}
}

func GetMainTemplate() *ReplaceMap {
	data := `
	package main
	
	import (
		"log"
		"net/http"
		"github.com/gorilla/mux"
	
		"{module}/internal/db"
		"{module}/internal/route"
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
	return &ReplaceMap{
		Input: data,
	}
}

func GetModelTemplate() string {
	return `package {package}

	import (
		"database/sql"
	)

	type {model_name} struct {
		Id  int
		{fields}
	}
	
	func Create{model_name}(db *sql.DB, {model_name_camel} *{model_name}) error {
		query := {insert_query}
		_, err := db.Exec(query, {insert_exec})
		if err != nil {
			return err
		}
		return nil
	}
	
	func Get{model_name}(db *sql.DB, id int) (*{model_name}, error) {
		query := "SELECT * FROM {model_name} WHERE id = ?"
		row := db.QueryRow(query, id)
	
		{model_name_camel} := &{model_name}{}
		err := row.Scan({fetch_row_scan})
		if err != nil {
			return nil, err
		}
		return {model_name_camel}, nil
	}
	
	func Update{model_name}(db *sql.DB, id int, {model_name_camel} *{model_name}) error {
		query := {update_query}
		_, err := db.Exec(query, {update_exec}, id)
		if err != nil {
			return err
		}
		return nil
	}
	
	func Delete{model_name}(db *sql.DB, id int) error {
		query := "DELETE FROM {model_name} WHERE id = ?"
    	_, err := db.Exec(query, id)
		if err != nil {
			return err
		}
		return nil
	}`
}

func GetHandlerTemplate() string {
	return `
	package {package}

	import (
		"database/sql"
		"encoding/json"
		"net/http"
		"strconv"
		
		"github.com/gorilla/mux"

		"{model_package_import}"
	)

	func Create{model_name}Handler(db *sql.DB, r *http.Request) error {
		{model_name_camel} := &{model_name}{}
		json.NewDecoder(r.Body).Decode({model_name_camel})
	
		err := {model_package}.Create{model_name}(db, {model_name_camel})
		if err != nil {
		 return err
		}
	
		return nil
	}
	
	func Get{model_name}Handler(db *sql.DB, r *http.Request) (*{model_package}.{model_name}, error) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
	
		o,err := {model_package}.Get{model_name}(db, id)
		if err != nil {
		 return nil, err
		}
	
		return o, nil
	 }
	
	func Update{model_name}Handler(db *sql.DB, r *http.Request) error {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		
		{model_name_camel} := &{model_package}.{model_name}{}
		json.NewDecoder(r.Body).Decode({model_name_camel})
	
		err = {model_package}.Update{model_name}(db, id, {model_name_camel})
		if err != nil {
		 return err
		}
	
		return nil
	}
	
	func Delete{model_name}Handler(db *sql.DB, r *http.Request) error {
		vars := mux.Vars(r)
		idStr := vars["id"]
		{model_name_camel}Id, err := strconv.Atoi(idStr)
	
		err = {model_package}.Delete{model_name}(db, id)
		if err != nil {
		 return err
		}
	
		return nil
	}`
}

func GetRouteTemplate() string {
	return `
	package {package}

	import (
		"{route_package_import}"
	)

	func Create{model_name}Route(w http.ResponseWriter, r *http.Request) error {
		// create db connection
		// call to handler
		// reture result in w
	}

	func Get{model_name}Route(w http.ResponseWriter, r *http.Request) error {
		// create db connection
		// call to handler
		// reture result in w
	}

	func Update{model_name}Route(w http.ResponseWriter, r *http.Request) error {
		// create db connection
		// call to handler
		// reture result in w
	}

	func Delete{model_name}Route(w http.ResponseWriter, r *http.Request) error {
		// create db connection
		// call to handler
		// reture result in w
	}
	`
}
