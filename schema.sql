CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE company
(
    id             UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name           VARCHAR(15) UNIQUE NOT NULL,
    description    VARCHAR(3000),
    employee_count INT                NOT NULL,
    registered     BOOLEAN            NOT NULL,
    type           VARCHAR(20)        NOT NULL CHECK ( type IN ('Corporations', 'NonProfit', 'Cooperative',
                                                                'Sole Proprietorship') )
);

CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(15) UNIQUE NOT NULL,
    password   VARCHAR(100)       NOT NULL,
    email      VARCHAR(100)       NOT NULL,
    company_id UUID               REFERENCES company (id) ON DELETE SET NULL
);