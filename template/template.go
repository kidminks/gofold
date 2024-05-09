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
	
	func (m *{model_name}) Create(db *sql.DB, {model_name_camel} *{model_name}) error {
		query := {insert_query}
		_, err := db.Exec(query, {insert_exec})
		if err != nil {
			return err
		}
		return nil
	}
	
	func GetUser(db *sql.DB, id int) (*{model_name}, error) {
		query := "SELECT * FROM {model_name} WHERE id = ?"
		row := db.QueryRow(query, id)
	
		{model_name_camel} := &{model_name}{}
		err := row.Scan({fetch_row_scan})
		if err != nil {
			return nil, err
		}
		return {model_name_camel}, nil
	}
	
	func UpdateUser(db *sql.DB, id int, {model_name_camel} *{model_name}) error {
		query := {update_query}
		_, err := db.Exec(query, {update_exec}, id)
		if err != nil {
			return err
		}
		return nil
	}
	
	func DeleteUser(db *sql.DB, id int) error {
		query := "DELETE FROM {model_name} WHERE id = ?"
    	_, err := db.Exec(query, id)
		if err != nil {
			return err
		}
		return nil
	}`
}
