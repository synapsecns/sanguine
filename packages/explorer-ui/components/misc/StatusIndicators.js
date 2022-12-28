import Tooltip from '@components/tailwind/Tooltip'

export function Indicator({ indicatorType }) {
  switch (indicatorType) {
    case 'pending':
      return (
        <div className="flex justify-end">
          <Tooltip
            content="Transaction still in progress"
            tooltipClassName="!-mt-24 !-ml-24"
          >
            <span className="flex w-3 h-3 mr-1">
              <span className="absolute inline-flex w-3 h-3 rounded-full opacity-75 animate-ping bg-amber-400"></span>
              <span className="relative inline-flex w-3 h-3 rounded-full bg-amber-500"></span>
            </span>
          </Tooltip>
        </div>
      )
    case 'indexing':
      return (
        <div className="flex justify-end">
          <Tooltip
            content="Transaction complete, data indexing"
            tooltipClassName="!-mt-24 !-ml-24"
          >
            <span className="flex w-3 h-3 mr-1">
              <span className="absolute inline-flex w-3 h-3 rounded-full opacity-75 animate-ping bg-slate-400"></span>
              <span className="relative inline-flex w-3 h-3 rounded-full bg-slate-500"></span>
            </span>
          </Tooltip>
        </div>
      )
    case 'complete':
      return (
        <div className="flex justify-end">
          <Tooltip
            content="Transaction completed"
            tooltipClassName="!-mt-24 !-ml-24"
          >
            <span className="flex w-3 h-3 mr-1">
              <span className="relative inline-flex w-3 h-3 bg-green-400 rounded-full"></span>
            </span>
          </Tooltip>
        </div>
      )
    case 'incomplete':
      return (
        <div className="flex justify-end">
          <Tooltip
            content="Transaction failed"
            tooltipClassName="!-mt-24 !-ml-24"
          >
            <span className="flex w-3 h-3 mr-1">
              <span className="relative inline-flex w-3 h-3 bg-red-500 rounded-full"></span>
            </span>
          </Tooltip>
        </div>
      )
  }
}
