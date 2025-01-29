CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(150) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bank_accounts (
                               id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                               user_id UUID NOT NULL,
                               bank_name VARCHAR(100) NOT NULL,
                               account_no VARCHAR(20) NOT NULL UNIQUE,
                               balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE credit_cards (
                              id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              user_id UUID NOT NULL,
                              bank_account_id UUID DEFAULT NULL,
                              card_name VARCHAR(100) NOT NULL,
                              card_number VARCHAR(20) NOT NULL UNIQUE,
                              credit_limit DECIMAL(15, 2) NOT NULL,
                              available DECIMAL(15, 2) NOT NULL,
                              expiration DATE NOT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
                              FOREIGN KEY (bank_account_id) REFERENCES bank_accounts (id) ON DELETE SET NULL
);

CREATE TABLE transactions (
                              id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              user_id UUID NOT NULL,
                              bank_account_id UUID DEFAULT NULL,
                              credit_card_id UUID DEFAULT NULL,
                              amount DECIMAL(15, 2) NOT NULL,
                              category VARCHAR(50) NOT NULL,
                              description TEXT DEFAULT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
                              FOREIGN KEY (bank_account_id) REFERENCES bank_accounts (id) ON DELETE SET NULL,
                              FOREIGN KEY (credit_card_id) REFERENCES credit_cards (id) ON DELETE SET NULL
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    token VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
)