import { formatDistanceToNow } from 'date-fns'

export function timeAgo({ timestamp }) {
  let timeAgo = '--'
  if (timestamp) {
    const unixTimestamp = timestamp
    const milliseconds = unixTimestamp * 1000
    const date = new Date(milliseconds)
    const timePeriod = formatDistanceToNow(date)

    timeAgo = `${timePeriod}`

    timeAgo = timeAgo.replace('minutes', 'min')
    timeAgo = timeAgo.replace('minute', 'min')
  }

  return timeAgo
}
