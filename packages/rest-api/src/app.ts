import express from 'express'

import routes from './routes'

const app = express()
const port = process.env.PORT || 3000

app.use(express.json())
app.use('/', routes)

export const server = app.listen(port, () => {
  console.log(`Server listening at ${port}`)
})
