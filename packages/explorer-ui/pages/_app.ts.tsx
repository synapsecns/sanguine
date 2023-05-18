import {ApolloProvider} from '@apollo/client'
import client from '../apollo-client'
import {PageWrapper} from '@components/layouts//MainLayout'
import '../styles/global.css'

function MyApp({ Component, pageProps }) {
  return (
    // @ts-expect-error TS(2749): 'ApolloProvider' refers to a value, but is being u... Remove this comment to see the full error message
    <ApolloProvider client={client}>
      // @ts-expect-error TS(2749): 'PageWrapper' refers to a value, but is being used... Remove this comment to see the full error message
      <PageWrapper >
        // @ts-expect-error TS(2749): 'Component' refers to a value, but is being used a... Remove this comment to see the full error message
        <Component {...pageProps} />
      </PageWrapper>
    </ApolloProvider>
  )
}

export default MyApp
