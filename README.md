# SimpleBank API

SimpleBank is a lightweight, high-performance HTTP API for banking operations. Written in **Go (Golang)**, it utilizes the **Gin** web framework to handle requests and manages database interactions via **SQLC** generated code.

## 🚀 Features

* **Account Management**: Create, retrieve, and list bank accounts.
* **Money Transfers**: Perform secure money transfers between accounts.
* **Data Validation**: Built-in request validation, including custom currency validation.
* **Type-Safe Database Interactions**: Uses an interface-based Store (`db.Store`) for decoupled and testable database operations.

## 🛠 Tech Stack

* **Language**: Go
* **Web Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
* **Database**: PostgreSQL (via [SQLC](https://sqlc.dev/))
* **Validation**: [go-playground/validator](https://github.com/go-playground/validator)

## 📂 API Reference

The server exposes the following RESTful endpoints:

### Accounts

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/accounts` | Create a new bank account |
| `GET` | `/accounts/:id` | Get details of a specific account by ID |
| `GET` | `/accounts` | List accounts (supports pagination parameters) |

### Transfers

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/transfers` | Create a new money transfer between two accounts |
