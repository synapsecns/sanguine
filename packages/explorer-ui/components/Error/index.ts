import {ContainerCard} from '@components/ContainerCard'
import {ExclamationIcon} from '@heroicons/react/outline'

export const Error = ({ text, param, subtitle }) => {
  return (
    // @ts-expect-error TS(2749): 'ContainerCard' refers to a value, but is being us... Remove this comment to see the full error message
    <ContainerCard
      // @ts-expect-error TS(2304): Cannot find name 'className'.
      className="px-10 mt-10"
      // @ts-expect-error TS(2304): Cannot find name 'icon'.
      icon={<ExclamationIcon className="w-5 h-5 text-red-500" />}
      // @ts-expect-error TS(2304): Cannot find name 'subtitle'.
      subtitle={subtitle}
    >
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="mt-3 text-white">
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="mt-5 mb-2 font-extralight">{text}</div>
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="text-base text-gray-400 break-words font font-extralight">
          // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
          {param}
        </div>
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="mt-2 font-extralight">
          // @ts-expect-error TS(2304): Cannot find name 'Please'.
          Please click{' '}
          // @ts-expect-error TS(2304): Cannot find name 'a'.
          <a href={'/'} className="text-gray-400 hover:underline">
            // @ts-expect-error TS(2304): Cannot find name 'here'.
            here
          </a>{' '}
          // @ts-expect-error TS(2304): Cannot find name 'to'.
          to go back to the main page.
        </div>
      </div>
    </ContainerCard>
  )
}
