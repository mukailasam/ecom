package models

func (dExec *Model) CreateProduct(name, category string, price float64, description string, other interface{}, contact string, owner_id int, owner_name string) (*int, error) {
	var productID int
	pstmt := `INSERT INTO products(product_name, product_category, product_price, product_description, other, contact, owner_id, owner_name) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING product_id`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(name, category, price, description, other, contact, owner_id, owner_name).Scan(&productID)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &productID, nil

}
