import swaggerJsdoc from 'swagger-jsdoc'

const options: swaggerJsdoc.Options = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: 'Syanpse Protocol REST API',
      version: '1.0.77',
      description: 'API documentation for the Synapse Protocol REST API',
    },
    servers: [
      {
        url: 'http://localhost:3000',
      },
    ],
  },
  apis: ['./src/routes/*.ts', './src/*.ts'],
}

export const specs = swaggerJsdoc(options)
