CREATE TABLE banks(
	bank_id SERIAL PRIMARY KEY,
    bank_name VARCHAR(100) NOT NULL UNIQUE,
    head_office VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO banks(bank_name, head_office)
VALUES ('sbi', 'mumbai'),
	   ('hdfc', 'delhi'); 

SELECT * FROM banks;

CREATE TABLE branches (
	branch_id SERIAL PRIMARY KEY,
	branch_name VARCHAR(100) NOT NULL,
	city VARCHAR(100) NOT NULL,
	address TEXT NOT NULL,

	bank_id INT NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	CONSTRAINTS fk_bank
		FOREIGN KEY (bank_id)
		REFERENCES banks(bank_id)
		ON DELETE RESTRICT
);

INSERT INTO branches(branch_name, city, address, bank_id)
VALUES ('connaught place', 'del');

CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,

    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,

    phone VARCHAR(15) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,

    address TEXT NOT NULL,

    branch_id INT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_customer_branch
        FOREIGN KEY (branch_id)
        REFERENCES branches(branch_id)
        ON DELETE RESTRICT
);

CREATE TABLE accounts(
	account_id SERIAL PRIMARY KEY,
	account_number VARCHAR(20) UNIQUE NOT NULL,
	customer_id INT NOT NULL,
	account_type VARCHAR(20)
		CHECK (account_type IN ('Savings', 'Current')),
	balance NUMERIC(15, 2) DEFAULT 0.00,
	status VARCHAR(10) DEFAULT 'Active'
		CHECK (status IN ('Active', 'Closed')),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	CONSTRAINT fk_customer_id
	FOREIGN KEY (customer_id) 
	REFERENCES customers(customer_id)
	ON DELETE RESTRICT,

	CONSTRAINT unique_customer_account_type
	UNIQUE (customer_id, account_type)
);

CREATE TABLE transactions(
    transaction_id SERIAL PRIMARY KEY,

    account_id INT NOT NULL,

    transaction_type VARCHAR(20)
        CHECK (transaction_type IN ('Deposit', 'Withdraw')),

    amount NUMERIC(15,2)
        CHECK (amount > 0),

    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_transaction_account
        FOREIGN KEY (account_id)
        REFERENCES accounts(account_id)
        ON DELETE RESTRICT
);


