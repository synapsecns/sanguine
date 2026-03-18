export const oneMonthInMinutes: number = 43200
export const oneDayInMinutes: number = 1440

export const getUnixTimeSecondsNow = (): number => {
  return Math.floor(Date.now() / 1000)
}

export const getUnixTimeMinutesFromNow = (minutesFromNow) => {
  return Math.round(getUnixTimeSecondsNow() + 60 * minutesFromNow)
}

export const getUnixTimeMinutesBeforeNow = (minutesBeforeNow) => {
  return Math.round(getUnixTimeSecondsNow() - 60 * minutesBeforeNow)
}

export const calculateTimeBetween = (
  timeBefore: number,
  timeAfter: number
): number => {
  return Math.abs(timeBefore - timeAfter) * 1000
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
  const date = new Date(unixTimestamp * 1000)
  return date.toLocaleString('en-US', { month: 'short', day: 'numeric' })
}

export const isTimestampToday = (unixTimestamp: number): boolean => {
  const currentDate = new Date()
  const dateFromTimestamp = new Date(unixTimestamp * 1000)

  return (
    dateFromTimestamp.getDate() === currentDate.getDate() &&
    dateFromTimestamp.getMonth() === currentDate.getMonth() &&
    dateFromTimestamp.getFullYear() === currentDate.getFullYear()
  )
}

export const convertMsToSeconds = (ms: number) => {
  return Math.ceil(ms / 1000)
}

type CompactDurationLabels = {
  minute: string
  second: string
}

const DEFAULT_COMPACT_DURATION_LABELS: CompactDurationLabels = {
  minute: 'm',
  second: 's',
}

export const formatCompactDuration = (
  durationInSeconds: number,
  labels: CompactDurationLabels = DEFAULT_COMPACT_DURATION_LABELS
): string => {
  const normalizedDuration = Math.max(0, Math.trunc(durationInSeconds))

  if (normalizedDuration < 60) {
    return `${normalizedDuration}${labels.second}`
  }

  const minutes = Math.floor(normalizedDuration / 60)
  const seconds = normalizedDuration % 60

  return `${minutes}${labels.minute}${seconds}${labels.second}`
}
