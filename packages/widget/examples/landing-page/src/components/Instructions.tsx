import { Install, Developer, Support } from './Icons'
import { PackageInstall } from './PackageInstall'
import { GetStarted } from './GetStarted'
import { Appearance } from './Appearance'
import { RecommendedParameters } from './RecommendedParameters'
import { TokenAndChainCustomization } from './TokenAndChainCustomization'

export function Instructions() {
  return (
    <article>
      <h2>
        <Install />
        Install
      </h2>
      <p>Install the Synapse Widget in your Next.js or React project</p>
      <PackageInstall />
      <h2>
        <Developer />
        Setup
      </h2>
      <GetStarted />
      <h3>Enhanced and Reliable Performance</h3>
      <RecommendedParameters />
      <h3>Token and Chain Customization</h3>
      <TokenAndChainCustomization />

      <h3>Appearance</h3>
      <Appearance />
      <h2>
        <Support />
        Support
      </h2>
      <p>
        For help and feedback, reach out to our Support team in the{' '}
        <a href="#" target="_blank" rel="noreferrer">
          Synapse Discord channel.
        </a>
      </p>
      <p>
        npm package is located{' '}
        <a
          href="https://www.npmjs.com/package/@synapsecns/widget"
          target="_blank"
        >
          here
        </a>
      </p>
    </article>
  )
}
