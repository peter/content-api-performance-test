import { Pool } from 'pg';

// Content represents a content record
export class PostgresContentStore {
    constructor(connString, maxConns = 50, minConns = 5) {
        this.pool = new Pool({
            connectionString: connString,
            max: maxConns,
            min: minConns,
            idleTimeoutMillis: 30 * 60 * 1000, // 30 minutes
            connectionTimeoutMillis: 2000,
        });

        // Test the connection
        this.pool.on('error', (err) => {
            console.error('Unexpected error on idle client', err);
        });
    }

    async close() {
        await this.pool.end();
    }

    async create(content) {
        const now = new Date();
        content.created_at = now;
        content.updated_at = now;

        const query = `
      INSERT INTO content (id, title, body, author, status, data, created_at, updated_at)
      VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `;

        try {
            await this.pool.query(query, [
                content.id,
                content.title,
                content.body,
                content.author,
                content.status,
                JSON.stringify(content.data),
                content.created_at,
                content.updated_at,
            ]);
        } catch (error) {
            throw new Error(`Failed to create content: ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
    }

    async getById(id) {
        const query = `
      SELECT id, title, body, author, status, data, created_at, updated_at
      FROM content WHERE id = $1
    `;

        try {
            const result = await this.pool.query(query, [id]);

            if (result.rows.length === 0) {
                return null;
            }

            const row = result.rows[0];
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
            const result = await this.pool.query(query);

            return result.rows.map(row => ({
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
        content.updated_at = new Date();

        const query = `
      UPDATE content 
      SET title = $1, body = $2, author = $3, status = $4, data = $5, updated_at = $6
      WHERE id = $7
    `;

        try {
            const result = await this.pool.query(query, [
                content.title,
                content.body,
                content.author,
                content.status,
                JSON.stringify(content.data),
                content.updated_at,
                content.id,
            ]);

            if (result.rowCount === 0) {
                throw new Error('Content not found');
            }
        } catch (error) {
            throw new Error(`Failed to update content: ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
    }

    async delete(id) {
        const query = `DELETE FROM content WHERE id = $1`;

        try {
            const result = await this.pool.query(query, [id]);

            if (result.rowCount === 0) {
                throw new Error('Content not found');
            }
        } catch (error) {
            throw new Error(`Failed to delete content: ${error instanceof Error ? error.message : 'Unknown error'}`);
        }
    }
}

// Database configuration - you can move this to environment variables
// postgres://postgres:postgres@localhost:5432/content_api?sslmode=disable
const dbConfig = {
    host: process.env.DB_HOST || 'localhost',
    port: parseInt(process.env.DB_PORT) || 5432,
    database: process.env.DB_NAME || 'content_api',
    user: process.env.DB_USER || 'postgres',
    password: process.env.DB_PASSWORD || 'postgres',
    maxConns: parseInt(process.env.DB_MAX_CONNS) || 50,
    minConns: parseInt(process.env.DB_MIN_CONNS) || 5
  };
  
// Factory function to create ContentStore instances
export function createContentStore() {
    console.log('Creating PostgresContentStore')
    const connString = `postgresql://${dbConfig.user}:${dbConfig.password}@${dbConfig.host}:${dbConfig.port}/${dbConfig.database}`;
    return new PostgresContentStore(
        connString,
        dbConfig.maxConns || 50,
        dbConfig.minConns || 5
    );
}
