import CloseIcon from '@components/icons/CloseIcon'

export default function ModalHeadline({
  title,
  subtitle,
  onClose,
  closeIconClassName,
}) {
  return (
    <div className="flex">
      <h3 className="pt-4 mb-3" id="modal-headline">
        <p className="mb-3 text-2xl font-medium leading-6 text-gray-900 dark:text-gray-300">
          {title}
        </p>
        <p className="text-gray-700 dark:text-gray-400">{subtitle}</p>
      </h3>
      <div className=" w-16 ml-auto cursor-pointer pt-1.5  -mr-2">
        <div
          className={`
            float-right hover:bg-gray-50 rounded-full p-1.5
            text-gray-600 hover:text-gray-900
            dark:text-gray-500 dark:hover:text-gray-300
            dark:hover:bg-gray-900
            ${closeIconClassName}
          `}
          onClick={onClose}
        >
          <CloseIcon />
        </div>
      </div>
    </div>
  )
}
