import './styles/global.css'
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import {PageWrapper} from '@components/layouts/MainLayout'
import {ApolloProvider} from '@apollo/client'
import client from './graphql/apollo-client'
import {ScrollToTop} from '@components/ScrollToTop'

const root = ReactDOM.createRoot(document.getElementById('root'))

root.render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <BrowserRouter>
        <ScrollToTop />
        <PageWrapper>
          <App />
        </PageWrapper>
      </BrowserRouter>
    </ApolloProvider>
  </React.StrictMode>
)
