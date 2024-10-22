import fs from 'fs'
import path from 'path'

// Define the path to your static swagger.json file
const swaggerFilePath = path.join(__dirname, '../swagger.json')

// Read the static swagger.json file
const specs = JSON.parse(fs.readFileSync(swaggerFilePath, 'utf8'))

// Export the specs for use in your application
export { specs }
