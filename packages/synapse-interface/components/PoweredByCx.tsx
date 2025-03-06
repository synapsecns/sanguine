'use client'

import React from 'react'
import Image from 'next/image'
import Link from 'next/link'

import { DOCS_CX_URL } from '@/constants/urls'

export const PoweredByCx = () => {
  return (
    <Link
      href={DOCS_CX_URL}
      className="flex gap-1.5 items-center text-xs text-secondaryTextColor hover:underline hover:text-white hover:bg-bgLighter rounded-md p-1 px-2 opacity-80 hover:opacity-100 transition-all"
      target="_blank"
      rel="noopener noreferrer"
    >
      <Image
        alt="Cortex icon"
        width="16"
        height="16"
        src="https://cortexprotocol.com/icon.svg"
        className="-translate-y-px"
      />
      Powered by Cortex
    </Link>
  )
}

export default PoweredByCx