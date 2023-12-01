//Enable us to import svg files cleanly
declare module '*.svg' {
  const content: any
  return content
}
