import numeral from 'numeral'

export const Chart = ({ data }) => {
  if (data) {
    const numbers = normalize(data)
    return (
      <div className="flex flex-col items-center w-full pb-6 rounded-lg shadow-xl sm:p-8">
        <div className="flex items-end flex-grow w-full mt-2 content-between">
          {numbers.map(({ value, normalizedValue, date }, i) => (
            <BarMaker
              value={value}
              height={normalizedValue}
              date={date}
              key={i}
            />
          ))}
        </div>
      </div>
    )
  }
}
const BarMaker = ({ value, height, date }) => {
  const h = `h-[${height}px]`
  const showValue = numeral(value).format('0,0')

  return (
    <div className="relative flex flex-col items-center flex-grow pb-5 ml-1 mr-1 group">
      <span className="absolute top-0 z-10 hidden -mt-6 text-xs text-white group-hover:block">
        {showValue}
      </span>
      <div
        className={`relative flex justify-center w-full ${h} bg-gradient-to-b from-[#FF00FF] to-[#AC8FFF] hover:opacity-50`}
      ></div>
      <span className="-rotate-45 text-white text-[5px] mt-3 l-0 pr-0">
        {date}
      </span>
    </div>
  )
}

export const ChartLoading = () => {
  return (
    <div className="flex flex-col items-center w-full pb-6 rounded-lg shadow-xl sm:p-8">
      <div className="flex items-end flex-grow w-full mt-2 content-between">
        {Array.from(Array(30).keys()).map((i) => (
          <BarMakerLoading key={i} />
        ))}
      </div>
    </div>
  )
}

function BarMakerLoading() {
  return (
    <div className="relative flex flex-col items-center flex-grow pb-5 ml-1 mr-1 group">
      <div
        className={`relative flex justify-center w-full h-[200px] animate-pulse bg-gradient-to-b from-slate-700 to-slate-500 hover:opacity-50`}
      ></div>
    </div>
  )
}

function normalize(data) {
  const maxHeight = 300

  let max = 0
  data.map((entry) => (entry.total > max ? (max = entry.total) : null))

  const newList = data.map((day) => {
    const date = new Date(day.date)
    const formattedDate = date.getUTCMonth() + 1 + '/' + date.getUTCDate()
    const n = (day.total / max) * maxHeight
    return {
      value: day.total,
      date: formattedDate,
      normalizedValue: Math.trunc(n),
    }
  })

  return newList
}
