import numeral from 'numeral'

export function Chart({ data }) {
  if (data) {
    const dailies = data.map((entry) => entry.total)
    const numbers = normalize(dailies)

    return (
      <div className="flex flex-col items-center w-full max-w-screen-md p-6 pb-6 rounded-lg shadow-xl sm:p-8">
        <div className="flex items-end flex-grow w-full mt-2 space-x-2 sm:space-x-3">
          {numbers.map(({ value, normalizedValue }, i) => (
            <BarMaker value={value} height={normalizedValue} key={i} />
          ))}
        </div>
      </div>
    )
  }
}

function BarMaker({ value, height }) {
  let h = `h-[${height}px]`
  let showValue = numeral(value).format('0,0')

  return (
    <div className="relative flex flex-col items-center flex-grow pb-5 group">
      <span className="absolute top-0 z-10 hidden -mt-6 text-xs text-white group-hover:block">
        {showValue}
      </span>
      <div
        className={`relative flex justify-center w-full ${h} bg-gradient-to-b from-[#FF00FF] to-[#AC8FFF] hover:opacity-50`}
      ></div>
    </div>
  )
}

export function ChartLoading() {
  return (
    <div className="flex flex-col items-center w-full max-w-screen-md p-6 pb-6 rounded-lg shadow-xl sm:p-8">
      <div className="flex items-end flex-grow w-full mt-2 space-x-2 sm:space-x-3">
        {[...Array(30).keys()].map((i) => (
          <BarMakerLoading key={i} />
        ))}
      </div>
    </div>
  )
}

function BarMakerLoading() {
  return (
    <div className="relative flex flex-col items-center flex-grow pb-5 group">
      <div
        className={`relative flex justify-center w-full h-[200px] animate-pulse bg-gradient-to-b from-slate-700 to-slate-500 hover:opacity-50`}
      ></div>
    </div>
  )
}

function normalize(list) {
  let maxHeight = 300

  let max = _.max(list)

  let newList = list.map((num) => {
    let n = (num / max) * maxHeight
    return {
      value: num,
      normalizedValue: Math.trunc(n),
    }
  })

  return newList
}
