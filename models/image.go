package models

func (dExec *Model) AddImage(name, url string, productID int64) error {
	pstmt := `INSERT INTO product_images(name, url, product_id) VALUES($1, $2, $3)`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, url, productID)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}
