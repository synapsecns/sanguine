import swaggerUi from 'swagger-ui-express'
import fs from 'fs'
import path from 'path'

// eslint-disable-next-line @typescript-eslint/no-var-requires
const packageJson = require('../package.json')

const isDevelopment = process.env.NODE_ENV === 'development'
const serverUrl = isDevelopment
  ? 'http://localhost:3000'
  : 'https://api.synapseprotocol.com'

// Load the merged swagger.json
const swaggerFilePath = path.resolve(__dirname, '../swagger.json')  // Corrected path
const swaggerDocument = JSON.parse(fs.readFileSync(swaggerFilePath, 'utf8'))

// Optional: Dynamically update the `servers` field if needed
swaggerDocument.info.version = packageJson.version
swaggerDocument.servers = [{ url: serverUrl }]

export const specs = swaggerDocument
