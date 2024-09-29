-- Create the memory_db database
CREATE DATABASE IF NOT EXISTS memory_db;

-- Use the memory_db database
USE memory_db;

-- Create the job_tab table
CREATE TABLE IF NOT EXISTS job_tab (
    job_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    job_status ENUM('queueing', 'processing', 'to approve', 'inserting', 'inserted', 'failed') NOT NULL,
    object_key VARCHAR(255) NOT NULL,
    created_time BIGINT NOT NULL,
    last_updated_time BIGINT NOT NULL
    job_summary TEXT NOT NULL
);

-- Create the subject_tab table
CREATE TABLE IF NOT EXISTS subject_tab (
    subject_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    subject_name VARCHAR(255) NOT NULL
    created_time BIGINT NOT NULL,
);

-- Create the topic_tab table
CREATE TABLE IF NOT EXISTS topic_tab (
    topic_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    subject_id BIGINT NOT NULL,
    topic_name VARCHAR(255) NOT NULL,
    created_time BIGINT NOT NULL,
    FOREIGN KEY (subject_id) REFERENCES subject_tab(subject_id) ON DELETE CASCADE
);

-- Create the memory_tab table
CREATE TABLE IF NOT EXISTS memory_tab (
    memory_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    topic_id BIGINT NOT NULL,
    job_id BIGINT NOT NULL,
    memory_title VARCHAR(255) NOT NULL,
    memory_body TEXT NOT NULL,
    created_time BIGINT NOT NULL,
    last_updated_time BIGINT NOT NULL,
    FOREIGN KEY (topic_id) REFERENCES topic_tab(topic_id) ON DELETE CASCADE,
    FOREIGN KEY (job_id) REFERENCES job_tab(job_id) ON DELETE CASCADE
);

