import Tooltip from '@components/tailwind/Tooltip'

export function Indicator({ indicatorType }) {
  switch (indicatorType) {
    case 'pending':
      return (
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex justify-end">
          // @ts-expect-error TS(2749): 'Tooltip' refers to a value, but is being used as ... Remove this comment to see the full error message
          <Tooltip
            // @ts-expect-error TS(2304): Cannot find name 'content'.
            content="Transaction still in progress"
            // @ts-expect-error TS(2304): Cannot find name 'tooltipClassName'.
            tooltipClassName="!-mt-24 !-ml-24"
          >
            // @ts-expect-error TS(2304): Cannot find name 'span'.
            <span className="flex w-3 h-3 mr-1">
              // @ts-expect-error TS(2304): Cannot find name 'span'.
              <span className="absolute inline-flex w-3 h-3 rounded-full opacity-75 animate-ping bg-amber-400"></span>
              // @ts-expect-error TS(2304): Cannot find name 'span'.
              <span className="relative inline-flex w-3 h-3 rounded-full bg-amber-500"></span>
            </span>
          </Tooltip>
        </div>
      )
    case 'indexing':
      return (
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex justify-end">
          // @ts-expect-error TS(2749): 'Tooltip' refers to a value, but is being used as ... Remove this comment to see the full error message
          <Tooltip
            // @ts-expect-error TS(2304): Cannot find name 'content'.
            content="Transaction complete, data indexing"
            // @ts-expect-error TS(2304): Cannot find name 'tooltipClassName'.
            tooltipClassName="!-mt-24 !-ml-24"
          >
            // @ts-expect-error TS(2304): Cannot find name 'span'.
            <span className="flex w-3 h-3 mr-1">
              // @ts-expect-error TS(2304): Cannot find name 'span'.
              <span className="absolute inline-flex w-3 h-3 rounded-full opacity-75 animate-ping bg-slate-400"></span>
              // @ts-expect-error TS(2304): Cannot find name 'span'.
              <span className="relative inline-flex w-3 h-3 rounded-full bg-slate-500"></span>
            </span>
          </Tooltip>
        </div>
      )
    case 'complete':
      return (
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex justify-end">
          // @ts-expect-error TS(2749): 'Tooltip' refers to a value, but is being used as ... Remove this comment to see the full error message
          <Tooltip
            // @ts-expect-error TS(2304): Cannot find name 'content'.
            content="Transaction completed"
            // @ts-expect-error TS(2304): Cannot find name 'tooltipClassName'.
            tooltipClassName="!-mt-24 !-ml-24"
          >
            // @ts-expect-error TS(2304): Cannot find name 'span'.
            <span className="flex w-3 h-3 mr-1">
              // @ts-expect-error TS(2304): Cannot find name 'span'.
              <span className="relative inline-flex w-3 h-3 bg-green-400 rounded-full"></span>
            </span>
          </Tooltip>
        </div>
      )
    case 'incomplete':
      return (
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex justify-end">
          // @ts-expect-error TS(2749): 'Tooltip' refers to a value, but is being used as ... Remove this comment to see the full error message
          <Tooltip
            // @ts-expect-error TS(2304): Cannot find name 'content'.
            content="Transaction failed"
            // @ts-expect-error TS(2304): Cannot find name 'tooltipClassName'.
            tooltipClassName="!-mt-24 !-ml-24"
          >
            // @ts-expect-error TS(2304): Cannot find name 'span'.
            <span className="flex w-3 h-3 mr-1">
              // @ts-expect-error TS(2304): Cannot find name 'span'.
              <span className="relative inline-flex w-3 h-3 bg-red-500 rounded-full"></span>
            </span>
          </Tooltip>
        </div>
      )
  }
}
