export class MemoryContentStore {
    constructor() {
        this.db = {};
    }

    async close() {        
    }

    async create(content) {
        if (this.db[content.id]) {
            throw new Error(`Cannot create content ${content.id} - already exists`);
        }
        const now = new Date();
        content.created_at = now.toISOString();
        content.updated_at = now.toISOString();
        this.db[content.id] = {
            id: content.id,
            title: content.title,
            body: content.body,
            author: content.author,
            status: content.status,
            data: content.data,
            created_at: content.created_at,
            updated_at: content.updated_at,
        };
    }

    async getById(id) {
        return this.db[id];
    }

    async list() {
        // TODO: sorting
        return Object.values(this.db).slice(0, 100);
    }

    async update(content) {
        if (!this.db[content.id]) {
            throw new Error(`Cannot update content ${content.id} - not found`);
        }
        content.updated_at = new Date().toISOString();
        this.db[content.id] = {
            ...this.db[content.id],
            ...{
                title: content.title,
                body: content.body,
                author: content.author,
                status: content.status,
                data: content.data,
                updated_at: content.updated_at,
            },
        }
    }

    async delete(id) {
        if (!this.db[id]) {
            throw new Error(`Cannot delete content ${id} - not found`);
        }
        delete this.db[id];
    }
}

export function createContentStore() {
    console.log('Creating MemoryContentStore')
    return new MemoryContentStore()
}
