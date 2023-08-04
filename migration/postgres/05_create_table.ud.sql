CREATE TABLE "sale_product"(
    
     id UUID PRIMARY KEY,
	 product_id UUID NOT NULL,
	 product_name VARCHAR,
	 produc_price INT NOT NULL,
	 discount INT,
	 discount_type VARCHAR NOT NULL,
	 price_with_discount INT ,
	 discount_price INT,
	 count INT ,
	 total_price INT,
	 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	 updated_at TIMESTAMP
);