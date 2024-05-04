package models

type Details struct {
	username string
	email    string
}

func (dbExec Model) UserExists(username string) error {
	dt := &Details{}
	stmt := "SELECT username FROM users WHERE username=$1"
	err := dbExec.DB.QueryRow(stmt, username).Scan(&dt.username)
	if err != nil {
		return err
	}

	return nil
}

func (dbExec *Model) EmailExists(email string) error {
	dt := &Details{}
	stmt := "SELECT email FROM users WHERE email=$1"
	err := dbExec.DB.QueryRow(stmt, email).Scan(&dt.email)
	if err != nil {
		return err
	}

	return nil
}
