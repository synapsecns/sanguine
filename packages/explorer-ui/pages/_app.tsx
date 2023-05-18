import { ApolloProvider } from '@apollo/client'
import { PageWrapper } from '@components/layouts//MainLayout'
import { GoogleAnalytics } from "nextjs-google-analytics";

import client from '../apollo-client'
import '../styles/global.css'

function App({ Component, pageProps }) {
  return (
    <>
      <GoogleAnalytics trackPageViews />
      <ApolloProvider client={client}>
        <PageWrapper>
          <Component {...pageProps} />
        </PageWrapper>
      </ApolloProvider>
    </>
  )
}

export default App
