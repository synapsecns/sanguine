/*
  Helper Function: timeout
  - setTimeout function to debounce bridge quote call
  */
export function timeout(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}
