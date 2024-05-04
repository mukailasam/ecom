package models

func (dbExec *Model) CreateUser(username, email, firstName, lastName, password, phoneNumber, salt, token string, verified bool) error {
	pstmt := `INSERT INTO users(username, email, first_name, last_name, 
	                            password, phone, salt, token, verified)
								VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	stmt, err := dbExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, email, firstName, lastName, password, phoneNumber, salt, token, verified)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}
