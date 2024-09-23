import swaggerJsdoc from 'swagger-jsdoc'

const isDevelopment = process.env.NODE_ENV === 'development'

const devServer = {
  url: 'http://localhost:3001/api',
  description: 'Local Development Server',
}

const prodServer = {
  url: 'https://triumphant-magic-production.up.railway.app/api',
  description: 'Production Server',
}

const options: swaggerJsdoc.Options = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: 'RFQ Indexer API',
      version: '1.0.00',
      description: 'API documentation for the RFQ Indexer API',
    },
    servers: isDevelopment ? [devServer, prodServer] : [prodServer, devServer],
  },
  apis: ['./src/routes/*.ts', './src/*.ts'],
}

export const specs = swaggerJsdoc(options)
