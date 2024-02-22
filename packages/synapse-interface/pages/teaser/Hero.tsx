import { useEffect, useRef, useState } from 'react'
import RightArrow from './icons/RightArrow'
import styles from './css-modules/ctaBlink.module.css'

export default function Hero() {
  const [h1, setH1] = useState<[cta: string] | [cta: string, index: number]>([
    'default',
  ])

  const bridgeRef = useRef(null)
  const buildRef = useRef(null)
  const parentRef = useRef(null)

  const [cta, index] = h1

  const ctas = {
    default: {
      tag: 'Synapse 2.0: The Modular Interchain Network',
    },
    bridge: {
      tag: 'Any asset to any chain',
      url: 'https://www.google.com',
    },
    build: {
      tag: 'Custom everything',
      url: '#',
    },
  }

  const { tag, url } = ctas[cta]

  const sleep = (time) => new Promise((resolve) => setTimeout(resolve, time))

  useEffect(() => {
    if (index < tag.length) {
      sleep((index / tag.length) * 5 + 5).then(() => setH1([cta, +index + 1]))
    } else {
      bridgeRef?.current?.addEventListener(
        'mousemove',
        () => setH1(['bridge', 0]),
        { once: true }
      )

      buildRef?.current?.addEventListener(
        'mousemove',
        () => setH1(['build', 0]),
        { once: true }
      )
    }
    // parentRef?.current?.addEventListener(
    //   'mouseleave',
    //   () => setH1(['default', 0]),
    //   { once: true }
    // )
  })

  const Tagline = () => {
    return (
      <>
        {tag.slice(0, index)}
        {index < tag.length - 4 && (
          <span className="text-fuchsia-500/60">
            {String.fromCharCode(Math.random() * 61 + 65)}
          </span>
        )}
        {index < tag.length - 5 && (
          <span className="text-purple-500/60">_</span>
        )}
      </>
    )
  }

  return (
    <header className="my-2 md:my-8 lg:my-16 text-center max-w-3xl grid place-items-center">
      <div className="hidden md:block text-3xl md:text-6xl font-semibold my-4">
        Modular Interchain Messages
      </div>
      <div ref={cta !== 'default' ? parentRef : null}>
        <h1 className="relative my-4 max-w-xl text-3xl md:text-2xl font-medium">
          {url ? (
            <a href={url} className={`cursor-pointer`}>
              <Tagline />
              {index === tag.length && <ArrowBounce />}
            </a>
          ) : (
            <Tagline />
          )}
        </h1>
        <div className="m-2">
          <a
            ref={cta !== 'bridge' ? bridgeRef : null}
            className="px-5 pt-1.5 pb-2 text-lg m-2 border border-zinc-500 hover:border-black hover:dark:border-white rounded inline-block bg-white hover:bg-zinc-100 dark:bg-zinc-950 hover:dark:bg-zinc-900"
            href={ctas.bridge.url}
          >
            Bridge
          </a>
          <a
            ref={cta !== 'build' ? buildRef : null}
            className="px-5 pt-1.5 pb-2 text-lg m-2 border border-fuchsia-500 hover:bg-fuchsia-100 hover:dark:bg-fuchsia-950 rounded inline-block"
            href={ctas.build.url}
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

const ArrowBounce = () => (
  <svg
    width="14"
    height="14"
    viewBox="0 -8 16 16"
    overflow="visible"
    stroke-width="4"
    fill="none"
    preserveAspectRatio="xMaxYMid"
    className="inline ml-2 mb-1"
    xmlns="http://www.w3.org/2000/svg"
  >
    <animate
      attributeName="width"
      values="14; 22; 14"
      dur=".5s"
      calcMode="spline"
      keySplines="0 0 0 1; .5 0 0 1"
    />
    <animate
      attributeName="stroke"
      values="hsl(275deg 100% 60%); hsl(290deg 100% 70%); hsl(275deg 100% 60%)"
      dur="2s"
      repeatCount="indefinite"
    />
    <path d="m16,0 -16,0 m8,-8 8,8 -8,8" />
  </svg>
)
