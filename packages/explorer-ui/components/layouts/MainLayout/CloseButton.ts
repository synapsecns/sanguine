import CloseIcon from '@components/icons/CloseIcon'

export default function CloseButton({ onClick }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'button'.
    <button
      onClick={onClick}
      // @ts-expect-error TS(2304): Cannot find name 'className'.
      className={`
        group ml-1 flex items-center justify-center
        h-10 w-10 rounded-full
        focus:outline-none bg-white
        shadow-md hover:shadow-lg
        transition
        hover:bg-indigo-50
        text-indigo-600 hover:text-indigo-700


        dark:bg-gray-800
        dark:text-gray-500
        dark:hover:text-purple-700
        dark:hover:bg-gray-700 dark:active:bg-gray-700
        `}
    >
      // @ts-expect-error TS(2304): Cannot find name 'span'.
      <span className="sr-only">Close sidebar</span>
      // @ts-expect-error TS(2304): Cannot find name 'className'.
      <CloseIcon className="w-6 h-6 " />
    </button>
  )
}
