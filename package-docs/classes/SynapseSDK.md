[sanguine](../README.md) / [Exports](../modules.md) / SynapseSDK

# Class: SynapseSDK

## Table of contents

### Constructors

- [constructor](SynapseSDK.md#constructor)

### Properties

- [synapseRouters](SynapseSDK.md#synapserouters)

### Methods

- [bridge](SynapseSDK.md#bridge)
- [bridgeQuote](SynapseSDK.md#bridgequote)

## Constructors

### constructor

• **new SynapseSDK**(`chainIds`, `providers`)

#### Parameters

| Name | Type |
| :------ | :------ |
| `chainIds` | `number`[] |
| `providers` | `Provider`[] |

#### Defined in

[sdk.ts:13](https://github.com/synapsecns/sanguine/blob/561a8e25/packages/sdk-router/src/sdk.ts#L13)

## Properties

### synapseRouters

• **synapseRouters**: `any`

#### Defined in

[sdk.ts:11](https://github.com/synapsecns/sanguine/blob/561a8e25/packages/sdk-router/src/sdk.ts#L11)

## Methods

### bridge

▸ **bridge**(`to`, `originChainId`, `destChainId`, `token`, `amount`, `originQuery`, `destQuery`): `Promise`<`any`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `to` | `string` |
| `originChainId` | `number` |
| `destChainId` | `number` |
| `token` | `string` |
| `amount` | `BigintIsh` |
| `originQuery` | `Object` |
| `originQuery.deadline` | `BigintIsh` |
| `originQuery.minAmountOut` | `BigintIsh` |
| `originQuery.rawParams` | `BytesLike` |
| `originQuery.swapAdapter` | `string` |
| `originQuery.tokenOut` | `string` |
| `destQuery` | `Object` |
| `destQuery.deadline` | `BigintIsh` |
| `destQuery.minAmountOut` | `BigintIsh` |
| `destQuery.rawParams` | `BytesLike` |
| `destQuery.swapAdapter` | `string` |
| `destQuery.tokenOut` | `string` |

#### Returns

`Promise`<`any`\>

#### Defined in

[sdk.ts:99](https://github.com/synapsecns/sanguine/blob/561a8e25/packages/sdk-router/src/sdk.ts#L99)

___

### bridgeQuote

▸ **bridgeQuote**(`originChainId`, `destChainId`, `tokenIn`, `tokenOut`, `amountIn`): `Promise`<`any`\>

#### Parameters

| Name | Type |
| :------ | :------ |
| `originChainId` | `number` |
| `destChainId` | `number` |
| `tokenIn` | `string` |
| `tokenOut` | `string` |
| `amountIn` | `BigintIsh` |

#### Returns

`Promise`<`any`\>

#### Defined in

[sdk.ts:27](https://github.com/synapsecns/sanguine/blob/561a8e25/packages/sdk-router/src/sdk.ts#L27)
