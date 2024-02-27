import FauxBridge from './FauxBridge'

export default function ValueProps() {
  return (
    <article className="grid gap-16 md:gap-24 p-4">
      <section>
        <ul className="w-fit md:w-max grid grid-cols-2 md:flex text-base sm:text-lg text-center items-center place-center bg-gradient-to-b from-white to-slate-100 dark:from-zinc-900 dark:to-zinc-950 border border-zinc-200 dark:border-zinc-800 rounded-md px-6 py-2 gap-x-8 gap-y-4 shadow-sm mx-auto cursor-default">
          <li className="-mt-1 p-1">50 blockchains</li>
          <li className="-mt-1 p-1">50,000 validators</li>
          <li className="-mt-1 p-1">10.2B messages</li>
          <li className="-mt-1 p-1">$1.2B transferred</li>
        </ul>
      </section>
      <section className="flex odd:flex-col even:flex-col-reverse md:grid grid-cols-2 gap-x-12 gap-y-4 items-center">
        <div className="grid place-items-center">
          <svg
            width="100%"
            height="100%"
            viewBox="-9 -9 170 118"
            vectorEffect="non-scaling-stroke"
            strokeWidth=".5"
            overflow="visible"
            // preserveAspectRatio="none"
            className="stroke-fuchsia-800/50 dark:stroke-fuchsia-300/50 fill-fuchsia-200/50 dark:fill-fuchsia-300/5 drop-shadow-xl h-64 max-w-sm"
            xmlns="http://www.w3.org/2000/svg"
          >
            <rect
              width="168"
              height="116"
              x="-8"
              y="-8"
              rx="4"
              className="stroke-purple-400/75 fill-[hsl(275deg_100%_96%)] dark:fill-[hsl(300deg_40%_10%)]"
            />
            <path d="m48,11 4,0 m22,0 4,0 m50,11 0,4 m0,22 0,4 m-24,24 -4,0 m-48,0 -4,0 m0,-39 4,0 m11,11 0,4 m0,-30 0,4" />
            <rect width="48" height="48" rx="2" />
            <rect width="22" height="22" x="52" rx="2" />
            <rect width="74" height="22" x="78" rx="2" />
            <rect width="100" height="22" x="52" y="26" rx="2" />
            <rect width="48" height="48" x="0" y="52" rx="2" />
            <rect width="48" height="48" x="52" y="52" rx="2" />
            <rect width="48" height="48" x="104" y="52" rx="2" />
          </svg>
        </div>
        <div>
          <h2 className="text-4xl font-medium mb-4">Reach every user</h2>
          <p className="text-lg leading-relaxed mb-3">
            Synapse enables developers to build truly native cross-chain
            applications with an economically secure method to reach consensus
            on interchain transactions,
          </p>
        </div>
      </section>
      {/* <section>
        <ul className="flex flex-wrap gap-12 bg-white dark:bg-zinc-950 px-8 py-6 border border-zinc-300 dark:border-zinc-800 rounded-lg shadow-lg">
          <li className="basis-56 grow">
            <h2 className="text-4xl font-medium mb-4">Extensible</h2>
            <p className="text-lg mb-2">
              Synapse’s cross-chain messaging contracts can be deployed across
              any blockchain
            </p>
          </li>
          <li className="basis-56 grow">
            <h2 className="text-4xl font-medium mb-4">Secure</h2>
            <p className="text-lg mb-2">
              Synapse employs an Optimistic security model to ensure integrity
              of cross-chain messages
            </p>
          </li>
          <li className="basis-56 grow">
            <h2 className="text-4xl font-medium mb-4">Generalized</h2>
            <p className="text-lg mb-2">
              Any arbitrary data can be sent across chains including contract
              calls, NFTs, snapshots, and more
            </p>
          </li>
        </ul>
      </section> */}
      <section className="flex odd:flex-col even:flex-col-reverse md:grid grid-cols-2 gap-12 items-center">
        <div>
          <h2 className="text-4xl font-medium mb-4">
            Build powerful decentralized apps
          </h2>
          <p className="text-lg leading-relaxed mb-4">
            Synapse Bridge is built on top of the cross-chain infrastructure
            enabling users to seamlessly transfer assets across all blockchains.
            The Bridge has become the most widely-used method to move assets
            cross-chain, offering low cost, fast, and secure bridging.
          </p>
        </div>
        <div className="grid justify-center">
          <FauxBridge />
        </div>
      </section>
      {/* <section>
        <ul className="flex flex-wrap gap-12 bg-white dark:bg-zinc-950 px-8 py-6 border border-zinc-300 dark:border-zinc-800 rounded-lg shadow-lg">
          <li className="basis-56 grow">
            <h2 className="text-4xl font-medium mb-4">Deep Liquidity</h2>
            <p className="text-lg mb-2">
              Swap native assets using our cross-chain AMM liquidity pools
            </p>
          </li>
          <li className="basis-56 grow">
            <h2 className="text-4xl font-medium mb-4">Wide Support</h2>
            <p className="text-lg mb-2">
              Access over 16 different EVM and non-EVM blockchains with more
              integrations coming soon
            </p>
          </li>
          <li className="basis-56 grow">
            <h2 className="text-4xl font-medium mb-4">Developer Friendly</h2>
            <p className="text-lg mb-2">
              Easily integrate cross-chain token bridging natively into your
              decentralized application
            </p>
          </li>
        </ul>
      </section> */}
      {/* <section className="w-full">
        <header className="p-4">
          <h2 className="text-4xl font-medium mb-4">Widely Integrated</h2>
          <p className="text-lg mb-2">
            Synapse is widely integrated across the most-used Layer 1 & 2
            networks for a seamless cross-chain experience.
          </p>
        </header>
        <ul className="grid grid-cols-3 md:grid-cols-4 lg:grid-cols-5 text-center gap-4 bg-white dark:bg-zinc-950 p-4 border border-zinc-300 dark:border-zinc-800 rounded-lg shadow-lg">
          {ChainList().map((a) => {
            return (
              <li className="border border-zinc-300 dark:border-zinc-700 rounded pb-4 pt-16">
                {a}
              </li>
            )
          })}
        </ul>
      </section> */}
      <section className="flex odd:flex-col even:flex-col-reverse md:grid grid-cols-2 gap-12 items-center">
        <div className="w-full h-full grid items-center justify-center">
          <svg
            width="100%"
            height="50%"
            viewBox="-4 -4 8 8"
            strokeWidth=".1"
            vectorEffect="non-scaling-size"
            fill="none"
            overflow="visible"
            className="stroke-fuchsia-500"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle r="4" />
          </svg>
        </div>
        <div>
          <h2 className="text-4xl font-medium my-4">Secure your routes</h2>
          <p className="text-lg leading-relaxed">
            Synapse has processed millions of transactions and tens of billions
            in bridged assets.
          </p>
        </div>
      </section>
    </article>
  )
}

/*
if (theme) {
    formStyle = `p-3 rounded-md border shadow-sm grid gap-4 absolute w-96 ${
      theme === 'dark'
        ? 'text-white bg-zinc-900 border-zinc-800 mr-8 mt-8'
        : 'text-black bg-neutral-100 border-zinc-300 ml-8 mb-8'
    }`
    sectionStyle = `rounded px-2.5 py-3 grid gap-3 grid-cols-2 ${
      theme === 'dark' ? 'bg-zinc-800' : 'bg-zinc-50 border border-zinc-200'
    }`
    selectStyle = `rounded w-fit cursor-pointer border ${
      theme === 'dark'
        ? 'bg-zinc-700 border-transparent'
        : 'bg-inherit border-zinc-300'
    }`
    inputWrapperStyle = `flex border rounded-md p-1.5 col-span-2 gap-1.5 ${
      theme === 'dark'
        ? 'bg-inherit border-zinc-700'
        : 'bg-white border-zinc-200 '
    }`
    inputStyle = `bg-inherit border-none w-full p-1.5 text-xxl font-normal tracking-wide rounded ${
      theme === 'dark' ? 'font-light' : 'font-normal'
    }`
  }
  */
