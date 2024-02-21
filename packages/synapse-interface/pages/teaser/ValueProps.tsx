import PulseDot from './PulseDot'

export default function ValueProps() {
  return (
    <article className="grid gap-24 px-4">
      <section className="flex even:flex-col odd:flex-col-reverse md:grid grid-cols-2 gap-12 items-center">
        <div className="grid place-items-center">
          <svg
            width="100%"
            height="100%"
            viewBox="-9 -9 170 118"
            vector-effect="non-scaling-stroke"
            stroke-width=".25"
            overflow="visible"
            // preserveAspectRatio="none"
            className="stroke-fuchsia-800/75 dark:stroke-purple-200/75 fill-fuchsia-200/50 dark:fill-fuchsia-300/5 drop-shadow-xl"
            xmlns="http://www.w3.org/2000/svg"
          >
            <rect
              width="168"
              height="116"
              x="-8"
              y="-8"
              rx="4"
              className="stroke-purple-400/80 fill-[hsl(275deg_100%_96%)] dark:fill-[hsl(300deg_40%_10%)]"
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
          <h2 className="text-4xl font-medium mb-4">
            Securely connect every blockchain
          </h2>
          <p className="text-lg leading-relaxed mb-3">
            Synapse enables developers to build truly native cross-chain
            applications with an economically secure method to reach consensus
            on interchain transactions,
          </p>
        </div>
      </section>
      <section className="flex even:flex-col odd:flex-col-reverse md:grid grid-cols-2 gap-12 items-center">
        <div>
          <h2 className="text-4xl font-medium mb-4">
            Powering the most popular bridge
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
      <section className="grid md:grid-cols-2 gap-x-12">
        <div>
          <h2 className="text-4xl font-medium my-4">
            Battle-tested infrastructure
          </h2>
          <p className="text-lg leading-relaxed">
            Synapse has processed millions of transactions and tens of billions
            in bridged assets.
          </p>
        </div>
        <div className="w-full h-full grid items-center justify-center">
          <svg
            width="100%"
            height="50%"
            viewBox="-4 -4 8 8"
            stroke-width=".1"
            vector-effect="non-scaling-size"
            fill="none"
            overflow="visible"
            className="stroke-fuchsia-500"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle r="4" />
          </svg>
        </div>
      </section>
      <section className="grid md:grid-cols-2 gap-x-12">
        <div>
          <h2 className="text-4xl font-medium my-4">Widely integrated</h2>
          <p className="text-lg leading-relaxed">
            Synapse is widely integrated across the most-used Layer 1 and Layer
            2 networks for a seamless cross-chain experience.
          </p>
        </div>
        <div className="w-full h-full grid items-center justify-center md:row-start-1 col-start-1 row-start-2">
          <svg
            width="100%"
            height="50%"
            viewBox="-4 -4 8 8"
            stroke-width=".1"
            vector-effect="non-scaling-size"
            fill="none"
            overflow="visible"
            className="stroke-fuchsia-500"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle r="4" />
          </svg>
        </div>
      </section>
    </article>
  )
}

const FauxBridge = () => {
  const cardStyle =
    'text-black dark:text-white bg-zinc-100 dark:bg-zinc-900 p-3 rounded-md border border-zinc-200 dark:border-zinc-800 shadow-xl grid gap-4 max-w-sm'
  const sectionStyle =
    'bg-zinc-50 dark:bg-zinc-800 rounded-md px-2.5 py-3 grid gap-3 grid-cols-2 border border-zinc-300 dark:border-transparent'
  const selectStyle =
    'rounded bg-inherit dark:bg-zinc-700 border-zinc-300 dark:border-zinc-700 w-fit cursor-pointer'
  const inputWrapperStyle =
    'flex bg-white dark:bg-inherit border border-zinc-200 dark:border-zinc-700 rounded-md gap-0 p-1.5 col-span-2 gap-1.5 items-center'
  const inputStyle =
    'bg-inherit border-none w-full p-1.5 text-xxl font-normal dark:font-light tracking-wide rounded'

  return (
    <div className={cardStyle}>
      <section className={sectionStyle}>
        <select className={`ml-0.5 ${selectStyle}`}>
          <option>Chain</option>
        </select>
        <div className="flex gap-2.5 items-center justify-self-end text-sm text-zinc-700 dark:text-zinc-300 mr-1 cursor-default">
          <PulseDot />
          <span className="mb-px">Connected</span>
        </div>
        <div className={inputWrapperStyle}>
          <select className={selectStyle}>
            <option>Token</option>
          </select>
          <input type="text" value="1000" className={inputStyle} />
          <button
            disabled
            className="px-4 py-1 bg-zinc-100 dark:bg-zinc-700 border border-zinc-200 dark:border-transparent h-fit rounded mr-1 cursor-pointer hover:bg-zinc-200 hover:dark:bg-zinc-600"
          >
            Max
          </button>
        </div>
      </section>
      <section className={sectionStyle}>
        <select className={`ml-0.5 ${selectStyle}`}>
          <option>Chain</option>
        </select>
        <div className={inputWrapperStyle}>
          <select className={selectStyle}>
            <option>Token</option>
          </select>
          <input
            disabled
            type="text"
            value="1,000"
            className={inputStyle}
          ></input>
        </div>
      </section>
      <a
        href="#"
        className="border border-fuchsia-500 p-2.5 rounded text-lg tracking-wider text-center hover:bg-purple-50 hover:dark:bg-fuchsia-950"
      >
        &nbsp;Bridge {'->'}
      </a>
    </div>
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
