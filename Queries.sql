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

	CONSTRAINT fk_bank
		FOREIGN KEY (bank_id)
		REFERENCES banks(bank_id)
		ON DELETE RESTRICT
);

INSERT INTO branches(branch_name, city, address, bank_id)
VALUES ('connaught place', 'delhi', 'F-40, connaught place, new delhi - 110001', 1);

-- SELECT * FROM BRANCHES;

INSERT INTO branches(branch_name, city, address, bank_id)
VALUES ('karol bagh' ,'delhi', '18A, Ajmal Khan Roadh, Karol Bagh, New Delhi - 110005', 2);

