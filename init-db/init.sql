CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS company (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  name VARCHAR(15) UNIQUE NOT NULL,
  description VARCHAR(3000),
  employee_count INT NOT NULL,
  registered BOOLEAN NOT NULL,
  type VARCHAR(20) NOT NULL CHECK ( type IN ('Corporations', 'NonProfit', 'Cooperative', 'Sole Proprietorship') )
);