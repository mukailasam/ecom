CREATE TABLE IF NOT EXISTS users(
	user_id serial PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	phone VARCHAR(20) NOT NULL,
	address VARCHAR DEFAULT '',
	salt VARCHAR,
	token VARCHAR,
	password_reset_token VARCHAR,
	verified bool DEFAULT false,
	texprire bool DEFAULT false,
	prt_expired bool DEFAULT false,
	
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS products (
	product_id serial PRIMARY KEY,
	product_name VARCHAR(255) NOT NULL,
	product_category VARCHAR(255),
	product_price DOUBLE PRECISION NOT NULL,
	product_description VARCHAR(255) NOT NULL,
	product_rating int DEFAULT 0,
	other json,
	contact VARCHAR(20) NOT NULL,
	owner_id int NOT NULL,
	owner_name VARCHAR(255) NOT NULL,
	
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	
	FOREIGN KEY (owner_id) REFERENCES users(user_id) ON DELETE CASCADE,
	FOREIGN KEY (owner_name) REFERENCES users(username) ON DELETE CASCADE
); 

CREATE TABLE IF NOT EXISTS product_images(
	image_id serial PRIMARY KEY,
	name VARCHAR NOT NULL,
	url VARCHAR NOT NULL,
	product_id int,
	
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now(),
	
	FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
	
);