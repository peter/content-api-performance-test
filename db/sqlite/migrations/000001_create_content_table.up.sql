CREATE TABLE IF NOT EXISTS content (
    id TEXT PRIMARY KEY, -- ULID as primary key
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    author TEXT NOT NULL,
    status TEXT DEFAULT 'draft' CHECK (status IN ('draft', 'published', 'archived')),
    data TEXT, -- JSON data column
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT DEFAULT (datetime('now'))
) STRICT;

CREATE INDEX idx_content_status ON content(status);
CREATE INDEX idx_content_author ON content(author);
CREATE INDEX idx_content_created_at ON content(created_at);
