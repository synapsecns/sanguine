import {
  BRIDGE_PATH,
  TERMS_OF_SERVICE_PATH,
  DISCORD_URL,
  DOCS_URL,
  PRIVACY_POLICY_PATH,
  GITHUB_URL,
  POOLS_PATH,
  STAKE_PATH,
  SWAP_PATH,
  TWITTER_URL,
  BUILD_ON_URL,
  MIRROR_URL, FORUM_URL, TELEGRAM_URL,
} from '@urls'
import Grid from '@components/tailwind/Grid'

import { SynapseTitleLogo } from '.'

interface FooterDataProps {
  text: string
  type: string | null
  url: string | null
}

enum FooterType {
  PATH = 'path',
  URL = 'url',
}

const functions: FooterDataProps[] = [
  {
    text: 'Functions',
    type: null,
    url: null,
  },
  {
    text: 'Swap',
    type: FooterType.URL,
    url: SWAP_PATH,
  },
  {
    text: 'Bridge',
    type: FooterType.URL,
    url: BRIDGE_PATH,
  },
  {
    text: 'Pools',
    type: FooterType.URL,
    url: POOLS_PATH,
  },
  {
    text: 'Stake',
    type: FooterType.URL,
    url: STAKE_PATH,
  },
]

const developers: FooterDataProps[] = [
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

const support: FooterDataProps[] = [
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

const admin = [
  {
    text: 'Terms of Use',
    type: 'path',
    url: TERMS_OF_SERVICE_PATH,
  },
  {
    text: 'Privacy Policy',
    type: 'path',
    url: PRIVACY_POLICY_PATH,
  },
]

export function PageFooter() {
  return (
    <footer>
      <div className="max-w-md px-4 pt-4 pb-6 mx-auto sm:max-w-3xl sm:pt-6 sm:px-6 lg:max-w-7xl lg:px-8">
        <Grid
          cols={{ xs: 2, sm: 2, md: 4, lg: 6 }}
          gap={4}
          className="px-12 py-6 md:py-12 lg:py-12"
        >
          <div className="items-center hidden col-span-3 lg:flex">
            <SynapseTitleLogo showText={true} />
          </div>
          <FooterBlock elements={functions} />
          <FooterBlock elements={developers} />
          <FooterBlock elements={support} />
        </Grid>
      </div>
      <div className="text-white pb-[70px] flex direction-row justify-center align-middle">
        <div className="mr-2">
          {' '}
          <DisplayText element={admin[0]} />{' '}
        </div>{' '}
        <p>｜</p>{' '}
        <div className="ml-2">
          <DisplayText element={admin[1]} />{' '}
        </div>
      </div>
    </footer>
  )
}

function FooterBlock({ elements }) {
  return (
    <div className="text-left text-white md:text-left lg:text-right hover:cursor-pointer">
      {elements.map((element, i) => (
        <DisplayText element={element} key={i} />
      ))}
    </div>
  )
}

function DisplayText({ element }) {
  const { text, url, type } = element

  if (type === 'url') {
    return (
      <div className="opacity-50 ">
        <a
          className="hover:opacity-40 hover:underline"
          href={url}
          target="_blank"
        >
          {text}
        </a>
      </div>
    )
  } else if (type === 'path') {
    return (
      <div className="opacity-50">
        <a className="hover:opacity-40 hover:underline" href={url}>
          {text}
        </a>
      </div>
    )
  } else {
    return <div className="opacity-80">{text}</div>
  }
}
