import { ApolloProvider } from '@apollo/client'
import { PageWrapper } from '@components/layouts//MainLayout'

import client from '../apollo-client'
import '../styles/global.css'

function MyApp({ Component, pageProps }) {
  return (
    <ApolloProvider client={client}>
      <PageWrapper>
        <Component {...pageProps} />
      </PageWrapper>
    </ApolloProvider>
  )
}

export default MyApp
