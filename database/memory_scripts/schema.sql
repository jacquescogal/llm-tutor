-- Create the memory_db database
CREATE DATABASE IF NOT EXISTS memory_db;

-- Use the memory_db database
USE memory_db;

-- Create the subject_tab table
CREATE TABLE IF NOT EXISTS subject_tab (
    subject_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                          -- Primary key with auto increment for subject_id
    subject_name VARCHAR(255) NOT NULL,                                             -- Subject name field with max length 255 and NOT NULL constraint
    subject_description TEXT NOT NULL,                                              -- Subject description field with TEXT data type and NOT NULL constraint                     
    is_public BOOLEAN NOT NULL,                                                     -- Boolean field to indicate if the subject is public or not
    created_time BIGINT UNSIGNED NOT NULL,                                          -- Created time field with BIGINT UNSIGNED data type and NOT NULL constraint
    updated_time BIGINT UNSIGNED NOT NULL                                      -- Last updated time field with BIGINT UNSIGNED data type and NOT NULL constraint
); -- subject_tab table stores the subject information.

-- full text search index on subject_name
CREATE FULLTEXT INDEX idx_subject_name ON subject_tab(subject_name);
-- index on created_time to sort subjects by created_time
CREATE INDEX idx_created_time ON subject_tab(created_time);


-- Create the member_access_tab table
CREATE TABLE IF NOT EXISTS member_access_tab (
    user_id BIGINT UNSIGNED NOT NULL,                                               -- Foreign key to user_tab
    subject_id BIGINT UNSIGNED NOT NULL,                                            -- Foreign key to subject_tab
    member_role INT UNSIGNED NOT NULL,                                              -- Enum field for member role defined in member proto file
    PRIMARY KEY (user_id, subject_id),                                              -- Composite primary key
    FOREIGN KEY (subject_id) REFERENCES subject_tab(subject_id) ON DELETE CASCADE   -- Foreign key constraint
); -- member_access_tab table is a many-to-many relationship table between user_tab and subject_tab tables. It stores the user's role for a specific subject.

-- Create the topic_tab table
CREATE TABLE IF NOT EXISTS topic_tab (
    topic_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                            -- Primary key with auto increment for topic_id         
    topic_name VARCHAR(255) NOT NULL,                                               -- Topic name field with max length 255 and NOT NULL constraint            
    topic_summary TEXT NOT NULL,                                                    -- Topic summary field with TEXT data type and NOT NULL constraint
    created_time BIGINT UNSIGNED NOT NULL,                                          -- Created time field with BIGINT UNSIGNED data type and NOT NULL constraint            
    updated_time BIGINT UNSIGNED NOT NULL                                      -- Last updated time field with BIGINT UNSIGNED data type and NOT NULL constraint                          
); -- topic_tab table stores the topic information.

-- full text search index on topic_name
CREATE FULLTEXT INDEX idx_topic_name ON topic_tab(topic_name);
-- index on created_time to sort topics by created_time
CREATE INDEX idx_created_time ON topic_tab(created_time);

-- Create the subject_topic_membership_tab table
CREATE TABLE IF NOT EXISTS subject_topic_membership_tab (
    subject_id BIGINT UNSIGNED NOT NULL,                                            -- Foreign key to subject_tab           
    topic_id BIGINT UNSIGNED NOT NULL,                                              -- Foreign key to topic_tab           
    is_master_topic BOOLEAN NOT NULL,                                               -- Boolean field to indicate if the topic is a master topic or not, only a master topic can be updated while a non-master topic is read-only
    PRIMARY KEY (subject_id, topic_id),                                             -- Composite primary key       
    FOREIGN KEY (subject_id) REFERENCES subject_tab(subject_id) ON DELETE CASCADE,  -- Foreign key constraint
    FOREIGN KEY (topic_id) REFERENCES topic_tab(topic_id) ON DELETE CASCADE         -- Foreign key constraint           
); -- subject_topic_membership_tab table is a many-to-many relationship table between subject_tab and topic_tab tables. It stores the membership of topics in a subject. Allowing for topics to be shared across multiple subjects.

-- Create the doc_tab table
CREATE TABLE IF NOT EXISTS doc_tab (
    doc_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                                                      -- Primary key with auto increment for doc_id
    topic_id BIGINT UNSIGNED NOT NULL,                                                                      -- Foreign key to topic_tab               
    doc_title VARCHAR(255) NOT NULL,                                                                             -- Doc title field with max length 255 and NOT NULL constraint
    doc_summary TEXT NOT NULL,                                                                              -- Doc summary field with TEXT data type and NOT NULL constraint           
    upload_status INT UNSIGNED NOT NULL,                                                                    -- Enum field for upload status defined in proto file
    s3_object_key VARCHAR(255) NOT NULL,                                                                       -- Object key field with max length 255 and NOT NULL constraint
    created_time BIGINT UNSIGNED NOT NULL,                                                                  -- Created time field with BIGINT UNSIGNED data type and NOT NULL constraint  
    updated_time BIGINT UNSIGNED NOT NULL,                                                             -- Last updated time field with BIGINT UNSIGNED data type and NOT NULL constraint
    FOREIGN KEY (topic_id) REFERENCES topic_tab(topic_id) ON DELETE CASCADE                                 -- Foreign key constraint
); -- doc_tab table stores the document information.

-- full text search index on doc_title
CREATE FULLTEXT INDEX idx_doc_title ON doc_tab(doc_title);
-- index on created_time to sort documents by created_time
CREATE INDEX idx_created_time ON doc_tab(created_time);


-- Create the memory_tab table
CREATE TABLE IF NOT EXISTS memory_tab (
    memory_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                            -- Primary key with auto increment for memory_id   
    doc_id BIGINT UNSIGNED NOT NULL,                                                 -- Foreign key to doc_tab          
    memory_title VARCHAR(255) NOT NULL,                                             -- Memory title field with max length 255 and NOT NULL constraint
    memory_content TEXT NOT NULL,                                                    -- Memory content field with TEXT data type and NOT NULL constraint
    created_time BIGINT UNSIGNED NOT NULL,                                           -- Created time field with BIGINT UNSIGNED data type and NOT NULL constraint
    updated_time BIGINT UNSIGNED NOT NULL,                                      -- Last updated time field with BIGINT UNSIGNED data type and NOT NULL constraint
    FOREIGN KEY (doc_id) REFERENCES doc_tab(doc_id) ON DELETE CASCADE                -- Foreign key constraint
); -- memory_tab table stores the memory information related to a document.

-- full text search index on memory_title
CREATE FULLTEXT INDEX idx_memory_title ON memory_tab(memory_title);
-- index on created_time to sort memories by created_time
CREATE INDEX idx_created_time ON memory_tab(created_time);

-- Create the question_tab table
CREATE TABLE IF NOT EXISTS question_tab (
    question_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                         -- Primary key with auto increment for question_id
    doc_id BIGINT UNSIGNED NOT NULL,                                                -- Foreign key to doc_tab
    question_title VARCHAR(255) NOT NULL,                                           -- Question title field with max length 255 and NOT NULL constraint
    question_blob BLOB NOT NULL,                                                    -- Question blob field to be unmarshalled into a question proto message
    question_type INT UNSIGNED NOT NULL,                                            -- Enum field for question type defined in question proto file and indicates how to unmarshal the question blob
    created_time BIGINT UNSIGNED NOT NULL,                                          -- Created time field with BIGINT UNSIGNED data type and NOT NULL constraint
    updated_time BIGINT UNSIGNED NOT NULL                                     -- Last updated time field with BIGINT UNSIGNED data type and NOT NULL constraint
); -- question_tab table stores question information related to a document.

-- full text search index on question_title
CREATE FULLTEXT INDEX idx_question_title ON question_tab(question_title);
-- index on created_time to sort questions by created_time
CREATE INDEX idx_created_time ON question_tab(created_time);