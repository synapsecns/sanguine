import {
  Bridge,
  USDC,
  USDT,
  DAI,
  ETH,
  CustomRpcs,
} from '@abtestingalpha/widget'
import { useEthereumWallet } from './hooks/useEthereumWallet'
import { ChangeEventHandler, useEffect, useState } from 'react'
import Header from './Header'
import Footer from './Footer'
// import { Install, Developer, Support } from './icons'
import Instructions from './Instructions'

const targetTokens = [USDC, USDT, DAI, ETH]

const customRpcs: CustomRpcs = {
  1: 'https://eth.llamarpc.com',
  42161: 'https://arbitrum.llamarpc.com',
}

function App() {
  const [customTheme, setCustomTheme] = useState({})
  const [container, setContainer] = useState(true)

  const inputChangeHandler = (
    e: React.ChangeEvent<HTMLSelectElement>
  ): void => {
    const target = e.target as HTMLSelectElement

    switch (target.value) {
      case 'dark':
      case 'light':
        setCustomTheme({ bgColor: target.value })
        break
      case 'dfk':
        setCustomTheme({
          bgColor: 'light',
          '--synapse-text': 'hsl(12deg 85% 13%)',
          '--synapse-secondary': 'hsl(12deg 85% 20%)',
          '--synapse-select-bg': 'hsl(35deg 100% 87%)',
          '--synapse-surface': 'hsl(32deg 69% 78%)',
          '--synapse-root': 'hsl(35deg 100% 87%)',
          '--synapse-border': 'hsl(29deg 53% 68%)',
          '--synapse-focus': 'hsl(12deg 85% 15%)',
          '--synapse-accent': 'hsl(12deg 85% 15%)',
        })
        break
      case 'gmx':
        setCustomTheme({
          '--synapse-text': 'white',
          '--synapse-secondary': '#ffffffb3',
          '--synapse-root': '#16182e',
          '--synapse-surface': 'linear-gradient(90deg, #1e223de6, #262b47e6)',
          '--synapse-border': 'transparent',
          '--synapse-select-bg': 'hsl(231.5deg 32% 19.5%',
          '--synapse-select-border': 'hsl(233deg 34% 34%)',
          '--synapse-button-bg': '#2d42fc',
        })
        break
      case 'hercules':
        setCustomTheme({
          bgColor: 'dark',
          '--synapse-button-bg':
            'linear-gradient(90deg, hsl(43deg 79% 74%), hsl(21deg 76% 60%))',
          '--synapse-button-text': 'black',
          '--synapse-focus': 'hsl(32deg 77.5% 67%)',
        })
        break
      default:
        setCustomTheme({ bgColor: 'dark' })
    }
  }

  function createCustomTheme() {
    const colorPicker = document.getElementById(
      'color-picker'
    ) as HTMLInputElement | null

    setCustomTheme({ bgColor: colorPicker?.value })
  }

  const { web3Provider, connectedAddress, connectedNetwork } =
    useEthereumWallet()

  const toggleContainer = (e: React.ChangeEvent<HTMLInputElement>) =>
    setContainer(e.target.checked)

  const bridgeContainerDisplayProperty = container ? 'grid' : 'block'

  return (
    <>
      <Header />

      <main>
        <header>
          <h1>Install the Synapse Bridge</h1>
          <pre>npm synapse-widget</pre>
          <p>
            Easily onboard new users by adding a custom instance of the Synapse
            Bridge to your React project.
          </p>
          <div id="example-container">
            <div
              id="bridge-container"
              style={{ display: bridgeContainerDisplayProperty }}
            >
              <Bridge
                web3Provider={web3Provider}
                customRpcs={customRpcs}
                targetTokens={targetTokens}
                customTheme={customTheme}
                container={container}
                targetChainIds={[42161]}
              />
            </div>
            <div style={{ display: 'flex', justifyContent: 'space-between' }}>
              <div style={{ display: 'flex', gap: '1rem' }}>
                <input
                  id="color-picker"
                  type="color"
                  onInput={createCustomTheme}
                />
                <div className="flex items-center gap-1">
                  <label>Container</label>{' '}
                  <input
                    type="checkbox"
                    checked={container}
                    onChange={toggleContainer}
                  />
                </div>
                <select onChange={inputChangeHandler}>
                  <option>Select preset</option>
                  <option value="dark">Synapse Dark</option>
                  <option value="light">Synapse Light</option>
                  <option value="dfk">DeFi Kingdoms</option>
                  <option value="gmx">GMX</option>
                  <option value="hercules">Hercules</option>
                </select>
              </div>
              <span className="desktop-only">Drag to resize</span>
            </div>
          </div>
        </header>
        <Instructions />
      </main>

      <Footer />
    </>
  )
}

export default App
