import React, { useState, useRef } from 'react'
import { useRouter } from 'next/router'
import { GlobeAltIcon } from '@heroicons/react/outline'

import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { useCloseOnEscape } from '@/utils/hooks/useCloseOnEscape'

const languages = [
  { code: 'en-US', name: 'English' },
  { code: 'fr', name: 'Français' },
  { code: 'lorem-ipsum', name: 'Lorem Ipsum' },
]

export const LanguageSelector = () => {
  const router = useRouter()
  const { pathname, asPath, query, locale: currentLocale } = router
  const [isOpen, setIsOpen] = useState(false)
  const dropdownRef = useRef(null)

  const toggleDropdown = () => setIsOpen(!isOpen)

  const handleLanguageChange = (lang) => {
    router.push({ pathname, query }, asPath, { locale: lang.code })
  }

  const closeDropdown = () => setIsOpen(false)

  useCloseOnOutsideClick(dropdownRef, closeDropdown)
  useCloseOnEscape(closeDropdown)

  return (
    <div className="relative" ref={dropdownRef}>
      <button onClick={toggleDropdown} className="flex items-center">
        <GlobeAltIcon className="w-6 h-6" />
      </button>
      {isOpen && (
        <div className="absolute right-0 z-10 w-48 mt-2 bg-white rounded-md shadow-lg ring-1 ring-black ring-opacity-5">
          <div
            className="py-1"
            role="menu"
            aria-orientation="vertical"
            aria-labelledby="options-menu"
          >
            {languages.map((lang) => (
              <div
                key={lang.code}
                className={`block px-4 py-2 text-sm cursor-pointer ${
                  currentLocale === lang.code
                    ? 'bg-blue-100 text-blue-900'
                    : 'text-gray-700 hover:bg-gray-100 hover:text-gray-900'
                }`}
                role="menuitem"
                onClick={() => handleLanguageChange(lang)}
                onKeyDown={(e) => {
                  if (e.key === 'Enter' || e.key === ' ') {
                    handleLanguageChange(lang)
                  }
                }}
                tabIndex={0}
              >
                {lang.name}
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  )
}
