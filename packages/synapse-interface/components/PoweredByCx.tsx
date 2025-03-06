'use client'

import React from 'react'
import Image from 'next/image'
import Link from 'next/link'

import { DOCS_CX_URL } from '@/constants/urls'

export const PoweredByCx = () => {
  return (
    <Link
      href={DOCS_CX_URL}
      className="flex gap-2 items-center text-sm text-secondaryTextColor hover:underline hover:text-white hover:bg-bgLighter rounded-md p-1.5 px-2.5 opacity-80 hover:opacity-100 transition-all"
      target="_blank"
      rel="noopener noreferrer"
    >
      <Image
        alt="Cortex icon"
        width="18"
        height="18"
        src="https://cortexprotocol.com/icon.svg"
        className="-translate-y-px"
      />
      Powered by Cortex
    </Link>
  )
}

export default PoweredByCx