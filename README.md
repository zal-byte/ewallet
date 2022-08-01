# latihan-cleanarchitecture-golang
Latihan clean architecture di bahasa pemrograman golang
Database bawaan **MySQL**

```go run main.go```


# User
| Method | Endpoint    | Description       |
| ------ | ------      | ------            |
| GET    | /users      | Fetch all data    |
| GET    | /users/{id} | Get data by id    |
| POST   | /users      | Insert new data   |
| PUT    | /users/{id} | Update data by id |
| DELETE | /users/{id} | Delete data by id |

# TO-DO User
- [x] Get all data
- [x] Get data by id
- [x] Insert data
- [x] Update data
- [x] Delete Data

# Wallet
| Method | Endpoint      | Description            |
| ------ | ------        | ------                 |
| GET    | /wallets/{id} | Get Wallet by id       |
| POST   | /wallets      | Create new user wallet |
| PUT    | /wallets/{id} | Update wallet by id    |
| DELETE | /wallets/{id} | Delete wallet by id    |

# TO-DO Wallet
- [x] Get Data By id
- [x] Insert Data
- [x] Update Data
- [x] Delete Data