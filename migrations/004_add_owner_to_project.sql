-- add owner_id to project table

ALTER TABLE projects
ADD COLUMN owner_id INT REFERENCES users(user_id) ON DELETE SET NULL;
