import {
  BRIDGE_PATH,
  TERMS_OF_SERVICE_PATH,
  DISCORD_URL,
  DOCS_URL,
  PRIVACY_POLICY_PATH,
  GITHUB_URL,
  MEDIUM_URL,
  POOLS_PATH,
  STAKE_PATH,
  SWAP_PATH,
  TWITTER_URL,
} from '@urls'

import Grid from '@components/tailwind/Grid'
import { SynapseTitleLogo } from '.'

const functions = [
  {
    text: 'Functions',
    type: null,
    url: null,
  },
  {
    text: 'Swap',
    type: 'url',
    url: SWAP_PATH,
  },
  {
    text: 'Bridge',
    type: 'url',
    url: BRIDGE_PATH,
  },
  {
    text: 'Pools',
    type: 'url',
    url: POOLS_PATH,
  },
  {
    text: 'Stake',
    type: 'url',
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
  }
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
  }
]


export function PageFooter() {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'footer'.
    <footer>
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="max-w-md px-4 pt-4 pb-6 mx-auto sm:max-w-3xl sm:pt-6 sm:px-6 lg:max-w-7xl lg:px-8">
        // @ts-expect-error TS(2749): 'Grid' refers to a value, but is being used as a t... Remove this comment to see the full error message
        <Grid
          // @ts-expect-error TS(2304): Cannot find name 'cols'.
          cols={{ xs: 2, sm: 2, md: 4, lg: 6 }}
          // @ts-expect-error TS(2304): Cannot find name 'gap'.
          gap={4}
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className="px-12 py-6 md:py-12 lg:py-12"
        >
          // @ts-expect-error TS(2304): Cannot find name 'div'.
          <div className="items-center hidden col-span-3 lg:flex">
            // @ts-expect-error TS(2749): 'SynapseTitleLogo' refers to a value, but is being... Remove this comment to see the full error message
            <SynapseTitleLogo showText={true} />
          </div>
          // @ts-expect-error TS(2304): Cannot find name 'elements'.
          <FooterBlock elements={functions} />
          // @ts-expect-error TS(2749): 'FooterBlock' refers to a value, but is being used... Remove this comment to see the full error message
          <FooterBlock elements={developers} />
          // @ts-expect-error TS(2304): Cannot find name 'elements'.
          <FooterBlock elements={support} />
        </Grid>
      </div>
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className='text-white pb-[70px] flex direction-row justify-center align-middle'>
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className='mr-2'> <DisplayText element={admin[0]} /> </div>   <p>｜</p>  <div className='ml-2'><DisplayText element={admin[1]} />  </div>
      </div>
    </footer>
  )
}

function FooterBlock({ elements }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className="text-left text-white md:text-left lg:text-right hover:cursor-pointer">
      {elements.map((element, i) => (
        // @ts-expect-error TS(2749): 'DisplayText' refers to a value, but is being used... Remove this comment to see the full error message
        <DisplayText element={element} key={i} />
      ))}
    </div>
  )
}

function DisplayText({ element }) {
  const { text, url, type } = element

  if (type === 'url') {
    return (
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="opacity-50 ">
        // @ts-expect-error TS(2304): Cannot find name 'a'.
        <a
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className="hover:opacity-40 hover:underline"
          // @ts-expect-error TS(2304): Cannot find name 'href'.
          href={url}
          // @ts-expect-error TS(2304): Cannot find name 'target'.
          target="_blank"
        >
          {text}
        </a>
      </div>
    )
  } else if (type === 'path') {
    return (
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="opacity-50">
        // @ts-expect-error TS(2304): Cannot find name 'a'.
        <a className="hover:opacity-40 hover:underline" href={url}>
          {text}
        </a>
      </div>
    )
  } else {
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    return <div className="opacity-80">{text}</div>
  }
}
