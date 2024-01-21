CREATE TABLE comment_changes (
    change_id SERIAL PRIMARY KEY,
    comment_id INT NOT NULL,
    old_text TEXT,
    new_text TEXT,
    change_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    changed_by INT NOT NULL,
    FOREIGN KEY (comment_id) REFERENCES comments (id),
    FOREIGN KEY (changed_by) REFERENCES users (id)
);

CREATE TABLE severity_changes (
    change_id SERIAL PRIMARY KEY,
    defect_id INT NOT NULL,
    old_severity VARCHAR(255),
    new_severity VARCHAR(255),
    change_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    changed_by INT NOT NULL,
    FOREIGN KEY (defect_id) REFERENCES defects (id),
    FOREIGN KEY (changed_by) REFERENCES users (id)
);

CREATE TABLE status_changes (
    change_id SERIAL PRIMARY KEY,
    defect_id INT NOT NULL,
    old_status VARCHAR(255),
    new_status VARCHAR(255),
    change_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    changed_by INT NOT NULL,
    FOREIGN KEY (defect_id) REFERENCES defects (id),
    FOREIGN KEY (changed_by) REFERENCES users (id)
);