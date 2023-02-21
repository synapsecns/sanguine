import {
  BRIDGE_PATH,
  CAREERS_URL,
  DISCORD_URL,
  DOCS_URL,
  FORUM_URL,
  GITHUB_URL,
  MEDIUM_URL,
  POOLS_PATH,
  STAKE_PATH,
  SWAP_PATH,
  TWITTER_URL,
} from '@urls'

import Grid from '@components/tailwind/Grid'
import {SynapseTitleLogo} from '.'

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
    text: 'Tutorials',
    type: 'url',
    url: MEDIUM_URL,
  },
  {
    text: 'GitHub',
    type: 'url',
    url: GITHUB_URL,
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
    text: 'Help Center',
    type: 'url',
    url: FORUM_URL,
  },
  {
    text: 'Careers',
    type: 'url',
    url: CAREERS_URL,
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
        <a className="hover:opacity-40 hover:underline" to={url}>
          {text}
        </a>
      </div>
    )
  } else {
    return <div className="opacity-80">{text}</div>
  }
}
