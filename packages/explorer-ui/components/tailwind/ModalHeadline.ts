import CloseIcon from '@components/icons/CloseIcon'

export default function ModalHeadline({
  title,
  subtitle,
  onClose,
  closeIconClassName,
}) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className="flex">
      // @ts-expect-error TS(2304): Cannot find name 'h3'.
      <h3 className="pt-4 mb-3" id="modal-headline">
        // @ts-expect-error TS(2304): Cannot find name 'p'.
        <p className="mb-3 text-2xl font-medium leading-6 text-gray-900 dark:text-gray-300">
          {title}
        </p>
        // @ts-expect-error TS(2304): Cannot find name 'p'.
        <p className="text-gray-700 dark:text-gray-400">{subtitle}</p>
      </h3>
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className=" w-16 ml-auto cursor-pointer pt-1.5  -mr-2">
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className={`
            float-right hover:bg-gray-50 rounded-full p-1.5
            text-gray-600 hover:text-gray-900
            dark:text-gray-500 dark:hover:text-gray-300
            dark:hover:bg-gray-900
            ${closeIconClassName}
          `}
          // @ts-expect-error TS(2304): Cannot find name 'onClick'.
          onClick={onClose}
        >
          // @ts-expect-error TS(2749): 'CloseIcon' refers to a value, but is being used a... Remove this comment to see the full error message
          <CloseIcon />
        </div>
      </div>
    </div>
  )
}
