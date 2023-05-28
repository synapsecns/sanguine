const monthNames = [
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

const getUTCDay = (date) => {
  const time = date.getTime()
  return time - (time % 86400000)
}

export function formatDateTime(d) {
  // const month = monthNames[d.getMonth()]
  // const day = d.getDate()
  // const year = d.getFullYear()
  //
  // const time = d.toTimeString()
  // return `${month} ${day}, ${year} ${time}`
  return d.toLocaleString()
}

export function formatDate(date) {
  if (!date) {
    return ''
  }
  const d = new Date(date.replaceAll('-', '/') + ' 00:00:00 UTC')
  const month = monthNames[d.getUTCMonth()]
  const day = d.getUTCDate()
  const year = d.getUTCFullYear()
  if (getUTCDay(d) === getUTCDay(new Date())) {
    return 'Today'
  }
  return `${month} ${day}, ${year}`
}
