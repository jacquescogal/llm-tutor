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


-- Create the user_subject_map_tab table
CREATE TABLE IF NOT EXISTS user_subject_map_tab (
    user_id BIGINT UNSIGNED NOT NULL,                                               -- Foreign key to user_tab
    subject_id BIGINT UNSIGNED NOT NULL,                                            -- Foreign key to subject_tab
    user_subject_role INT UNSIGNED NOT NULL,                                        -- Enum field for user id role in subject defined in memory proto file
    is_favourite BOOLEAN NOT NULL,                                                   -- Boolean field to indicate if the subject is a favorite
    PRIMARY KEY (user_id, subject_id),                                              -- Composite primary key
    FOREIGN KEY (subject_id) REFERENCES subject_tab(subject_id) ON DELETE CASCADE   -- Foreign key constraint
); -- user_subject_map_tab table is a many-to-many relationship table between user_tab and subject_tab tables. It stores relationship between user and subject.

-- Create the module_tab table
CREATE TABLE IF NOT EXISTS module_tab (
    module_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                            -- Primary key with auto increment for module_id         
    module_name VARCHAR(255) NOT NULL,                                               -- Module name field with max length 255 and NOT NULL constraint            
    module_description TEXT NOT NULL,                                                -- Module description field with TEXT data type and NOT NULL constraint
    is_public BOOLEAN NOT NULL,                                                      -- Boolean field to indicate if the module is public or not
    created_time BIGINT UNSIGNED NOT NULL,                                          -- Created time field with BIGINT UNSIGNED data type and NOT NULL constraint            
    updated_time BIGINT UNSIGNED NOT NULL                                      -- Last updated time field with BIGINT UNSIGNED data type and NOT NULL constraint                          
); -- module_tab table stores the module information.

-- full text search index on module_name
CREATE FULLTEXT INDEX idx_module_name ON module_tab(module_name);
-- index on created_time to sort modules by created_time
CREATE INDEX idx_created_time ON module_tab(created_time);

-- Create the subject_module_map_tab table
CREATE TABLE IF NOT EXISTS subject_module_map_tab (
    subject_id BIGINT UNSIGNED NOT NULL,                                            -- Foreign key to subject_tab           
    module_id BIGINT UNSIGNED NOT NULL,                                             -- Foreign key to module_tab 
    PRIMARY KEY (subject_id, module_id),                                             -- Composite primary key       
    FOREIGN KEY (subject_id) REFERENCES subject_tab(subject_id) ON DELETE CASCADE,  -- Foreign key constraint
    FOREIGN KEY (module_id) REFERENCES module_tab(module_id) ON DELETE CASCADE         -- Foreign key constraint           
); -- subject_module_map_tab table is a many-to-many relationship table between subject_tab and module_tab tables. It stores the membership of modules in a subject. Allowing for modules to be shared across multiple subjects.

-- Create the user_module_map_tab table
CREATE TABLE IF NOT EXISTS user_module_map_tab (
    user_id BIGINT UNSIGNED NOT NULL,                                               -- Foreign key to user_tab
    module_id BIGINT UNSIGNED NOT NULL,                                             -- Foreign key to module_tab
    user_module_role INT UNSIGNED NOT NULL,                                         -- Enum field for user id role in module defined in memory proto file
    is_favourite BOOLEAN NOT NULL,                                                   -- Boolean field to indicate if the module is a favorite
    PRIMARY KEY (user_id, module_id),                                               -- Composite primary key
    FOREIGN KEY (module_id) REFERENCES module_tab(module_id) ON DELETE CASCADE       -- Foreign key constraint
); -- user_module_map_tab table is a many-to-many relationship table between user_tab and module_tab tables. It stores relationship between user and module.


-- Create the doc_tab table
CREATE TABLE IF NOT EXISTS doc_tab (
    doc_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                                                      -- Primary key with auto increment for doc_id
    module_id BIGINT UNSIGNED NOT NULL,                                                                      -- Foreign key to module_tab               
    doc_name VARCHAR(255) NOT NULL,                                                                             -- Doc title field with max length 255 and NOT NULL constraint
    doc_description TEXT NOT NULL,                                                                              -- Doc description field with TEXT data type and NOT NULL constraint
    doc_summary TEXT NOT NULL,                                                                              -- Doc summary field with TEXT data type and NOT NULL constraint           
    upload_status INT UNSIGNED NOT NULL,                                                                    -- Enum field for upload status defined in proto file
    s3_object_key VARCHAR(255) NOT NULL,                                                                       -- Object key field with max length 255 and NOT NULL constraint
    created_time BIGINT UNSIGNED NOT NULL,                                                                  -- Created time field with BIGINT UNSIGNED data type and NOT NULL constraint  
    updated_time BIGINT UNSIGNED NOT NULL,                                                             -- Last updated time field with BIGINT UNSIGNED data type and NOT NULL constraint
    FOREIGN KEY (module_id) REFERENCES module_tab(module_id) ON DELETE CASCADE                                 -- Foreign key constraint
); -- doc_tab table stores the document information.

-- full text search index on doc_name
CREATE FULLTEXT INDEX idx_doc_name ON doc_tab(doc_name);
-- index on created_time to sort documents by created_time
CREATE INDEX idx_created_time ON doc_tab(created_time);


-- Create the memory_tab table
CREATE TABLE IF NOT EXISTS memory_tab (
    memory_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                            -- Primary key with auto increment for memory_id   
    doc_id BIGINT UNSIGNED NOT NULL,                                                 -- Foreign key to doc_tab          
    user_id BIGINT UNSIGNED NOT NULL,                                                -- Foreign key to user_tab
    memory_title TEXT NOT NULL,                                             -- Memory title field
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
    user_id BIGINT UNSIGNED NOT NULL,                                                -- Foreign key to user_tab
    question_title TEXT NOT NULL,                                           -- Question title field
    question_blob BLOB NOT NULL,                                                    -- Question blob field to be unmarshalled into a question proto message
    question_type INT UNSIGNED NOT NULL,                                            -- Enum field for question type defined in question proto file and indicates how to unmarshal the question blob
    created_time BIGINT UNSIGNED NOT NULL,                                          -- Created time field with BIGINT UNSIGNED data type and NOT NULL constraint
    updated_time BIGINT UNSIGNED NOT NULL                                     -- Last updated time field with BIGINT UNSIGNED data type and NOT NULL constraint
); -- question_tab table stores question information related to a document.

-- full text search index on question_title
CREATE FULLTEXT INDEX idx_question_title ON question_tab(question_title);
-- index on created_time to sort questions by created_time
CREATE INDEX idx_created_time ON question_tab(created_time);