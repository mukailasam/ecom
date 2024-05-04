package models

type userDetails struct {
	UserId    int
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Phone     string
	Salt      string
	Address   string
}

func (dbExec *Model) GetByUsername(username string) (*userDetails, error) {
	ud := userDetails{}
	pstmt := `Select user_id, username, email, password, salt, first_name, last_name, phone, address FROM users WHERE username=$1`
	stmt, err := dbExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&ud.UserId, &ud.Username, &ud.Email, &ud.Password, &ud.Salt, &ud.FirstName, &ud.LastName, &ud.Phone, &ud.Address)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &ud, nil
}

func (dbExec *Model) GetByEmail(email string) (*userDetails, error) {
	ud := userDetails{}
	pstmt := `SELECT username, email, password, first_name, last_name, salt, phone, address FROM users WHERE email=$1`

	stmt, err := dbExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(email).Scan(&ud.Username, &ud.Email, &ud.Password, &ud.FirstName, &ud.LastName, &ud.Salt, &ud.Phone, &ud.Address)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &ud, nil

}

func (dbExec *Model) GetEmail(email string) (*string, error) {
	ud := userDetails{}
	pstmt := `SELECT email FROM users WHERE email=$1`
	stmt, err := dbExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(email).Scan(&ud.Email)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &ud.Email, nil

}

func (dbExec *Model) GetSalt(username string) (*string, error) {
	ud := userDetails{}
	pstmt := `SELECT salt FROM users WHERE username=$1`
	stmt, err := dbExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&ud.Salt)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &ud.Salt, nil

}

/*
func (dbExec *Model) GetID(username string) (*int, error) {
	var uID int
	pstmt := `SELECT id FROM users WHERE username=$1`
	stmt, err := dbExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err := stmt.QueryRow(username).Scan(&uID)
	if err != nil {
		return nil, err
	}

	return uID, nil

}

*/
