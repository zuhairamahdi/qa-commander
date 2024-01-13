-- add severity table
CREATE TABLE severity (
    severity_id SERIAL PRIMARY KEY,
    severity_name VARCHAR(50) NOT NULL,
);

-- add severity data
INSERT INTO severity (severity_name)
VALUES
    ('Low'),
    ('Medium'),
    ('High'),
    ('Critical');


-- add status table
CREATE TABLE status (
    status_id SERIAL PRIMARY KEY,
    status_name VARCHAR(50) NOT NULL,
);

-- add status data
INSERT INTO status (status_name)
VALUES
    ('Open'),
    ('Closed'),
    ('Reopened'),
    ('In Progress'),
    ('Rejected'),
    ('Duplicate'),
    ('Deferred'),
    ('Not a Defect'),
    ('Not Reproducible'),
    ('Won''t Fix'),
    ('Cannot Reproduce'),
    ('Fixed');

-- update defects table to include severity_id and status_id 
ALTER TABLE defects
ADD COLUMN severity_id INT REFERENCES severity(severity_id) ON DELETE SET NULL,
ADD COLUMN status_id INT REFERENCES status(status_id) ON DELETE SET NULL,
ADD COLUMN assigner_id INT REFERENCES users(user_id) ON DELETE SET NULL,
ADD COLUMN assignee_id INT REFERENCES users(user_id) ON DELETE SET NULL,
DROP COLUMN status,
DROP COLUMN severity;
