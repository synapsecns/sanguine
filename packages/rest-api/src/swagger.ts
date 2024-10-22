// import swaggerJsdoc from 'swagger-jsdoc'
//
// // eslint-disable-next-line @typescript-eslint/no-var-requires
// const packageJson = require('../package.json')
//
// const isDevelopment = process.env.NODE_ENV === 'development'
// const serverUrl = isDevelopment
//   ? 'http://localhost:3000'
//   : 'https://api.synapseprotocol.com'
//
// const options: swaggerJsdoc.Options = {
//   definition: {
//     openapi: '3.0.0',
//     info: {
//       title: 'Synapse Protocol REST API',
//       version: packageJson.version,
//       description: 'API documentation for the Synapse Protocol REST API',
//     },
//     servers: [
//       {
//         url: serverUrl,
//       },
//     ],
//   },
//   apis: ['./src/routes/*.ts', './src/*.ts'],
// }
//
// export const specs = swaggerJsdoc(options)

import fs from 'fs'
import path from 'path'

// Define the path to your static swagger.json file
const swaggerFilePath = path.join(__dirname, '../swagger.json')

// Read the static swagger.json file
const specs = JSON.parse(fs.readFileSync(swaggerFilePath, 'utf8'))

// Export the specs for use in your application
export { specs }
