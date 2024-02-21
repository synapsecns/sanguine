import { useEffect, useRef, useState } from 'react'
import styles from './ctaBlink.module.css'

export default function Hero() {
  const [h1, setH1] = useState('default')
  const bridgeRef = useRef(null)
  const buildRef = useRef(null)
  const parentRef = useRef(null)

  useEffect(() => {
    if (h1 !== 'bridge' && bridgeRef.current)
      bridgeRef.current.addEventListener('mouseover', setBridge, {
        once: true,
      })

    if (h1 !== 'build' && buildRef.current)
      buildRef.current.addEventListener('mouseover', setBuild, {
        once: true,
      })

    if (h1 !== 'default' && parentRef.current)
      parentRef.current.addEventListener('mouseleave', Reset, { once: true })
  })

  function setBridge() {
    newH1('bridge', 'Any asset to any chain')
  }
  function setBuild() {
    newH1('build', 'Custom everything')
  }
  function Reset() {
    newH1('default', 'Synapse 2.0: The Modular Interchain Network')
  }

  const h1Fragment = (() => {
    switch (h1) {
      case 'bridge':
        return (
          <a href="#">
            Any asset to any chain<span className={styles.underscore}>_</span>
          </a>
        )
      case 'build':
        return (
          <a href="#">
            Custom everything<span className={styles.underscore}>_</span>
          </a>
        )
      default:
        return (
          <>
            Synapse 2.0: The Modular Interchain Network
            <span className={styles.underscore}>_</span>
          </>
        )
    }
  })()

  function newH1(newState, newText) {
    console.log('begin animation')

    let node = document.querySelector('h1')

    if (node.innerText.slice(0, 15) === newText.slice(0, 15)) return

    let start, previousTimeStamp
    let done = false
    let max = newText.length
    const az = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
    let randAZ = () => az[Math.round(Math.random() * (az.length - 1))]

    function step(timeStamp) {
      if (start === undefined) start = timeStamp

      const elapsed = timeStamp - start

      if (previousTimeStamp !== timeStamp) {
        const count = Math.max(1, Math.min((elapsed * elapsed) / 3000, max))
        node.innerText = newText.slice(0, Math.round(count))

        if (node.innerText.length < newText.length - 1)
          node.innerHTML += `<span style="color: hsla(285deg 100% 50% / 1.5);">${randAZ()}</span>`

        if (node.innerText.length < newText.length)
          node.innerHTML += `<span style="color: hsla(280deg 100% 50% / 1.5);">_</span>`

        if (count === max) done = true
      }

      previousTimeStamp = timeStamp

      if (done) {
        console.log('end animation', h1)
        if (newState !== 'default')
          node.innerHTML += `<span class=${styles.arrow}> -></span>`
        setH1(newState)
      } else {
        window.requestAnimationFrame(step)
      }
    }

    window.requestAnimationFrame(step)
  }

  return (
    <header className="my-2 md:my-8 lg:my-16 text-center max-w-3xl grid place-items-center">
      <div className="hidden md:block text-3xl md:text-6xl font-semibold my-4">
        Modular Interchain Messages
      </div>
      <div ref={parentRef}>
        <h1 className="relative my-4 max-w-xl text-3xl md:text-2xl font-medium overflow-hidden">
          <div>{h1Fragment}</div>
        </h1>
        <div className="m-2">
          <a
            ref={bridgeRef}
            className="px-5 pt-1.5 pb-2 text-lg m-2 border border-zinc-500 hover:border-black hover:dark:border-white rounded inline-block bg-white hover:bg-zinc-100 dark:bg-zinc-950 hover:dark:bg-zinc-900"
            href="#"
          >
            Bridge
          </a>
          <a
            ref={buildRef}
            className="px-5 pt-1.5 pb-2 text-lg m-2 border border-fuchsia-500 hover:bg-fuchsia-100 hover:dark:bg-fuchsia-950 rounded inline-block"
            href="#"
          >
            Build
          </a>
        </div>
      </div>
      <p className="leading-relaxed max-w-xl m-2 text-lg dark:font-light tracking-wider">
        Say goodbye to centralized resource pools for cross-chain communication.
        Synapse lets you customize literally every aspect of your interchain
        communications.
      </p>
      <ul className="w-fit md:w-max grid grid-cols-2 md:flex text-base sm:text-lg text-center items-center place-center bg-gradient-to-b from-white to-slate-100 dark:from-zinc-900 dark:to-zinc-950 border border-zinc-200 dark:border-zinc-800 rounded-md px-6 py-2 gap-x-8 gap-y-4 shadow-sm my-4">
        <li className="-mt-1 p-1">50 blockchains</li>
        <li className="-mt-1 p-1">50,000 validators</li>
        <li className="-mt-1 p-1">10.2B messages</li>
        <li className="-mt-1 p-1">$1.2B transferred</li>
      </ul>
    </header>
  )
}
