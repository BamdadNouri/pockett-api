## Pockett API

Pocket is a simple money management tool. This is pockett's API, written in Go. Web app repository could be found [here](https://github.com/MhmdNdri/Pockett). Alsoe database documents could be found [here](https://github.com/sepidehghanadi/dbpockett)

### About the architecture
This application has a stateful modules. HTTP handlers could be found in `internal/handlers` package. Checkout the list of endpoints bellow for more information. Modules could be found in `internal/modules` package.

### Authentication
Pockett uses JWT for authentication. Both Register & Login endpoints return a JWT token, as well as the user informations. Providing token in the request header grants access to authenticated endpoints. Access endpoints by setting `"Authorization": "TOKEN"` header.

### Security
Pockett hashes passwords and does not store plain text.

### Configurations
Pockett API uses flags for it's configs. The complete list of configs could be found in `internal/config/config.go`

### Dependencies
Pockett API uses MySQL for data storage. It can be set up using `docker-compose.yml` in the project. Using `database_user, database_password, database_name, database_port` flags database configs could be set to the application.

### Endpoints
`Base_URL` is _`https://pockett.bamdad.dev/api`_.

Health

`GET /healthz`, the response should be `status: 200, "healthy"`

### User

register: `POST /user/register`
```
request: {
    "email": "user@pockett.ir",
    "username": "pockett_user",
    "password": "1234p"
}
```
```
response: {
    "token": "unique_JWT_token",
    "user": {
        "id": 192,
        "email": "user@pockett.ir",
        "username": "pockett_user",
        "theme": 1,
        "defaultWallet": 437,
        "active": true
        }
}
```

login: `POST /user/login`
```
request: {
    "id": "pockett_user" // could be email or username
    "password": "1234p
}
```
```
response: {
    "token": "unique_JWT_token",
    "user": {
        "id": 192,
        "email": "user@pockett.ir",
        "username": "pockett_user",
        "theme": 1,
        "defaultWallet": 437,
        "active": true
        }
}
```
### Transaction
All transaction endpoins require authentications.

create `POST /transaction`
```
request: {
    "amount": 32000,
    "type": 1, // 0: earned, 1: spent
    "description": "groceries",
    "wallet_id": 437
}
```
```
response: {
    "id": 1544,
    "amount": 32000,
    "type": 1,
    "description": "groceries",
    "wallet_id": 437
}
```

get list of transactions GET `/transaction/:wallet_id`
this endpoint has an optional pagination, by providing page and size of pagination in the URL, as an example: `?page=2&size=15`. The default values are page: 1 & size: 10. That means not providing page & size is equals to `?page=1&size=10`
```
response: {
"transactions": [ // list of transactions
    {
    "id": 1544,
    "amount": 32000,
    "type": 1,
    "description": "groceries",
    "wallet_id": 437
    }
],
"wallet": {
    "balance": 32000, // this can be negative, based on overall earn and spendings od this wallet
    "curr": "IRT"
}
}
```
