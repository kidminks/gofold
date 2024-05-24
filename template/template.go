package template

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
