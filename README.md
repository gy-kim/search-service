# search-service [![Build Status](https://travis-ci.org/gy-kim/search-service.svg?branch=master)](https://travis-ci.org/gy-kim/search-service)

A simple REST API for Searching products.

## Requirements
- Go v1.13
- Docker 

## Run
### Docker-Compose
    $ docker-compose up -d

### Development
    $ cd dev
    $ docker-compose up -d

    $ cd ..
    $ go run main.go

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
