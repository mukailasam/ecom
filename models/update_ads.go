package models

func (dExec *Model) UpdateProduct(name, category string, price float64, description string, other interface{}, pID int64) error {
	pstmt := `UPDATE products SET product_name=$1, product_category=$2, product_price=$3, product_description=$4, other=$5 WHERE product_id=$6`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, category, price, description, other, pID)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil

}
