# Simple web api project written with GO language

## Launch project instructions

1. Clone repository 

```git clone https://github.com/MaxDMT16/go-web-api.git```  

2. Move to a project folder 

```cd go-web-api/```

3. Launch a project 

```go run main.go```    

Ready! Project is launched on port `9912`

## API endpoints

#### /api/google-pay/{id}
* `GET` : Get payment link from Google Pay

#### /api/apple-pay/{id}
* `GET` : Get payment link from Apple Pay

#### /api/pay-pal/{id}
* `GET` : Get payment link from PayPal