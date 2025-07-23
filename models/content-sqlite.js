
// better-sqlite3 API: https://github.com/WiseLibs/better-sqlite3/blob/master/docs/api.md
// better-sqlite3 Performance: https://github.com/photostructure/node-sqlite/tree/main/benchmark#-summary
import Database from 'better-sqlite3'

const DATABASE_PATH = 'db/sqlite/content-api.db'

export class SQLiteContentStore {
    constructor() {
        const verbose = process.env.DATABASE_VERBOSE === 'true' ? console.log : null;
        this.db = new Database(DATABASE_PATH, { verbose });
        this.db.pragma('journal_mode = WAL');
    }

    async close() {
        this.db.close();
    }

    async create(content) {
        const now = new Date();
        content.created_at = now.toISOString();
        content.updated_at = now.toISOString();

        const query = `
      INSERT INTO content (id, title, body, author, status, data, created_at, updated_at)
      VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `;

        try {
            this.db.prepare(query).run(
                content.id,
                content.title,
                content.body,
                content.author,
                content.status,
                JSON.stringify(content.data),
                content.created_at,
                content.updated_at,
            );
        } catch (error) {
            throw new Error(`Failed to create content: ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
    }

    async getById(id) {
        const query = `
      SELECT id, title, body, author, status, data, created_at, updated_at
      FROM content WHERE id = ?
    `;

        try {
            const row = this.db.prepare(query).get(id);
            if (!row) return null;
            return {
                id: row.id,
                title: row.title,
                body: row.body,
                author: row.author,
                status: row.status,
                data: typeof row.data === 'string' ? JSON.parse(row.data) : row.data,
                created_at: row.created_at,
                updated_at: row.updated_at,
            };
        } catch (error) {
            throw new Error(`Failed to get content: ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
    }

    async list() {
        const query = `
      SELECT id, title, body, author, status, data, created_at, updated_at
      FROM content ORDER BY created_at DESC LIMIT 100
    `;

        try {
            const rows = this.db.prepare(query).all();
            return rows.map(row => ({
                id: row.id,
                title: row.title,
                body: row.body,
                author: row.author,
                status: row.status,
                data: typeof row.data === 'string' ? JSON.parse(row.data) : row.data,
                created_at: row.created_at,
                updated_at: row.updated_at,
            }));
        } catch (error) {
            throw new Error(`Failed to list content: ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
    }

    async update(content) {
        content.updated_at = new Date().toISOString();

        const query = `
      UPDATE content 
      SET title = ?, body = ?, author = ?, status = ?, data = ?, updated_at = ?
      WHERE id = ?
    `;

        try {
            const result = this.db.prepare(query).run(
                content.title,
                content.body,
                content.author,
                content.status,
                JSON.stringify(content.data),
                content.updated_at,
                content.id,
            );
            if (result.changes === 0) {
                throw new Error(`Update yielded no changes - Content ${content.id} not found?`);
            }
        } catch (error) {
            throw new Error(`Failed to update content: ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
    }

    async delete(id) {
        const query = `DELETE FROM content WHERE id = ?`;

        try {
            const result = this.db.prepare(query).run(id);
            if (result.changes === 0) {
                throw new Error(`Delete yielded no changes for ${id} - Content not found?`);
            }
        } catch (error) {
            throw new Error(`Failed to delete content: ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
    }
}

export function createContentStore() {
    console.log('Creating SQLiteContentStore')
    return new SQLiteContentStore()
}
