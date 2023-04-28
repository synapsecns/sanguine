import Grid from '@tw/Grid'
import Link from 'next/link'

import {
  BRIDGE_PATH,
  DISCORD_URL,
  DOCS_URL,
  FORUM_URL,
  GITHUB_URL,
  MIRROR_URL,
  POOLS_PATH,
  PRIVACY_POLICY_PATH,
  STAKE_PATH,
  SWAP_PATH,
  TERMS_OF_SERVICE_PATH,
  TWITTER_URL,
} from '@/constants/urls'
import { SynapseTitleLogo } from '.'

const functions = [
  {
    text: 'Functions',
    type: null,
    url: null,
  },
  {
    text: 'Swap',
    type: 'path',
    url: SWAP_PATH,
  },
  {
    text: 'Bridge',
    type: 'path',
    url: BRIDGE_PATH,
  },
  {
    text: 'Pools',
    type: 'path',
    url: POOLS_PATH,
  },
  {
    text: 'Stake',
    type: 'path',
    url: STAKE_PATH,
  },
]

const developers = [
  {
    text: 'Developers',
    type: null,
    url: null,
  },
  {
    text: 'References',
    type: 'url',
    url: GITHUB_URL,
  },
  {
    text: 'Documentation',
    type: 'url',
    url: DOCS_URL,
  },
  {
    text: 'GitHub',
    type: 'url',
    url: GITHUB_URL,
  },
  {
    text: 'Blog',
    type: 'url',
    url: MIRROR_URL,
  },
]

const support = [
  {
    text: 'Support',
    type: null,
    url: null,
  },
  {
    text: 'Discord',
    type: 'url',
    url: DISCORD_URL,
  },
  {
    text: 'Twitter',
    type: 'url',
    url: TWITTER_URL,
  },
  {
    text: 'Forum',
    type: 'url',
    url: FORUM_URL,
  },
  //   {
  //     text: 'Careers',
  //     type: 'url',
  //     url: CAREERS_URL,
  //   },
]

export function PageFooter() {
  return (
    <footer>
      <div className="max-w-md px-6 pt-4 pb-6 mx-auto md:max-w-5xl md:px-8">
        <Grid
          cols={{ xs: 2, sm: 2, md: 4, lg: 6 }}
          gap={4}
          className="px-8 py-6 md:py-12"
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
        <div className="text-opacity-50 text-secondaryTextColor mr-2">
          <a
            className="duration-75 hover:text-white hover:text-opacity-100 transform-gpu hover:transition-all"
            href={TERMS_OF_SERVICE_PATH}
            target="_blank"
            rel="noreferrer"
          >
            Terms of Use
          </a>
        </div>
        <p>｜</p>
        <div className="text-opacity-50 text-secondaryTextColor ml-2">
          <a
            className="duration-75 hover:text-white hover:text-opacity-100 transform-gpu hover:transition-all"
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

function FooterBlock({ elements }: { elements: any[] }) {
  return (
    <div className="text-base leading-8 text-left text-white md:text-left lg:text-right">
      {elements.map((element, i) => (
        <DisplayText element={element} key={i} />
      ))}
    </div>
  )
}

function DisplayText({ element }: { element: any }) {
  const { text, url, type } = element

  if (type === 'url') {
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
  } else if (type === 'path') {
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
    return <div className="opacity-80">{text}</div>
  }
}
