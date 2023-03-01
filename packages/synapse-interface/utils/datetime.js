export function getFormattedTimeString() {
  const now = new Date()
  return now.toLocaleTimeString()
}

/**
 * @param {Date} date
 */
export function formatDate(date) {
  return new Intl.DateTimeFormat(
    'en-US',
    {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: 'numeric',
      minute: 'numeric'
    }
  ).format(date)
}

/**
 * @param {number} timestamp
 */
export function formatTimestampToDate(timestamp) {
  return formatDate(new Date(timestamp * 1000))
}