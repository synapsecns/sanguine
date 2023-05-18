import { ApolloProvider } from '@apollo/client'
import { PageWrapper } from '@components/layouts//MainLayout'

import client from '../apollo-client'
import '../styles/global.css'
import { useAnalyticsReporter } from "@components/analytics";

function App({ Component, pageProps }) {
  useAnalyticsReporter()

  return (
    <ApolloProvider client={client}>
      <PageWrapper>
        <Component {...pageProps} />
      </PageWrapper>
    </ApolloProvider>
  )
}

export default App
