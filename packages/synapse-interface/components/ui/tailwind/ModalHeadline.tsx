import { twMerge } from 'tailwind-merge'

const baseClassname = `
  mb-3 text-sm text-secondaryTextColor text-opacity-50
`

export default function ModalHeadline({
  title,
  subtitle,
  onClose,
  titleClassName,
}: {
  title: string
  subtitle: string
  onClose: any
  titleClassName?: string
}) {
  const mergedTitleClassName = twMerge(`${baseClassname} ${titleClassName}`)

  return (
    <div>
      <div className="flex items-center">
        <h3 className="pt-3" id="modal-headline">
          <p className={mergedTitleClassName}>{title}</p>
        </h3>
        <div className="ml-auto cursor-pointer ">
          <div
            className="float-right text-sm text-red-500 hover:underline"
            onClick={onClose}
          >
            Clear
          </div>
        </div>
      </div>
      <p className="text-gray-400">{subtitle}</p>
    </div>
  )
}
