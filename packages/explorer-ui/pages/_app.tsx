import Head from 'next/head'
import { ApolloProvider } from '@apollo/client'
import { PageWrapper } from '@components/layouts//MainLayout'
import { GoogleAnalytics } from '@next/third-parties/google'
import { Analytics } from '@vercel/analytics/react'

import client from '../apollo-client'
import '../styles/global.css'

const App = ({ Component, pageProps }) => {
  return (
    <>
      <Head>
        <title>Synapse Explorer</title>
      </Head>
      <GoogleAnalytics gaId={'G-BBC13LQXBD'} />
      <ApolloProvider client={client}>
        <PageWrapper>
          {/* <Banner /> */}
          <Component {...pageProps} />
          <Analytics />
        </PageWrapper>
      </ApolloProvider>
    </>
  )
}

// const Banner = () => {
//   return (
//     <div className="flex items-center justify-center w-full h-12 p-4 text-white bg-purple-700">
//       We're updating the explorer at the moment, and some data may be
//       inaccurate. Thank you for your patience.
//     </div>
//   )
// }

export default App
