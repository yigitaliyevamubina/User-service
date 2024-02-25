CREATE TABLE IF NOT EXISTS account_verification (
    verification_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    token VARCHAR(100),
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP NOT NULL,
    status BOOLEAN NOT NULL
);

CREATE INDEX user_id_iddx ON account_verification(user_id);
CREATE INDEX token_iddx ON account_verification(token);