-- Create the userdb database
CREATE DATABASE IF NOT EXISTS user_db;

-- Use the userdb database
USE user_db;

-- Create the user_tab table
CREATE TABLE IF NOT EXISTS user_tab (
    user_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,     -- Primary key with auto increment for user_id
    username VARCHAR(50) NOT NULL,                          -- Username field with max length 50 and NOT NULL constraint
    hash_salt_password CHAR(60) NOT NULL                    -- Password field with fixed length 60 (bcrypt hash length) and NOT NULL constraint
); -- user_tab table stores the user information.

