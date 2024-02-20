import { useEffect, useRef, useState } from 'react'

export default function Hero() {
  const [h1, setH1] = useState('default')
  const bridgeRef = useRef(null)
  const buildRef = useRef(null)
  const parentRef = useRef(null)

  useEffect(() => {
    console.log('render', h1)

    if (h1 !== 'bridge' && bridgeRef.current)
      bridgeRef.current.addEventListener('mouseover', mouseEnterBridgeButton, {
        once: true,
      })

    if (h1 !== 'build' && buildRef.current)
      buildRef.current.addEventListener('mouseover', mouseEnterBuildButton, {
        once: true,
      })

    if (h1 !== 'default' && parentRef.current)
      parentRef.current.addEventListener('mouseleave', mouseLeaveButton)
  })

  const h1Fragment = (() => {
    switch (h1) {
      case 'bridge':
        return <a href="#">Any asset to any chain_</a>
      case 'build':
        return <a href="#">Custom everything_</a>
      default:
        return 'Synapse 2.0: The Modular Interchain Network'
    }
  })()

  function mouseEnterBridgeButton() {
    newH1('bridge', 'Any asset to any chain')
  }
  function mouseEnterBuildButton() {
    newH1('build', 'Custom everything')
  }
  function mouseLeaveButton() {
    newH1('default', 'Synapse 2.0: The Modular Interchain Network')
  }

  function newH1(newState, newText) {
    console.log('begin animation')

    let h1 = document.querySelector('h1')

    if (h1.innerText.slice(0, 15) === newText.slice(0, 15)) return

    // h1.outerHTML = h1.outerHTML
    // h1 = document.querySelector('h1')

    let start, previousTimeStamp
    let done = false
    let max = newText.length
    const az = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
    let randAZ = () => az[Math.round(Math.random() * (az.length - 1))]

    function step(timeStamp) {
      if (start === undefined) start = timeStamp

      const elapsed = timeStamp - start

      if (previousTimeStamp !== timeStamp) {
        const count = Math.min((elapsed * elapsed) / 3000, max)
        h1.innerText = newText.slice(0, Math.round(count))
        if (h1.innerText.length < newText.length - 1)
          h1.innerHTML += `<span style="color: hsla(285deg 100% 50% / .5);">${randAZ()}</span>`
        if (h1.innerText.length < newText.length)
          h1.innerHTML += `<span style="color: hsla(280deg 100% 50% / .5);">_</span>`
        if (count === max) done = true
      }

      previousTimeStamp = timeStamp

      console.log('end animation')

      done ? setH1(newState) : window.requestAnimationFrame(step)
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
