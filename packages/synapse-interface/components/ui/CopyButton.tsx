import React, { useState } from 'react'

const copyToClipboard = (text: string) => {
  navigator.clipboard.writeText(text)
}

export const CopyButton = ({ text }: { text: string }) => {
  const [isClicked, setIsClicked] = useState(false)

  const handleCopyClick = () => {
    copyToClipboard(text)
    setIsClicked(true)
  }

  return (
    <button
      onClick={() => handleCopyClick()}
      onMouseLeave={() => setIsClicked(false)}
      className="group/copy"
    >
      <span className="text-white/65 hover:opacity-100 group-hover/copy:underline active:opacity-65 py-0.5 inline-block">
        {isClicked ? 'copied' : 'copy'}
      </span>
    </button>
  )
}
