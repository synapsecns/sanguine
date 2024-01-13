import Grid from '@tw/Grid'
import Link from 'next/link'

import {
  BRIDGE_PATH, BUILD_ON_URL,
  DISCORD_URL,
  DOCS_URL,
  FORUM_URL,
  GITHUB_URL,
  MIRROR_URL,
  POOLS_PATH,
  PRIVACY_POLICY_PATH,
  STAKE_PATH,
  SWAP_PATH, TELEGRAM_URL,
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
    url: DOCS_URL,
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
      <div className="p-10 flex flex-wrap gap-8 justify-between max-w-6xl m-auto">
          <SynapseTitleLogo showText={true} />
          <div className="flex flex-wrap gap-8">
            <FooterBlock elements={functionsList} />
            <FooterBlock elements={developersList} />
            <FooterBlock elements={supportList} />
          </div>
      </div>
      <div className="pb-12 flex gap-2 justify-center">
        <a
          className="opacity-60 hover:opacity-100"
          href={TERMS_OF_SERVICE_PATH}
          target="_blank"
          rel="noreferrer"
        >
          Terms of Use
        </a>
      <span className="opacity-60">|</span>
        <a
          className="opacity-60 hover:opacity-100"
          href={PRIVACY_POLICY_PATH}
          target="_blank"
          rel="noreferrer"
        >
          Privacy Policy
        </a>
      </div>
    </footer>
  )
}

function FooterBlock({ elements }: { elements: FooterDataProps[] }) {
  return (
    <div className="text-base leading-8 text-left md:text-left lg:text-right">
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
      <div>
        <a
          className="opacity-60 hover:opacity-100"
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
      <div>
        <Link
          className="opacity-60 hover:opacity-100"
          href={url}
        >
          {text}
        </Link>
      </div>
    )
  } else {
    return <>{text}</>
  }
}
