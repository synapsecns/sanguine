import React, { useState, useEffect, useRef } from 'react'
import { SpeakerphoneIcon } from '@heroicons/react/outline'

import { titleCase } from '@/utils/titleCase'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'

const CHANGELOG_RAW_URL =
  'https://raw.githubusercontent.com/synapsecns/sanguine/fe-release/packages/synapse-interface/CHANGELOG.md'

const CHANGELOG_VISIBLE_URL =
  'https://github.com/synapsecns/sanguine/blob/fe-release/packages/synapse-interface/CHANGELOG.md'

interface Feature {
  version: string
  description: string
}

export const ToggleChangelogButton = () => {
  const ref = useRef(null)
  const [isChangelogVisible, setIsChangelogVisible] = useState(false)
  const [features, setFeatures] = useState([])
  const [isNewVersion, setIsNewVersion] = useState(false)

  useEffect(() => {
    fetch(CHANGELOG_RAW_URL)
      .then((response) => response.text())
      .then((text) => {
        const f = parseChangelog(text)
        setFeatures(f)
        const lastSeenVersion = localStorage.getItem('lastSeenVersion')
        if (f.length > 0 && f[0].version !== lastSeenVersion) {
          setIsNewVersion(true)
        }
      })
  }, [])

  const toggleChangelog = () => {
    setIsChangelogVisible(!isChangelogVisible)
    if (!isChangelogVisible && features.length > 0) {
      localStorage.setItem('lastSeenVersion', features[0].version)
      setIsNewVersion(false)
    }
  }

  useCloseOnOutsideClick(ref, () => setIsChangelogVisible(false))

  return (
    <div id="toggle-changelog" className="relative" ref={ref}>
      <button
        onClick={toggleChangelog}
        className="relative p-2 text-white border border-bgLight hover:bg-zinc-700 focus:outline-none focus:ring"
      >
        {isNewVersion && (
          <span className="absolute top-0 w-3 h-3 transform translate-x-1/2 -translate-y-1/2 rounded-full -left-3 bg-fuchsia-500" />
        )}
        <SpeakerphoneIcon className="w-6 h-6 text-secondary" />
      </button>
      {isChangelogVisible && (
        <div className="absolute right-0 z-50 mt-2 w-80">
          <ChangelogFeatures features={features} />
        </div>
      )}
    </div>
  )
}
const ChangelogFeatures = ({ features }: { features: Feature[] }) => {
  return (
    <div className="p-3 text-white border rounded-sm shadow-lg border-fuchsia-800 z-100 bg-gradient-to-r from-fuchsia-900 to-purple-900">
      <div className="mb-2 text-lg font-bold text-center">Latest Updates</div>
      <ul>
        {features.slice(0, 3).map((feature, index) => (
          <li
            key={index}
            className="p-2 mb-4 space-y-1 text-sm border"
            style={{
              borderImage:
                'linear-gradient(to right, rgba(255, 0, 255, 1), rgba(172, 143, 255, 1)) 1',
            }}
          >
            <div className="flex items-center space-x-2">
              <div className="px-[8px] py-[2px] bg-fuchsia-500 rounded-md">
                NEW
              </div>
              <div className="font-semibold">v{feature.version}</div>
            </div>
            <div>{titleCase(feature.description)}</div>
          </li>
        ))}
      </ul>
      <a
        className="text-sm hover:underline hover:text-blue-500"
        href={CHANGELOG_VISIBLE_URL}
        target="_blank"
      >
        See all changes
      </a>
    </div>
  )
}

function parseChangelog(changelogText) {
  const features = []
  const lines = changelogText.split('\n')

  let currentVersion = null
  let currentDescription = null

  for (const line of lines) {
    const versionMatch = line.match(/^# \[(\d+\.\d+\.\d+)\]/)
    if (versionMatch) {
      if (currentVersion && currentDescription) {
        features.push({
          version: currentVersion,
          description: currentDescription.trim(),
        })
      }
      currentVersion = versionMatch[1]
      currentDescription = ''
    } else if (line.startsWith('* **synapse-interface:**')) {
      const description = line.replace('* **synapse-interface:**', '').trim()
      const cleanedDescription = description.replace(
        /\((\[#\d+\]|\[\w+\])\(https?:\/\/[^)]+\)\)/g,
        ''
      )
      currentDescription += cleanedDescription.trim() + ' '
    }
  }

  if (currentVersion && currentDescription) {
    features.push({
      version: currentVersion,
      description: currentDescription.trim(),
    })
  }

  return features
}
