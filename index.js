import Fastify from 'fastify';
import { ulid } from 'ulid';
import * as pgContent from './models/content-postgres.js';
import * as sqliteContent from './models/content-sqlite.js';

const fastify = Fastify({
  logger: { level: 'info' }
});


// Helper function to create a new Content instance
export function createContent(data = {}) {
  return {
      id: data.id || '',
      title: data.title || '',
      body: data.body || '',
      author: data.author || '',
      status: data.status || 'draft',
      data: data.data || {},
      created_at: data.created_at || new Date(),
      updated_at: data.updated_at || new Date(),
  };
}

// Create content store instance
const contentStore = process.env.DATABASE_ENGINE === 'postgres' ? pgContent.createContentStore() : sqliteContent.createContentStore()

// Request/Response schemas for validation
const createContentSchema = {
  schema: {
    body: {
      type: 'object',
      required: ['title', 'body', 'author'],
      properties: {
        title: { type: 'string', minLength: 1 },
        body: { type: 'string', minLength: 1 },
        author: { type: 'string', minLength: 1 },
        status: { 
          type: 'string', 
          enum: ['draft', 'published', 'archived'],
          default: 'draft'
        },
        data: { type: 'object' }
      }
    },
    response: {
      201: {
        type: 'object',
        properties: {
          id: { type: 'string' },
          title: { type: 'string' },
          body: { type: 'string' },
          author: { type: 'string' },
          status: { type: 'string' },
          data: { type: 'object' },
          created_at: { type: 'string', format: 'date-time' },
          updated_at: { type: 'string', format: 'date-time' }
        }
      }
    }
  }
};

const getContentSchema = {
  schema: {
    params: {
      type: 'object',
      required: ['id'],
      properties: {
        id: { type: 'string', minLength: 1 }
      }
    },
    response: {
      200: {
        type: 'object',
        properties: {
          id: { type: 'string' },
          title: { type: 'string' },
          body: { type: 'string' },
          author: { type: 'string' },
          status: { type: 'string' },
          data: {
            type: 'object',
            properties: {
              run_id: { type: 'string' },
              created_at: { type: 'string', format: 'date-time' },
            },
          },
          created_at: { type: 'string', format: 'date-time' },
          updated_at: { type: 'string', format: 'date-time' }
        }
      }
    }
  }
};

const updateContentSchema = {
  schema: {
    params: {
      type: 'object',
      required: ['id'],
      properties: {
        id: { type: 'string', minLength: 1 }
      }
    },
    body: {
      type: 'object',
      properties: {
        title: { type: 'string', minLength: 1 },
        body: { type: 'string', minLength: 1 },
        author: { type: 'string', minLength: 1 },
        status: { 
          type: 'string', 
          enum: ['draft', 'published', 'archived']
        },
        data: { type: 'object' }
      }
    },
    response: {
      200: {
        type: 'object',
        properties: {
          id: { type: 'string' },
          title: { type: 'string' },
          body: { type: 'string' },
          author: { type: 'string' },
          status: { type: 'string' },
          data: { type: 'object' },
          created_at: { type: 'string', format: 'date-time' },
          updated_at: { type: 'string', format: 'date-time' }
        }
      }
    }
  }
};

const deleteContentSchema = {
  schema: {
    params: {
      type: 'object',
      required: ['id'],
      properties: {
        id: { type: 'string', minLength: 1 }
      }
    },
    response: {
      204: {
        type: 'object',
        properties: {}
      }
    }
  }
};

// Hello world route
fastify.get('/', async (request, reply) => {
  return { message: 'Hello World!' };
});

// Health check route
fastify.get('/health', async (request, reply) => {
  return { status: 'ok', timestamp: new Date().toISOString() };
});

// POST /content - Create new content
fastify.post('/content', createContentSchema, async (request, reply) => {
  const { title, body, author, status = 'draft', data = {} } = request.body;
  
  // Generate ULID for content ID (lowercase)
  const id = ulid().toLowerCase();
  
  fastify.log.info('Creating new content', {
    content_id: id,
    title,
    author
  });

  const content = createContent({
    id,
    title,
    body,
    author,
    status,
    data
  });

  try {
    await contentStore.create(content);
    
    fastify.log.info('Successfully created content', {
      content_id: id
    });
    
    reply.code(201);
    return content;
  } catch (error) {
    fastify.log.error('Failed to create content', {
      content_id: id,
      error: error.message
    });
    throw new Error(`Failed to create content: ${error.message}`);
  }
});

// GET /content - List all content
fastify.get('/content', async (request, reply) => {
  try {
    const contents = await contentStore.list();
    return contents;
  } catch (error) {
    fastify.log.error('Failed to list content', { error: error.message });
    throw new Error(`Failed to list content: ${error.message}`);
  }
});

// GET /content/:id - Get content by ID
fastify.get('/content/:id', getContentSchema, async (request, reply) => {
  const { id } = request.params;
  
  fastify.log.info('Fetching content by ID', { content_id: id });

  try {
    const content = await contentStore.getById(id);
    
    if (!content) {
      fastify.log.warn('Content not found', { content_id: id });
      return reply.code(404).send({ error: 'Content not found' });
    }
    
    fastify.log.info('Successfully retrieved content', {
      content_id: id,
      title: content.title
    });
    
    return reply.send(content);
  } catch (error) {
    fastify.log.error('Failed to get content', {
      content_id: id,
      error: error.message
    });
    throw new Error(`Failed to get content: ${error.message}`);
  }
});

// PUT /content/:id - Update content
fastify.put('/content/:id', updateContentSchema, async (request, reply) => {
  const { id } = request.params;
  const updateData = request.body;
  
  try {
    // Get existing content
    const existingContent = await contentStore.getById(id);
    if (!existingContent) {
      return reply.code(404).send({ error: 'Content not found' });
    }
    
    // Update only provided fields
    const updatedContent = {
      ...existingContent,
      ...updateData,
      id // Ensure ID doesn't get overwritten
    };
    
    await contentStore.update(updatedContent);
    
    return updatedContent;
  } catch (error) {
    fastify.log.error('Failed to update content', {
      content_id: id,
      error: error.message
    });
    throw new Error(`Failed to update content: ${error.message}`);
  }
});

// DELETE /content/:id - Delete content
fastify.delete('/content/:id', deleteContentSchema, async (request, reply) => {
  const { id } = request.params;
  
  try {
    await contentStore.delete(id);
    reply.code(204);
    return {};
  } catch (error) {
    fastify.log.error('Failed to delete content', {
      content_id: id,
      error: error.message
    });
    return reply.code(404).send({ error: 'Content not found' });
  }
});

// Graceful shutdown
const gracefulShutdown = async () => {
  fastify.log.info('Shutting down gracefully...');
  await contentStore.close();
  await fastify.close();
  process.exit(0);
};

process.on('SIGTERM', gracefulShutdown);
process.on('SIGINT', gracefulShutdown);

// Start the server
const start = async () => {
  try {
    await fastify.listen({ port: 8888, host: '0.0.0.0' });
    console.log('Server is running on http://localhost:8888');
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};

start();
