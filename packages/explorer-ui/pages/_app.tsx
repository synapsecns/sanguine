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
          <Component {...pageProps} />
          <Analytics />
        </PageWrapper>
      </ApolloProvider>
    </>
  )
}

export default App
