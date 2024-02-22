import { useEffect, useRef, useState } from 'react'

export default function Hero() {
  const [h1, setH1] = useState<[cta: string] | [cta: string, index: number]>([
    'default',
  ])

  const bridgeRef = useRef(null)
  const buildRef = useRef(null)

  const [cta, index] = h1

  const ctas = {
    default: {
      tag: 'Synapse 2.0: The Modular Interchain Network',
    },
    bridge: {
      tag: 'Any asset to any chain',
      url: '#',
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

    if (cta !== 'default') {
      document.addEventListener('mousemove', () => setH1(['default', 0]), {
        once: true,
      })
    }
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

  const ctaButtonBaseStyle =
    'px-5 pt-1.5 pb-2 text-lg m-2 border rounded inline-block'

  return (
    <header className="my-0 xs:my-4 md:my-8 lg:my-12 text-center max-w-3xl grid place-items-center">
      <div className="hidden sm:block text-3xl sm:text-6xl font-semibold my-4 cursor-default">
        Modular Interchain Messages
      </div>
      <div
        className="pt-4 pb-12 grid gap-4"
        onMouseMove={(e) => e.stopPropagation()}
      >
        <h1 className="max-w-xl text-3xl md:text-2xl font-medium cursor-default">
          {url ? (
            <a
              href={url}
              onMouseEnter={(e) => {
                const target = e.target as HTMLAnchorElement
                target.querySelector('animate')?.beginElement()
              }}
              className="p-4 hover:text-black hover:dark:text-white"
            >
              <Tagline />
              {index === tag.length && <ArrowBounce />}
            </a>
          ) : (
            <Tagline />
          )}
        </h1>
        <div>
          <a
            ref={cta !== 'bridge' ? bridgeRef : null}
            className={`${ctaButtonBaseStyle} border-zinc-500 hover:border-black hover:dark:border-white bg-white hover:bg-zinc-100 dark:bg-zinc-950 hover:dark:bg-zinc-900`}
            href={ctas.bridge.url}
          >
            Bridge
          </a>
          <a
            ref={cta !== 'build' ? buildRef : null}
            className={`${ctaButtonBaseStyle} border-fuchsia-500 hover:bg-fuchsia-100 hover:dark:bg-fuchsia-950`}
            href={ctas.build.url}
          >
            Build
          </a>
        </div>
      </div>
      <p className="leading-relaxed max-w-xl text-lg dark:font-light tracking-wider -mt-8">
        Say goodbye to centralized resource pools for cross-chain communication.
        Synapse lets you customize literally every aspect of your interchain
        communications.
      </p>
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
