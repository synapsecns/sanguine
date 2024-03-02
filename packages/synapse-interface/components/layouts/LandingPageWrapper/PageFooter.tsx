import Grid from '@tw/Grid'
import Link from 'next/link'

import {
  BRIDGE_PATH,
  BUILD_ON_URL,
  DISCORD_URL,
  SYNAPSE_DOCS_URL,
  FORUM_URL,
  GITHUB_URL,
  MIRROR_URL,
  POOLS_PATH,
  PRIVACY_POLICY_PATH,
  STAKE_PATH,
  SWAP_PATH,
  TELEGRAM_URL,
  TERMS_OF_SERVICE_PATH,
  TWITTER_URL,
} from '@/constants/urls'
import { SynapseTitleLogo } from '.'

enum FooterType {
  PATH = 'path',
  URL = 'url',
}

interface FooterDataProps {
  text: string
  type: string | null
  url: string | null
}

const functionsList: FooterDataProps[] = [
  {
    text: 'Functions',
    type: null,
    url: null,
  },
  {
    text: 'Swap',
    type: FooterType.PATH,
    url: SWAP_PATH,
  },
  {
    text: 'Bridge',
    type: FooterType.PATH,
    url: BRIDGE_PATH,
  },
  {
    text: 'Pools',
    type: FooterType.PATH,
    url: POOLS_PATH,
  },
  {
    text: 'Stake',
    type: FooterType.PATH,
    url: STAKE_PATH,
  },
]

const developersList: FooterDataProps[] = [
  {
    text: 'Developers',
    type: null,
    url: null,
  },
  {
    text: 'Build on Synapse',
    type: FooterType.URL,
    url: BUILD_ON_URL,
  },
  {
    text: 'Documentation',
    type: FooterType.URL,
    url: SYNAPSE_DOCS_URL,
  },
  {
    text: 'GitHub',
    type: FooterType.URL,
    url: GITHUB_URL,
  },
  {
    text: 'Blog',
    type: FooterType.URL,
    url: MIRROR_URL,
  },
]

const supportList: FooterDataProps[] = [
  {
    text: 'Support',
    type: null,
    url: null,
  },
  {
    text: 'Discord',
    type: FooterType.URL,
    url: DISCORD_URL,
  },
  {
    text: 'Twitter',
    type: FooterType.URL,
    url: TWITTER_URL,
  },
  {
    text: 'Forum',
    type: FooterType.URL,
    url: FORUM_URL,
  },
  {
    text: 'Telegram',
    type: 'url',
    url: TELEGRAM_URL,
  },
]

export function PageFooter() {
  return (
    <footer>
      <div className="flex flex-wrap justify-between max-w-4xl gap-8 p-8 m-auto">
        <SynapseTitleLogo showText={true} />
        <div className="flex flex-wrap gap-8">
          <FooterBlock elements={functionsList} />
          <FooterBlock elements={developersList} />
          <FooterBlock elements={supportList} />
        </div>
      </div>
      <div className="flex justify-center gap-2 pb-12 text-secondaryTextColor">
        <div className="text-opacity-50 text-secondaryTextColor">
          <a
            className="duration-75 hover:text-white transform-gpu hover:transition-all"
            href={TERMS_OF_SERVICE_PATH}
            target="_blank"
            rel="noreferrer"
          >
            Terms of Use
          </a>
        </div>
        <p>｜</p>
        <div className="text-opacity-50 text-secondaryTextColor">
          <a
            className="duration-75 hover:text-white transform-gpu hover:transition-all"
            href={PRIVACY_POLICY_PATH}
            target="_blank"
            rel="noreferrer"
          >
            Privacy Policy
          </a>
        </div>
      </div>
    </footer>
  )
}

function FooterBlock({ elements }: { elements: FooterDataProps[] }) {
  return (
    <div className="text-base leading-8 text-left text-white md:text-left lg:text-right">
      {elements.map((element, i) => (
        <DisplayText element={element} key={i} />
      ))}
    </div>
  )
}

function DisplayText({ element }: { element: FooterDataProps }) {
  const { text, url, type } = element

  if (type === FooterType.URL) {
    return (
      <div className="text-opacity-50 text-secondaryTextColor">
        <a
          className="duration-75 hover:text-white hover:text-opacity-100 transform-gpu hover:transition-all"
          href={url}
          target="_blank"
          rel="noreferrer"
        >
          {text}
        </a>
      </div>
    )
  } else if (type === FooterType.PATH) {
    return (
      <div className="text-opacity-50 text-secondaryTextColor">
        <Link
          className="duration-75 hover:text-white hover:text-opacity-100 transform-gpu hover:transition-all"
          href={url}
        >
          {text}
        </Link>
      </div>
    )
  } else {
    return <div className="text-opacity-80">{text}</div>
  }
}
