### 1. Initializing project

```bash
go mod init _project_name_
```

### 2. Installation

```bash
go get -u github.com/gorilla/mux
```

```bash
go get -u gorm.io/driver/postgres
```

```bash
go get -u gorm.io/gorm
```

```bash
go get -u github.com/go-playground/validator/v10
```

```bash
go get -u github.com/golang-jwt/jwt/v4
```

### 3. Endpoint

| Endpoint           | Relative Path  | Method | Description                        |
| ------------------ | -------------- | ------ | ---------------------------------- |
| Create Asset       | _/asset_       | POST   | Endpoint to create new asset       |
| Create Wallet      | _/wallet_      | POST   | Endpoint to create new wallet      |
| Create Transaction | _/transaction_ | POST   | Endpoint to create new transaction |
| Get Asset          | _/assets_      | GET    | Endpoint to get all asset          |
| Get Wallet         | _/wallets_     | GET    | Endpoint to get all wallet         |
| Update Asset       | _/asset/id_    | PATCH  | Endpoint to update asset           |
| Update Wallet      | _/wallet/id_   | PATCH  | Endpoint to update wallet          |
| Delete Asset       | _/assets/id_   | DELETE | Endpoint to delete asset           |
| Delete Wallet      | _/wallets/id_  | DELETE | Endpoint to delete wallet          |

### 4. Running

Running Your App with this commad

```
go run main.go
```

or

```
nodemon --exec go run main.go
```
