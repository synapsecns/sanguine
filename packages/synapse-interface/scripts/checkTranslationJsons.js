const fs = require('fs')
const path = require('path')

// Read the en-US.json file
const enUSPath = path.join(__dirname, '../messages', 'en-US.json')
const enUSContent = JSON.parse(fs.readFileSync(enUSPath, 'utf8'))

// Get all JSON files in the messages directory
const messagesDir = path.join(__dirname, '../messages')
const files = fs
  .readdirSync(messagesDir)
  .filter((file) => file.endsWith('.json') && file !== 'en-US.json')

function compareKeys(obj1, obj2, path = '') {
  const differences = []

  for (const key in obj1) {
    if (!(key in obj2)) {
      differences.push(`Missing key "${path}${key}" in compared file`)
    } else if (typeof obj1[key] === 'object' && obj1[key] !== null) {
      differences.push(...compareKeys(obj1[key], obj2[key], `${path}${key}.`))
    }
  }

  for (const key in obj2) {
    if (!(key in obj1)) {
      differences.push(`Extra key "${path}${key}" in compared file`)
    }
  }

  return differences
}

files.forEach((file) => {
  const filePath = path.join(messagesDir, file)
  const content = JSON.parse(fs.readFileSync(filePath, 'utf8'))

  console.log(`\nChecking ${file}:`)
  const differences = compareKeys(enUSContent, content)

  if (differences.length === 0) {
    console.log('All keys match en-US.json')
  } else {
    differences.forEach((diff) => console.log(diff))
  }
})
