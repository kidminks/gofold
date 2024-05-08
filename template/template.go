package template

func GetModelTemplate() string {
	return `package {package}

	type {model_name} struct {
		{fields}
	}
	
	func (m *{model_name}) Create(db *sql.DB, {model_name_camel} *{model_name}) error {
		query := {insert_query}
		_, err := {insert_exec}
		if err != nil {
			return err
		}
		return nil
	}
	
	func GetUser(db *sql.DB, id int) (*{model_name}, error) {
		query := "SELECT * FROM {model_name} WHERE id = ?"
		row := db.QueryRow(query, id)
	
		{model_name_camel} := &{model_name}{}
		err := {fetch_row_scan}
		if err != nil {
			return nil, err
		}
		return {model_name_camel}, nil
	}
	
	func UpdateUser(db *sql.DB, id int, {model_name_camel} *{model_name}) error {
		query := {update_query}
		_, err := {update_exec}
		if err != nil {
			return err
		}
		return nil
	}
	
	func DeleteUser(db *sql.DB, id int) error {
		query := {delete_query}
		_, err := {delete_exec}
		if err != nil {
			return err
		}
		return nil
	}`
}
