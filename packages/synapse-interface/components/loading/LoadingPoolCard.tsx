import Grid from '@tw/Grid'
import Card from '@tw/Card'

const LoadingPoolCard = () => {
  return (
    <Card
      title={<div className="px-6 h-5 bg-slate-700 rounded w-[60%]"></div>}
      titleClassName="text-white font-light text-xl"
      className={`
        items-center row-span-1
      bg-bgBase transition-all rounded-xl
        py-6 mt-4 pr-2 h-40
        border border-transparent
        hover:bg-bgLight animate-pulse
      `}
      divider={false}
    >
      <div className="h-4 my-2 mt-9 bg-slate-700 rounded w-[90%]"></div>
      <div className="h-4 my-2 bg-slate-700 rounded w-[90%]"></div>
    </Card>
  )
}
export default LoadingPoolCard
