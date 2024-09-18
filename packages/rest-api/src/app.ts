// import express from 'express'

// import routes from './routes'

// const app = express()
// const port = process.env.PORT || 3000

// app.use(express.json())
// app.use('/', routes)

// export const server = app.listen(port, () => {
//   console.log(`Server listening at ${port}`)
// })

import express from 'express'
import cors from 'cors'

import routes from './routes'

const app = express()
const port = process.env.PORT || 3000

// Configure CORS to accept requests from localhost:3001
const corsOptions = {
  origin: 'http://localhost:3001',
  methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
  credentials: true, // If you need to support cookies or authentication headers
}

app.use(cors(corsOptions)) // Apply the CORS middleware
app.use(express.json())
app.use('/', routes)

export const server = app.listen(port, () => {
  console.log(`Server listening at ${port}`)
})
