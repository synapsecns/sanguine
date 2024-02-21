import PulseDot from './PulseDot'

export default function ValueProps() {
  return (
    <article className="grid gap-8 px-4">
      <section className="grid md:grid-cols-2 gap-x-12">
        <div>
          <h2 className="text-4xl font-medium my-4">
            Securely connect every blockchain
          </h2>
          <p className="text-lg leading-relaxed">
            Synapse is comprised of a cross-chain messaging framework and an
            economically secure method to reach consensus on the validity of
            cross-chain transactions, enabling developers to build truly native
            cross-chain apps.
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
          <h2 className="text-4xl font-medium my-4">
            Powering the most popular bridge
          </h2>
          <p className="text-lg leading-relaxed">
            Synapse Bridge is built on top of the cross-chain infrastructure
            enabling users to seamlessly transfer assets across all blockchains.
            The Bridge has become the most widely-used method to move assets
            cross-chain, offering low cost, fast, and secure bridging.
          </p>
        </div>
        <div className="w-full h-full grid items-center justify-center md:row-start-1 col-start-1 row-start-2 relative">
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
  const formStyle =
    'text-black dark:text-white bg-zinc-100 dark:bg-zinc-900 p-3 rounded-md border border-zinc-200 dark:border-zinc-800 shadow-xl grid gap-4 w-96'
  const sectionStyle =
    'bg-zinc-50 dark:bg-zinc-800 rounded-md px-2.5 py-3 grid gap-3 grid-cols-2 border border-zinc-300 dark:border-transparent'
  const selectStyle =
    'rounded bg-inherit dark:bg-zinc-700 border-zinc-300 dark:border-zinc-700 w-fit cursor-pointer'
  const inputWrapperStyle =
    'flex bg-white dark:bg-inherit border border-zinc-200 dark:border-zinc-700 rounded-md gap-0 p-1.5 col-span-2 gap-1.5 items-center'
  const inputStyle =
    'bg-inherit border-none w-full p-1.5 text-xxl font-normal dark:font-light tracking-wide rounded'

  return (
    <form className={formStyle}>
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
    </form>
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
