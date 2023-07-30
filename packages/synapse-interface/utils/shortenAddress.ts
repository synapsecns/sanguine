export const shortenAddress = (address: string | undefined, chars = 6) => {
  if (address) {
    const start = address.slice(0, chars + 2)
    const end = address.slice(-chars)
    return `${start}...${end}`
  } else {
    return ''
  }
}
