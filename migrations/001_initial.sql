-- Create a table for user roles
CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Populate some initial roles (you can customize as needed)
INSERT INTO roles (name) VALUES 
    ('Admin'),
    ('Project Manager'),
    ('QA Engineer'),
    ('Developer');

-- Create a table for users
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    role_id INT REFERENCES roles(role_id) ON DELETE SET NULL, -- SET NULL to allow users without a role
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Create a table for projects
CREATE TABLE projects (
    project_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Create a table for project defects
CREATE TABLE defects (
    defect_id SERIAL PRIMARY KEY,
    project_id INT REFERENCES projects(project_id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'Open',
    severity VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_project_id ON defects(project_id);

-- Create a table for defect comments
CREATE TABLE defect_comments (
    comment_id SERIAL PRIMARY KEY,
    defect_id INT REFERENCES defects(defect_id) ON DELETE CASCADE,
    user_id INT, -- You might have a user table for authentication
    comment_text TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_defect_id ON defect_comments(defect_id);

-- Create a table to store project-user relationships (permissions)
CREATE TABLE project_user_permissions (
    permission_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    project_id INT REFERENCES projects(project_id) ON DELETE CASCADE,
    can_create_project BOOLEAN DEFAULT false,
    can_view_project BOOLEAN DEFAULT true,
    can_edit_project BOOLEAN DEFAULT false,
    can_comment BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Create an index for faster user_id and project_id lookups in the project_user_permissions table
CREATE INDEX idx_user_project_permissions ON project_user_permissions(user_id, project_id);
