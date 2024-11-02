module.exports = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: 'Synapse REST API',
      version: '1.8.4',
      description: 'A node.js project exposing a rest api for synapse sdk quotes',
    }
  },
  apis: ['./src/routes/*.ts']  // This will look in all .ts files in src directory
};
