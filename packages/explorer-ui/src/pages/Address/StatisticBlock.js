import { GetStatistic } from './GetStatistic'

export function StatisticBlock({
  title,
  logo,
  address,
  type,
  duration,
  prefix,
}) {
  return (
    <div className="flex flex-col text-center ">
      <dd className="text-3xl font-bold text-slate-300">
        {
          <GetStatistic
            address={address}
            type={type}
            prefix={prefix}
            duration={duration}
          />
        }
      </dd>
      <dt className="text-gray-500 ">
        {/* <div className="flex items-center"> */}
        <span className="inline mr-2 align-middle">{logo}</span>
        <span className="text-sm">{title}</span>
        {/* </div> */}
      </dt>
    </div>
  )
}
