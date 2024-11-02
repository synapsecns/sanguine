const swaggerJsdoc = require('swagger-jsdoc');
const fs = require('fs');

const options = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: 'Synapse REST API',
      version: '1.8.4',
      description: 'A node.js project exposing a rest api for synapse sdk quotes',
    }
  },
  apis: ['./src/routes/*.ts']
};

const swaggerSpec = swaggerJsdoc(options);
fs.writeFileSync('swagger.json', JSON.stringify(swaggerSpec, null, 2));
console.log('Swagger JSON generated!');
