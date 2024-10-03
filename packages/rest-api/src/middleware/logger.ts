import { Writable } from 'stream'

import winston from 'winston'

const transports = []

if (process.env.NODE_ENV === 'test') {
  transports.push(
    new winston.transports.Stream({
      stream: new Writable({
        write: () => {},
      }),
    })
  )
} else {
  transports.push(new winston.transports.Console())
}

export const logger = winston.createLogger({
  level: 'info',
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.json()
  ),
  transports,
})
