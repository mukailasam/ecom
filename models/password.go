package models

func (dExec *Model) UpdatePassword(hashPassword, username string) error {
	pstmt := ` UPDATE users SET password=$1 WHERE username=$2`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(hashPassword, username)
	if err != nil {
		return err
	}

	return nil
}

func (dExec *Model) ResetPassword(token string, prtExpired bool, email string) error {
	pstmt := ` UPDATE users SET password_reset_token=$1, prt_expired=$2 WHERE email=$3`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(token, prtExpired, email)
	if err != nil {
		return err
	}

	return nil

}

type prToken struct {
	Prt        string
	PrtExpired bool
}

func (dExec *Model) Getprt(username string) (*prToken, error) {
	prt := &prToken{}

	pstmt := ` SELECT password_reset_token, prt_expired FROM users WHERE username=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(username).Scan(&prt.Prt, &prt.PrtExpired)
	if err != nil {
		return nil, err
	}

	return prt, nil
}

func (dExec *Model) Setprt(prtExpired bool, username, password string) error {
	pstmt := ` UPDATE users SET password=$1, prt_expired=$2 WHERE username=$3`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(password, prtExpired, username)
	if err != nil {
		return err
	}

	return nil

}
