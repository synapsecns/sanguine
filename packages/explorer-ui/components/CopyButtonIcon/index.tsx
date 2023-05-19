import { ClipboardCheckIcon, ClipboardCopyIcon } from '@heroicons/react/outline'
import { useState } from 'react'
import Tooltip from '@components/tailwind/Tooltip'

export function CopyButtonIcon({
                                 tooltipText = '',
                                 className,
                                 text = 'Text to Copy',
                               }) {
  const [copied, setCopied] = useState(false)

  const copyToClipboard = () => {
    navigator.clipboard.writeText(text).then(
      () => {
        setCopied(true)
        setTimeout(() => {
          setCopied(false)
        }, 5000)
      },
      (err) => {
        console.log(err)
      }
    )
  }

  const toCopy = ({ tooltipText, className = 'text-white' }) => {
    return (
      <Tooltip content={`Copy ${tooltipText}`} tooltipClassName="!-mt-16 !-ml-16">
        <ClipboardCopyIcon className={`w-5 h-5 ${className}`} strokeWidth={1} />
      </Tooltip>
    )
  }

  const hasBeenCopied = ({ tooltipText, className = 'text-green-200' }) => {
    return (
      <Tooltip
        content={`Copied ${tooltipText}`}
        tooltipClassName="!-mt-16 !-ml-16"
      >
        <ClipboardCheckIcon className={`w-5 h-5 ${className}`} strokeWidth={1} />
      </Tooltip>
    )
  }

  return (
    <>
      <button onClick={copyToClipboard}>
        {copied
          ? hasBeenCopied({ tooltipText, className })
          : toCopy({ tooltipText, className })}
      </button>
    </>
  )
}
