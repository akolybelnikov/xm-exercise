CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS company
(
    id             UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name           VARCHAR(15) UNIQUE NOT NULL,
    description    VARCHAR(3000),
    employee_count INT                NOT NULL,
    registered     BOOLEAN            NOT NULL,
    type           VARCHAR(20)        NOT NULL CHECK ( type IN ('Corporations', 'NonProfit', 'Cooperative',
                                                                'Sole Proprietorship') )
);

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(15) UNIQUE NOT NULL,
    password   VARCHAR(100)       NOT NULL,
    email      VARCHAR(100)       NOT NULL,
    company_id UUID               REFERENCES company (id) ON DELETE SET NULL
);

INSERT INTO company (name, description, employee_count, registered, type)
VALUES ('Apple',
        'Apple Inc. is an American multinational technology company that specializes in consumer electronics, computer software, and online services.',
        147000, TRUE, 'Corporations');
INSERT INTO users (username, password, email, company_id)
VALUES ('admin', crypt('admin', gen_salt('bf')), 'admin@apple.inc', (SELECT id FROM company WHERE name = 'Apple'));