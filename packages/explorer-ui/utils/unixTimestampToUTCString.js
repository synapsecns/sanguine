export function unixTimestampToUTCString(unixTimestamp) {
  if (unixTimestamp) {
    return new Date(unixTimestamp * 1000).toUTCString()
  } else {
    return ''
  }
}
