CREATE TABLE IF NOT EXISTS content (
    id VARCHAR(26) PRIMARY KEY, -- ULID as primary key (26 characters)
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    author TEXT NOT NULL,
    status TEXT DEFAULT 'draft' CHECK (status IN ('draft', 'published', 'archived')),
    data JSONB NOT NULL DEFAULT '{}', -- JSONB data column for better performance
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX idx_content_status ON content(status);
CREATE INDEX idx_content_author ON content(author);
CREATE INDEX idx_content_created_at ON content(created_at);

-- Create a GIN index on the JSONB data column for efficient JSON queries
CREATE INDEX idx_content_data_gin ON content USING GIN (data);
