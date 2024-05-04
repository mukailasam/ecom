package models

import (
	"encoding/json"
)

type Products struct {
	Product_id  int
	OwnerID     int
	OwnerName   string
	Contact     string
	Name        string
	Category    string
	Price       float64
	Description string
	Other       interface{}
}

func (dExec *Model) ReadProduct(id int64) (*Products, error) {
	pd := &Products{}
	pstmt := `SELECT product_id, owner_id, owner_name, contact, product_name, product_category, product_price, product_description, other from products WHERE product_id=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	var other []byte
	err = stmt.QueryRow(id).Scan(&pd.Product_id, &pd.OwnerID, &pd.OwnerName, &pd.Contact, &pd.Name, &pd.Category, &pd.Price, &pd.Description, &other)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(other, &pd.Other)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return pd, nil
}

func (dExec *Model) ListProducts() (interface{}, error) {
	pd := []Products{}
	stmt := `SELECT product_id, owner_id, owner_name, contact, product_name, product_category, product_price, product_description, other from products`

	rows, err := dExec.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		p := Products{}
		var other []byte
		err := rows.Scan(&p.Product_id, &p.OwnerID, &p.OwnerName, &p.Contact, &p.Name, &p.Category, &p.Price, &p.Description, &other)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(other, &p.Other)
		if err != nil {
			return nil, err
		}

		pd = append(pd, p)
		_ = pd

	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &pd, nil

}

func (dExec *Model) ListProductsByCategory(category string) (interface{}, error) {
	pd := []Products{}
	pstmt := `SELECT product_id, owner_id, owner_name, contact, product_name, product_category, product_price, product_description, other from products WHERE product_category=$1`
	stmt, err := dExec.DB.Prepare(pstmt)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(category)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := Products{}
		var other []byte
		err = rows.Scan(&p.Product_id, &p.OwnerID, &p.OwnerName, &p.Contact, &p.Name, &p.Category, &p.Price, &p.Description, &other)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(other, &p.Other)
		if err != nil {
			return nil, err
		}

		pd = append(pd, p)
		_ = pd
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &pd, nil

}
