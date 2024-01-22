import { hotjar } from 'react-hotjar'
import { useEffect } from 'react'

const HOTJAR_ID = 3835898

function HotJar() {
  useEffect(() => {
    // Hotjar tracking requires window to be defined
    if (typeof window !== 'undefined') {
      console.log('hotjar initialized')
      hotjar.initialize(HOTJAR_ID, 6)
    }
  })

  return <></>
}

export default HotJar
