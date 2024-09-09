import React, { useState, useRef, useEffect } from 'react'
import { useRouter } from 'next/router'
import { useTranslations } from 'next-intl'

import { GlobeAltIcon } from '@heroicons/react/outline'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { useCloseOnEscape } from '@/utils/hooks/useCloseOnEscape'

const languages = [
  { code: 'en-US', name: 'English' },
  { code: 'ar', name: 'العربية' },
  { code: 'es', name: 'Español' },
  { code: 'fr', name: 'Français' },
  { code: 'tr', name: 'Türkçe' },
]

export const LanguageSelector = () => {
  const router = useRouter()
  const t = useTranslations('LanguageSelector')
  const { pathname, asPath, query } = router
  const [isOpen, setIsOpen] = useState(false)
  const [currentLocale, setCurrentLocale] = useState(router.locale)
  const dropdownRef = useRef(null)

  useEffect(() => {
    const storedLanguage = localStorage.getItem('selectedLanguage')
    if (storedLanguage && storedLanguage !== router.locale) {
      router.push({ pathname, query }, asPath, { locale: storedLanguage })
    }
  }, [])

  useEffect(() => {
    setCurrentLocale(router.locale)
  }, [router.locale])

  const toggleDropdown = () => setIsOpen(!isOpen)

  const handleLanguageChange = (lang) => {
    router.push({ pathname, query }, asPath, { locale: lang.code })
    localStorage.setItem('selectedLanguage', lang.code)
  }

  const closeDropdown = () => setIsOpen(false)

  useCloseOnOutsideClick(dropdownRef, closeDropdown)
  useCloseOnEscape(closeDropdown)

  return (
    <div className="relative" ref={dropdownRef}>
      <button onClick={toggleDropdown} className="flex items-center">
        <GlobeAltIcon
          className={`w-6 h-6 hover:opacity-80 ${isOpen && 'opacity-80'}`}
        />
      </button>
      {isOpen && (
        <div className="absolute left-0 z-10 w-48 mt-2 bg-white rounded-sm ">
          <div
            className="py-1"
            role="menu"
            aria-orientation="vertical"
            aria-labelledby="options-menu"
          >
            <div className="px-2 mb-2 text-black ">{t('Language')}</div>
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
