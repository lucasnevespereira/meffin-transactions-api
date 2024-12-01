CREATE SEQUENCE transactions_id_seq;
CREATE SEQUENCE categories_id_seq;

CREATE TABLE transactions
(
    id           BIGINT NOT NULL          DEFAULT nextval('transactions_id_seq'),
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP WITH TIME ZONE,
    user_id      TEXT,
    type         TEXT,
    description  TEXT,
    amount       NUMERIC,
    is_fixed     BOOLEAN,
    day_of_month BIGINT,
    end_date     TEXT,
    category     TEXT,
    PRIMARY KEY (id)
);

CREATE TABLE categories
(
    id         BIGINT NOT NULL          DEFAULT nextval('categories_id_seq'),
    user_id    TEXT,
    type      TEXT,
    name       TEXT,
    color      TEXT,
    PRIMARY KEY (id)
);

-- Index for faster query performance on user_id
-- CREATE INDEX idx_transactions_user_id ON transactions(user_id);


-- CREATE TABLE IF NOT EXISTS transactions (
--     id SERIAL PRIMARY KEY,
--     user_id VARCHAR(255) NOT NULL,
--     type VARCHAR(255) NOT NULL,
--     description TEXT,
--     amount DECIMAL(15, 2) NOT NULL,
--     is_fixed BOOLEAN NOT NULL,
--     day_of_month INT,
--     end_date DATE,
--     category VARCHAR(255),
--     created_at TIMESTAMPTZ DEFAULT NOW(),
--     updated_at TIMESTAMPTZ DEFAULT NOW(),
--     deleted_at TIMESTAMPTZ
-- );