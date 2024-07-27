# RATA.id Product (Golang v1.19)

This project is **Code Challenge** purpose from **RATA.id** reqrutement test by **Ahmad Saekoni**


## Architecture: Monolith

Kenapa **monolith** ? **monolith** lebih mudah untuk diimplementasikan dan di-maintain untuk tahap awal pengembangan. namun microservice dapat di pertimbangkan ketika saat skala sistem meningkat.

## Structure Directory
- cmd `entrypoint apps`
- internal `internal apps`
-- common `common helper`
-- entity `entity domain`
-- handler `http handler`
-- repo `repository`
-- usecase  `usecase`

## Requirement Development

Sebelum menjalankan project ini pastikan sudah menginstall dan memastikan semua berjalan dengan baik. 
1. **Golang v1.19**


## How to Run

Untuk menjalankan project ini, silahkan ikuti step-step seperti dibawah ini:

1. Clone project

	```bash
	git clone https://github.com/lowkruc/rataid-product-golang.git
	```
2. Masuk ke directory project

	```bash
	cd rataid-product-golang
	```
3. Validasi dependecy dengan perintah

	```bash
	go mod tidy
	```
4. Jalankan service dengan perintah

	```bash
	go run ./cmd/*
	```

## API Contract

- **Get Products**
Endpoint: `/v1/products`
Method: `GET`
Params:
-- `limit` : `int` default `10`
-- `page`: `int` default `1`
Response:

	```json
		{
  "Page": 1,
  "Next": false,
  "Products": [
	    {
	      "id_product": 1,
	      "name": "Product 1",
	      "description": "Product test 1",
	      "price": 12000,
	      "category": "Category 1",
	      "is_out_of_stock": false
	    },
	    {
	      "id_product": 2,
	      "name": "Product 2",
	      "description": "Product test 2",
	      "price": 14000,
	      "category": "Category 2",
	      "is_out_of_stock": false
	    },
	    {
	      "id_product": 3,
	      "name": "Product 3",
	      "description": "Product test 3",
	      "price": 16000,
	      "category": "Category 3",
	      "is_out_of_stock": true
	    }
	  ]
	```

## TODO
Dalam project ini masih banyak yang harus dikembangkan lagi

1. Logging to NewRelic
2. Database Repository with PostgreSQL
3. Error Handling
4. Slack Notification
5. Google Screet Manager for KeyValue
6. Config Implementation with yml, ini, or json for store
7. Define env with staging, production, and development
