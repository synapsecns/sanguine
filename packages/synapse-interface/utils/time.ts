export const getTimeMinutesFromNow = (minutesFromNow) => {
  const currentTimeSeconds = new Date().getTime() / 1000

  return Math.round(currentTimeSeconds + 60 * minutesFromNow)
}
export const formatDate = (date) => {
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric',
  }).format(date)
}

export const formatTimestampToDate = (timestamp) => {
  return formatDate(new Date(timestamp * 1000))
}
