# search-service [![Build Status](https://travis-ci.org/gy-kim/search-service.svg?branch=master)](https://travis-ci.org/gy-kim/search-service)

A simple REST API for Searching products.

## Requirements
- Go v1.13
- Docker 

## Run
### Docker-Compose
    $ make up
* Wait at least 1 minute for booting up elastic search.

### Development
    $ cd dev
    $ docker-compose up -d

    $ cd ..
    $ go run main.go

## Down
### Docker-Compose
    $ make down

## Usage
### Health
    $ curl --header "x-access-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o"  http://127.0.0.1:9000/health

    {
        "state": "OK"
    }

### GetProducts
    $ curl --header "x-access-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o" http://127.0.0.1:9000/v1/products?q=black_shoes&filter=brand:adidas&page=1&sort=name&sort_asc=false

    {
        "products": [
            {
                "id": "6_product_id_adidas",
                "type": "black_shoes",
                "brand": "adidas",
                "name": "6_product_6_name"
            },
            {
                "id": "5_product_id_adidas",
                "type": "black_shoes",
                "brand": "adidas",
                "name": "5_product_5_name"
            },
            {
                "id": "4_product_id_adidas",
                "type": "black_shoes",
                "brand": "adidas",
                "name": "4_product_4_name"
            },
            {
                "id": "3_product_id_adidas",
                "type": "black_shoes",
                "brand": "adidas",
                "name": "3_product_3_name"
            },
            {
                "id": "2_product_id_adidas",
                "type": "black_shoes",
                "brand": "adidas",
                "name": "2_product_2_name"
            }
        ],
        "page": 1,
        "total_count": 6
    }


## API Documentation
### GetProduct
#### URL : http://127.0.0.1:9000/v1/products?q={query}&filter={property:value}&page={page}&sort={sort}&sort_asc={sort_asc}
   - query: query for product type.  i.e) `black_shoes` | `white_shoes`
   - filter : filter for properties.(propery:value) i.e) `brand:adidas` | `brand:nike`
   - page : pagination. i.e) `1` | `2` | `3`
   - sort : sorting property. i.e) `id` | `product_type` | `brand` | `name`
   - sort_asc : Soring ascending. i.e) `true` | `false`
#### Authentication : 
 -  Key : x-access-token
 -  JWT Secret Key : `secret` (You can generate token through `https://jwt.io/`)


## Sameple Data
	products := []Product{
		{ID: "1_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "1_product_1_name"},
		{ID: "2_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "2_product_2_name"},
		{ID: "3_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "3_product_3_name"},
		{ID: "4_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "4_product_4_name"},
		{ID: "5_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "5_product_5_name"},
		{ID: "6_product_id_adidas", Type: "black_shoes", Brand: "adidas", Name: "6_product_6_name"},
		{ID: "7_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "7_product_7_name"},
		{ID: "8_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "8_product_8_name"},
		{ID: "9_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "9_product_9_name"},
		{ID: "10_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "10_product_10_name"},
		{ID: "11_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "11_product_11_name"},
		{ID: "12_product_id_adidas", Type: "white_shoes", Brand: "adidas", Name: "12_product_12_name"},

		{ID: "13_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "13_product_name"},
		{ID: "14_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "14_product_name"},
		{ID: "15_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "15_product_name"},
		{ID: "16_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "16_product_name"},
		{ID: "17_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "17_product_name"},
		{ID: "18_product_id_nike", Type: "black_shoes", Brand: "nike", Name: "18_product_name"},
		{ID: "19_product_id_nike", Type: "white_shoes", Brand: "nike", Name: "19_product_name"},
		{ID: "20_product_id_nike", Type: "white_shoes", Brand: "nike", Name: "20_product_name"},
		{ID: "21_product_id_nike", Type: "white_shoes", Brand: "nike", Name: "21_product_name"},
	}
