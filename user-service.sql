CREATE DATABASE IF NOT EXISTS userdb;

USE DATABASE userdb; /*  \c userdb  */

/*registration*/
CREATE TABLE IF NOT EXISTS users {
    id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
};

CREATE UNIQUE INDEX unique_email_deleted_at_null_idx ON users(email) WHERE deleted IS NULL; --unique email, users with non-soft-deleted accounts cannot register again with the same email.
CREATE INDEX deleted_at_idx ON users(deleted_at);

/*track user actions -> audit_logs*/

CREATE TABLE IF NOT EXISTS audit_logs {
    log_id SERIAL PRIMAY KEY,
    user_id UUID REFERENCES users(id),
    user_action VARCHAR(100),
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    details TEXT
};

CREATE INDEX user_id_idx ON audit_logs(user_id);
CREATE INDEX created_time_idx ON audit_logs(created_time);

/*account varification*/
CREATE TABLE IF NOT EXISTS account_verification {
    verificatin_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    token VARCHAR(100),
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP NOT NULL,
    status BOOLEAN NOT NULL
};

CREATE INDEX user_id_iddx ON account_verification(user_id);
CREATE INDEX token_iddx ON account_verification(token);

