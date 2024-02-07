export const shortenAddress = (address: string | undefined, chars = 4) => {
  if (address) {
    const start = address.slice(0, chars)
    const end = address.slice(-chars)
    return `${start}...${end}`
  } else {
    return ''
  }
}
