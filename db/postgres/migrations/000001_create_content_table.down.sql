-- Drop indexes first
DROP INDEX IF EXISTS idx_content_data_gin;
DROP INDEX IF EXISTS idx_content_created_at;
DROP INDEX IF EXISTS idx_content_author;
DROP INDEX IF EXISTS idx_content_status;

-- Drop the table
DROP TABLE IF EXISTS content;
