import { formatDistanceToNow } from 'date-fns'

export const timeAgo = ({ timestamp }) => {
  let timeAgo = '--'
  if (timestamp) {
    const unixTimestamp = timestamp
    const milliseconds = unixTimestamp * 1000
    const date = new Date(milliseconds)
    const timePeriod = formatDistanceToNow(date, { addSuffix: false })

    timeAgo = `${timePeriod}`

    timeAgo = timeAgo.replace('minutes', 'min')
    timeAgo = timeAgo.replace('minute', 'min')
    timeAgo = timeAgo.replace('hours', 'hrs')
    timeAgo = timeAgo.replace('hour', 'hr')
  }

  return timeAgo
}
