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
      // tag: 'with the Synapse blockchain protocol suite.',
      tag: 'The Web connects devices. We connect blockchains.',
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
    // return
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

  const ctaButtonBaseStyle = 'px-6 p-2 text-lg border rounded inline-block'

  return (
    <>
      <div className="hidden xs:block text-5xl sm:text-6xl font-semibold text-center mb-4">
        Reach every chain.
      </div>
      <div className="pb-10 sm:pb-0" onMouseMove={(e) => e.stopPropagation()}>
        <h1 className="text-3xl sm:text-2xl font-medium text-center mt-6 mb-4">
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
        <div className="flex gap-4 text-base sm:text-lg whitespace-nowrap justify-center mt-6">
          <a
            ref={cta !== 'bridge' ? bridgeRef : null}
            href={ctas.bridge.url}
            className={`${ctaButtonBaseStyle} border-zinc-500 hover:border-black hover:dark:border-white bg-white hover:bg-zinc-100 dark:bg-zinc-950 hover:dark:bg-zinc-900`}
          >
            Bridge
          </a>
          <a
            ref={cta !== 'build' ? buildRef : null}
            href={ctas.build.url}
            className={`${ctaButtonBaseStyle} border-fuchsia-500 hover:bg-fuchsia-100 hover:dark:bg-fuchsia-950`}
          >
            Build
          </a>
        </div>
      </div>
    </>
  )
}

const ArrowBounce = () => (
  <svg
    width="14"
    height="14"
    viewBox="0 -8 16 16"
    overflow="visible"
    strokeWidth="4"
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
