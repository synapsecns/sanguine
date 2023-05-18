import Button from '@components/tailwind/Button'
import {DotsHorizontalIcon} from '@heroicons/react/outline'

export default function MoreButton({ open, onClick, className, ...props }) {
  return (
    // @ts-expect-error TS(2749): 'Button' refers to a value, but is being used as a... Remove this comment to see the full error message
    <Button
      onClick={onClick}
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`
        w-full cursor-pointer rounded-lg px-4 py-4 group border-none dark:hover:bg-[#111111] ${className} text-sm
        ${open && ' bg-[#111111]'}
      `}
      // @ts-expect-error TS(2304): Cannot find name 'outline'.
      outline={true}
      // @ts-expect-error TS(2304): Cannot find name 'props'.
      {...props}
    >
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="space-x-2">
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="inline-block rounded-md">
          // @ts-expect-error TS(2749): 'DotsHorizontalIcon' refers to a value, but is bei... Remove this comment to see the full error message
          <DotsHorizontalIcon
            // @ts-expect-error TS(2304): Cannot find name 'className'.
            className={`
              ${
                open && 'opacity-100'
              } inline-block w-4 h-4 text-white dark:text-white group-hover:opacity-100
              `}
          // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
          />
        </div>
      </div>
    </Button>
  )
}
