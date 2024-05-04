package models

type Profile struct {
	Username  string
	Email     string
	Phone     string
	FirstName string
	LastName  string
	Address   string
}

func (dExec *Model) ReadProfile(userName string) (*Profile, error) {
	profile := &Profile{}
	pstmt := ` SELECT username, email, first_name, last_name, phone, address FROM users Where username=$1 `
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(userName).Scan(&profile.Username, &profile.Email, &profile.FirstName, &profile.LastName, &profile.Phone, &profile.Address)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (dExec *Model) UpdateProfile(firstName, lastName, phone, address string, usr string) (*string, error) {
	var pass string
	pstmt := ` UPDATE users SET first_name=$1, last_name=$2, phone=$3, address=$4 WHERE username=$5 RETURNING 'pass'`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(firstName, lastName, phone, address, usr).Scan(&pass)
	if err != nil {
		return nil, err
	}

	return &pass, nil

}

func (dExec *Model) DeleteProfile(userName string) (*string, error) {
	var pass string
	pstmt := ` DELETE FROM users WHERE username=$1 RETURNING 'deleted' `
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(userName).Scan(&pass)
	if err != nil {
		return nil, err
	}

	return &pass, nil
}
