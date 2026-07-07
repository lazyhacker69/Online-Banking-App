#Online Banking App

--- Database design ---

banks
---------
bank_id
bank_name
head_office
created_at

branches
---------
branch_id
branch_name
city
address
bank_id

customers
---------
customer_id
first_name
last_name
phone
email
address
branch_id
created_at

accounts
---------
account_id
account_number
customer_id
account_type
balance
status
created_at

transactions
------------
transaction_id
account_id
transaction_type
amount
transaction_date

loans
---------
loan_id
customer_id
loan_amount
interest_rate
remaining_amount
status
created_at

loan_payments
---------
payment_id
loan_id
amount
payment_date

--- Project Structure ---

online-banking/
│
├── go.mod
├── main.go
│
├── database/
│   └── database.go
│
├── models/
│   ├── bank.go
│   ├── branch.go
│   ├── customer.go
│   ├── account.go
│   ├── transaction.go
│   ├── loan.go
│   └── loan_payment.go
│
├── handlers/
│   ├── customer.go
│   ├── account.go
│   ├── transaction.go
│   └── loan.go
│
└── routes/
    └── routes.go