import React from 'react'
import clsx from 'clsx'
import Link from '@docusaurus/Link'
import useDocusaurusContext from '@docusaurus/useDocusaurusContext'
import Layout from '@theme/Layout'

import styles from './index.module.css'
import Card from '../components/Card'

export const HomepageHeader = () => {
  const { siteConfig } = useDocusaurusContext()
  return (
    <header className={clsx('hero hero--primary')}>
      <div className="container">
        <h1 className="hero__title">{siteConfig.title}</h1>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons}>
          <Link className="button button--secondary button--lg" to="/docs/sdk">
            SDK Tutorial - 5min ⏱️
          </Link>
        </div>
      </div>
    </header>
  )
}

const Home = (): JSX.Element => {
  return (
    <Layout title="Homepage" description="Synapse Dev Docs">
      <main>
        <br />
        <h1 style={{ fontWeight: '750', textAlign: 'center' }}>
          Get Started Below!
        </h1>
        <section className={styles.features}>
          <div className="container">
            <div className="row cards__container">
              <Card
                to="docs/solidity/intro"
                header={{
                  label: '🚀 Send your first message',
                }}
                body={{
                  label:
                    'Learn how to send a cross-chain message in under five minutes',
                }}
              />

              <Card
                to="docs/consensus"
                header={{
                  label: '🛠 Learn about Synapse',
                }}
                body={{
                  label: 'Discover how cross-chain messaging works.',
                }}
              />

              <Card
                to="docs/offchain"
                header={{
                  label: '😎 Run an Agent',
                }}
                body={{
                  label:
                    "Join Synapse's Testnet to help secure cross-chain messages. Learn how to run a notary, guard or executor in under 5 minutes.",
                }}
              />

              <Card
                to="https://docs.synapseprotocol.com/developers/rest-api"
                header={{
                  label: '💻 View Bridge API docs',
                }}
                body={{
                  label:
                    'Access bridge api docs to learn how to interact with the bridge.',
                }}
              />

              <Card
                to="/docs/sdk/"
                header={{
                  label: '🛠️ View Bridge SDK Docs',
                }}
                body={{
                  label:
                    'Learn how to interact with the synapse bridge using the bridge sdk and send your first bridge in 5 minutes!',
                }}
              />

              <Card
                to="https://docs.synapseprotocol.com/protocol/synapse-chain"
                header={{
                  label: '⛓️️ Develop on Synapse Chain',
                }}
                body={{
                  label: 'Learn how to deploy your dapp on the synapse chain.',
                }}
              />
            </div>
          </div>
        </section>
      </main>
    </Layout>
  )
}
export default Home
