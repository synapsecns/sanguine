export const oneMonthInMinutes: number = 43200
export const oneDayInMinutes: number = 1440

export const getTimeMinutesFromNow = (minutesFromNow) => {
  const currentTimeSeconds = new Date().getTime() / 1000

  return Math.round(currentTimeSeconds + 60 * minutesFromNow)
}

export const getTimeMinutesBeforeNow = (minutesBeforeNow) => {
  const currentTimeSeconds = new Date().getTime() / 1000

  return Math.round(currentTimeSeconds - 60 * minutesBeforeNow)
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

export const convertUnixTimestampToMonthAndDate = (
  unixTimestamp: number
): string => {
  const months = [
    'Jan',
    'Feb',
    'Mar',
    'Apr',
    'May',
    'Jun',
    'Jul',
    'Aug',
    'Sep',
    'Oct',
    'Nov',
    'Dec',
  ]

  const date = new Date(unixTimestamp * 1000) // Convert seconds to milliseconds
  const monthIndex = date.getMonth()
  const day = date.getDate()

  const formattedDate = `${months[monthIndex]} ${day}`
  return formattedDate
}
