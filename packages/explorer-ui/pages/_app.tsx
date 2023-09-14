import { ApolloProvider } from '@apollo/client'
import { PageWrapper } from '@components/layouts//MainLayout'
import { GoogleAnalytics } from 'nextjs-google-analytics'
import { Analytics } from '@vercel/analytics/react'

import client from '../apollo-client'
import '../styles/global.css'

function App({ Component, pageProps }) {
  return (
    <>
      <GoogleAnalytics trackPageViews />
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
