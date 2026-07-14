# 🏦 Online Banking System API

A RESTful Online Banking Backend built using **Go**, **Gin**, **GORM**, and **PostgreSQL**. The project provides APIs for customer management, account management, transactions, and loan management while following good backend development practices such as DTOs, database transactions, foreign key constraints, and layered architecture.

---

## 🚀 Features

### Customer Management
- Create Customer
- Get Customer Details

### Account Management
- Create Savings/Current Account
- Deposit Money
- Withdraw Money
- Check Account Balance
- View Last 10 Transactions
- Close Account

### Loan Management
- Apply for Loan
- Repay Loan
- View Loan Details
- View All Loans of an Account

### Database Features
- PostgreSQL
- Foreign Key Constraints
- CHECK Constraints
- UNIQUE Constraints
- Atomic Transactions using GORM Transactions

### Backend Features
- REST APIs
- DTO Pattern
- Layered Architecture
- Cloud Deployment on Render
- PostgreSQL Database hosted on Render

---

# 🛠 Tech Stack

| Technology | Usage |
|------------|------|
| Go | Programming Language |
| Gin | HTTP Framework |
| GORM | ORM |
| PostgreSQL | Database |
| Render | Deployment |
| Postman | API Testing |

---

# 📂 Project Structure

```
online-banking-app/
│
├── database/
│   └── database.go
│
├── dto/
│   ├── request/
│   └── response/
│
├── handlers/
│   ├── customer.go
│   ├── account.go
│   └── loans.go
│
├── models/
│   ├── customer.go
│   ├── account.go
│   ├── transaction.go
│   ├── loan.go
│   └── loan_payment.go
│
├── routes/
│   ├── customer_routes.go
│   ├── accounts_routes.go
│   ├── loan_routes.go
│   └── routes.go
│
├── utils/
│   └── constants.go
│
├── Queries.sql
├── main.go
├── go.mod
└── README.md
```

---

# 🗄 Database Design

## Banks

| Column | Type |
|---------|------|
| bank_id | SERIAL |
| bank_name | VARCHAR |
| head_office | VARCHAR |
| created_at | TIMESTAMP |

---

## Branches

| Column | Type |
|---------|------|
| branch_id | SERIAL |
| branch_name | VARCHAR |
| city | VARCHAR |
| address | TEXT |
| bank_id | INT |
| created_at | TIMESTAMP |

Relationship

```
Bank
   │
   └──────< Branches
```

---

## Customers

| Column | Type |
|---------|------|
| customer_id | SERIAL |
| first_name | VARCHAR |
| last_name | VARCHAR |
| phone | VARCHAR |
| email | VARCHAR |
| address | TEXT |
| branch_id | INT |
| created_at | TIMESTAMP |

Relationship

```
Branch
   │
   └──────< Customers
```

---

## Accounts

| Column | Type |
|---------|------|
| account_id | SERIAL |
| account_number | VARCHAR |
| customer_id | INT |
| account_type | Savings / Current |
| balance | NUMERIC |
| status | Active / Closed |
| created_at | TIMESTAMP |

Relationship

```
Customer
    │
    └──────< Accounts
```

Each customer can have

- One Savings Account
- One Current Account

---

## Transactions

| Column | Type |
|---------|------|
| transaction_id | SERIAL |
| account_id | INT |
| transaction_type | Deposit / Withdraw / Loan Credit |
| amount | NUMERIC |
| created_at | TIMESTAMP |

Relationship

```
Account
    │
    └──────< Transactions
```

---

## Loans

| Column | Type |
|---------|------|
| loan_id | SERIAL |
| account_id | INT |
| loan_type | Personal / Home / Car / Education |
| principal_amount | NUMERIC |
| remaining_amount | NUMERIC |
| interest_rate | NUMERIC |
| status | Active / Paid / Defaulted |
| created_at | TIMESTAMP |

Relationship

```
Account
    │
    └──────< Loans
```

---

## Loan Payments

| Column | Type |
|---------|------|
| payment_id | SERIAL |
| loan_id | INT |
| amount | NUMERIC |
| payment_date | TIMESTAMP |

Relationship

```
Loan
   │
   └──────< Loan Payments
```

---

# 🧩 Database ER Diagram

```
Bank
 │
 └──────< Branch
              │
              └──────< Customer
                             │
                             └──────< Account
                                          │
             ┌────────────────────────────┴───────────────┐
             │                                            │
             ▼                                            ▼
      Transactions                                   Loans
                                                          │
                                                          ▼
                                                  Loan Payments
```

---

# 📌 REST APIs

## Customer APIs

### Create Customer

```
POST /customers
```

### Get Customer

```
GET /customers/:id
```

---

## Account APIs

### Create Account

```
POST /accounts
```

---

### Deposit Money

```
POST /accounts/:id/deposit
```

---

### Withdraw Money

```
POST /accounts/:id/withdraw
```

---

### Get Balance

```
GET /accounts/:id/balance
```

---

### Get Transactions

Returns latest 10 transactions.

```
GET /accounts/:id/transactions
```

---

### Close Account

```
PATCH /accounts/:id/close
```

---

## Loan APIs

### Apply Loan

```
POST /loans
```

---

### Repay Loan

```
POST /loans/:id/payment
```

---

### Get Loan

```
GET /loans/:id
```

---

### Get Loans of an Account

```
GET /accounts/:id/loans
```

---

# 🔄 Transaction Flow

## Deposit

```
Receive Request
       │
Validate Input
       │
Find Account
       │
Verify Active Account
       │
Database Transaction
       │
 ├── Update Balance
 └── Create Transaction
       │
Commit
       │
Return Response
```

---

## Withdraw

```
Receive Request
       │
Validate Input
       │
Find Account
       │
Check Balance
       │
Database Transaction
       │
 ├── Update Balance
 └── Create Transaction
       │
Commit
```

---

## Loan Creation

```
Validate Loan
       │
Verify Account
       │
Create Loan
       │
Credit Account Balance
       │
Create Loan Credit Transaction
```

---

# 📦 HTTP Status Codes Used

| Status | Meaning |
|---------|----------|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 404 | Not Found |
| 409 | Conflict |
| 500 | Internal Server Error |

---

# ▶️ Running Locally

Clone the repository

```bash
git clone https://github.com/<your-username>/Online-Banking-App.git
```

Install dependencies

```bash
go mod tidy
```

Configure your PostgreSQL database connection.

Run

```bash
go run main.go
```

Server starts on

```
localhost:8080
```

---

# 🧪 Testing

The APIs can be tested using

- Postman
- Thunder Client
- cURL

---

# 🌐 Deployment

Backend is deployed on **Render**.

Database is hosted on **Render PostgreSQL**.

---

# 🚧 Future Improvements

- JWT Authentication
- Password Hashing (bcrypt)
- Authorization Middleware
- Swagger Documentation
- Docker Support
- Unit Testing
- Role-Based Access Control
- Logging Middleware
- Rate Limiting
- Refresh Tokens
- CI/CD Pipeline

---

# 👨‍💻 Author

**Ashwani Vahal**

B.Tech CSE, Delhi Technological University (DTU)

Backend Developer | Go | PostgreSQL | Gin | GORM