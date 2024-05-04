package models

func (dExec *Model) DeleteProduct(id int64) error {
	pstmt := `DELETE FROM products WHERE product_id=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}
