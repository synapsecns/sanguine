---
title: Widget
---

import SVGWidget from '@site/src/components/SVGWidget'

# Bridge Widget

The Synapse Widget lets you quickly and easily add bridging to their DeFi application, access specific tokens supported by the Synapse protocol, or even custom-build your own Synapse frontend.

<figure>
    <SVGWidget />
    <figcaption>Synapse Widget</figcaption>
</figure>

## Install

Requires `React`, and the `npm` or `yarn` package manager.

| Options |
|-|
| `npm install @synapsecns/widget`
| `yarn add @synapsecns/widget`

## Quick start

`import { Bridge } from @synapsecns/widget` into your app, and initialize it with your `web3Provider`.

```jsx
import { Bridge } from '@synapsecns/widget'

const MyApp = () => {
  const web3Provider = new ethers.BrowserProvider(window.ethereum)

  return <Bridge web3Provider={web3Provider} />
}
```

This will result in a fully operational Bridge integrating the routes and tokens supported by the Synapse protocol.

## Properties

While the widget works out-of-the-box without any setup beyond `web3Provider`, configuration is recommended for reliability and performance.

* `web3Provider` (required): Handles wallet connections.
* `customRpcs` (recommended):  JSON-RPC endpoints.
* `targetChainIds`: List of destination chain IDs. Defaults to *all*.
* `targetTokens`: List of tokens to display. These tokens are imported from the widget package. Defaults to *all*.
* `customTheme`: Custom theme for the widget. Defaults to light mode. see [Theme](#theme) section for details.
* `container`: Includes a solid-background container if `true`. Defaults to `false`.
* `protocolName`: Short name by which to identify the protocol. Defaults to `"Target"`.
* `hideConsoleErrors`: Hide SDK and Widget `console.error` messages. Defaults to `false`.

## web3Provider

While the demo landing page uses the `ethers` library, any similar provider can be used.

```jsx
// Ethers v5
const web3Provider = new ethers.providers.Web3Provider(window.ethereum, 'any')

// Ethers v6
const web3Provider = new ethers.BrowserProvider(window.ethereum)
```

## customRpcs

Set preferred RPC endpoints for each `chainId`. Defaults to Synapse fallback values for undefined chains.

```jsx
import { Bridge, CustomRpcs } from '@synapsecns/widget'

const customRpcs: CustomRpcs = {
  1: 'https://ethereum.my-custom-rpc.com',
  10: 'https://optimism.my-custom-rpc.com',
  42161: 'https://arbitrum.my-custom-rpc.com',
}

const MyApp = () => {
  const web3Provider = new ethers.BrowserProvider(window.ethereum)

  return <Bridge web3Provider={web3Provider} customRpcs={customRpcs} />
}
```

## targetChainIds & targetTokens

Shows only the chains and tokens your project supports for onboarding, while still allowing users to onboard from, or bridge back to, any preferred chain or token.

```jsx
import { Bridge, CustomRpcs, ETH, USDC, USDT } from '@synapsecns/widget'

const MyApp = () => {
  const web3Provider = new ethers.BrowserProvider(window.ethereum)

  return (
    <Bridge
      web3Provider={web3Provider}
      targetChainIds={[42161, 43114]}
      targetTokens={[ETH, USDC, USDT]}
    />
  )
}
```

:::tip Token names

Token names must match the definitions in `src/constants/bridgeable.ts`. Metis USDC, for example, is `"METISUSDC"`.

For chains, see `src/constants/chains.ts` as well.

:::

## Theme

The widget will automatically generate a color palette from `bgColor` in the `customTheme` object, which you can also use to override individual color variables.

:::tip Color modes

If your application has multiple color modes, such as light and dark, reload the widget with the appropriate theme colors when your application’s theme changes.

:::

```jsx
const customTheme = {
  // Generate from base color, 'dark', or 'light'
  bgColor: '#08153a',

  // Basic customization
  '--synapse-text': 'white',
  '--synapse-secondary': '#ffffffb3',
  '--synapse-root': '#16182e',
  '--synapse-surface': 'linear-gradient(90deg, #1e223de6, #262b47e6)',
  '--synapse-border': 'transparent',

  // Full customization (Uses 'basic' colors by default)
  '--synapse-focus': 'var(--synapse-secondary)',
  '--synapse-select-bg': 'var(--synapse-root)',
  '--synapse-select-text': 'var(--synapse-text)',
  '--synapse-select-border': 'var(--synapse-border)',
  '--synapse-button-bg': 'var(--synapse-surface)',
  '--synapse-button-text': 'var(--synapse-text)',
  '--synapse-button-border': 'var(--synapse-border)',

  // Transaction progress colors (set bgColor to auto-generate)
  '--synapse-progress': 'hsl(265deg 100% 65%)',
  '--synapse-progress-flash': 'hsl(215deg 100% 65%)',
  '--synapse-progress-success': 'hsl(120deg 100% 30%)',
  '--synapse-progress-error': 'hsl(15deg 100% 65%)',
}

<Bridge
  web3Provider={web3Provider}
  customTheme={customTheme}
/>
```

## Container

The widget is full-width with a transparent background by default, but can supply its own solid background container.

```jsx
<Bridge web3Provider={web3Provider} container={true} />
```

## useBridgeSelections hook

An active list of dropdown selections are available from the `useBridgeSelections` React hook.

```jsx
const {
  originChain, 
  originToken, 
  destinationChain, 
  destinationToken, 
} = useBridgeSelections()
```

### Structure

* `Chain: { id: number, name: string }`
* `Token: { symbol: string, address: string }`

## Example Apps

For reference, you can find three example apps in the repository’s `/examples` folder.

| `examples/` | Description
|-----------------|-
| `landing-page`  | Functional demo with basic customization
| `with-react`    | Simple React implementation
| `with-next`     | Simple Next.js implementation

## Support

If you have questions or need help implementing the widget, reach out to the team on [Discord](https://discord.gg/synapseprotocol).
