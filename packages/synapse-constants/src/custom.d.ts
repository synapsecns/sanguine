//Enable us to import svg files cleanly
declare module '*.svg' {
  const content: any
  export default content
}

declare module '*.png' {
  const value: any
  export default value
}
