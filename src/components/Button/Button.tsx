import { useEffect } from 'react'

export const Button = () => {
  useEffect(() => {
    console.log('hello')
  }, [])
  return <button onClick={() => alert('clicked')}>Click me</button>
}
