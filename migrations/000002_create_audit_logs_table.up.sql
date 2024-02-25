CREATE TABLE IF NOT EXISTS audit_logs (
    log_id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    user_action VARCHAR(100),
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    details TEXT
    );

CREATE INDEX user_id_idx ON audit_logs(user_id);
CREATE INDEX created_time_idx ON audit_logs(created_time);