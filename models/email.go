package models

func (dExec *Model) VerifyEmail(username string) error {
	pstmt := `UPDATE users SET verified=true WHERE username=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (dExec *Model) IsVerify(email string) (*bool, error) {
	verified := false
	pstmt := `SELECT verified FROM users WHERE email=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(email).Scan(&verified)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &verified, nil

}

func (dExec *Model) GetToken(username string) (string, error) {
	var token string
	pstmt := `SELECT token FROM users Where username=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return "", err
	}

	err = stmt.QueryRow(username).Scan(&token)
	if err != nil {
		return "", err
	}

	err = stmt.Close()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (dExec *Model) TokenExpired(username string) (*bool, error) {
	tExpired := false
	pstmt := `SELECT texprire FROM users WHERE username=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&tExpired)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &tExpired, nil
}

func (dExec *Model) ExpiredToken(username string) error {
	pstmt := `UPDATE users SET texprire=true where username=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil

}
