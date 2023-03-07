import {
  ANYSWAP_ROUTER_ADDRESSES,
  ANYNATIVE_ADDRESSES,
} from '@constants/anyswap'
import { ChainId, CHAIN_PARAMS } from '@constants/networks'
import { WAVAX, WBNB, WETH, WMATIC } from '@constants/tokens/basic'
import { useRevokeToken } from '@hooks/actions/useRevokeToken'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useCheckAnyToken } from '@hooks/tokens/useCheckAnyToken'
import Card from '@tw/Card'
import Grid from '@tw/Grid'
import Button from '@tw/Button'

export default function AnyswapWarningCard() {
  const wethParams = {
    targetChainId: ChainId.ETH,
    token: WETH,
    dappAddr: ANYSWAP_ROUTER_ADDRESSES[ChainId.ETH],
    type: 'router',
  }
  const wbnbParams = {
    targetChainId: ChainId.BSC,
    token: WBNB,
    dappAddr: ANYSWAP_ROUTER_ADDRESSES[ChainId.BSC],
    type: 'router',
  }
  const wmaticParams = {
    targetChainId: ChainId.POLYGON,
    token: WMATIC,
    dappAddr: ANYSWAP_ROUTER_ADDRESSES[ChainId.POLYGON],
    type: 'router',
  }
  const wavaxParams = {
    targetChainId: ChainId.AVALANCHE,
    token: WAVAX,
    dappAddr: ANYSWAP_ROUTER_ADDRESSES[ChainId.AVALANCHE],
    type: 'router',
  }

  const wethAnyParams = {
    targetChainId: ChainId.ETH,
    token: WETH,
    dappAddr: ANYNATIVE_ADDRESSES[ChainId.ETH],
    type: 'native',
  }
  const wbnbAnyParams = {
    targetChainId: ChainId.BSC,
    token: WBNB,
    dappAddr: ANYNATIVE_ADDRESSES[ChainId.BSC],
    type: 'native',
  }
  const wmaticAnyParams = {
    targetChainId: ChainId.POLYGON,
    token: WMATIC,
    dappAddr: ANYNATIVE_ADDRESSES[ChainId.POLYGON],
    type: 'native',
  }
  const wavaxAnyParams = {
    targetChainId: ChainId.AVALANCHE,
    token: WAVAX,
    dappAddr: ANYNATIVE_ADDRESSES[ChainId.AVALANCHE],
    type: 'native',
  }

  const wethStatus = useCheckAnyToken(wethParams)
  const wbnbStatus = useCheckAnyToken(wbnbParams)
  const wmaticStatus = useCheckAnyToken(wmaticParams)
  const wavaxStatus = useCheckAnyToken(wavaxParams)

  const wethAnyStatus = useCheckAnyToken(wethAnyParams)
  const wbnbAnyStatus = useCheckAnyToken(wbnbAnyParams)
  const wmaticAnyStatus = useCheckAnyToken(wmaticAnyParams)
  const wavaxAnyStatus = useCheckAnyToken(wavaxAnyParams)

  const tokenArr = [
    [wethParams, wethStatus],
    [wethAnyParams, wethAnyStatus],
    [wbnbParams, wbnbStatus],
    [wbnbAnyParams, wbnbAnyStatus],
    [wmaticParams, wmaticStatus],
    [wmaticAnyParams, wmaticAnyStatus],
    [wavaxParams, wavaxStatus],
    [wavaxAnyParams, wavaxAnyStatus],
  ]

  if (wethStatus || wbnbStatus || wmaticStatus || wavaxStatus) {
    return (
      <Grid
        cols={{ xs: 1 }}
        gap={6}
        className="justify-center px-2 py-16 sm:px-6 md:px-8"
      >
        <div className="pb-3 place-self-center">
          <Card
            title="Your Account is Affected by AnySwap Hack"
            className="shadow-red-2xl"
            divider={false}
          >
            <div className="pb-2 text-slate-400">
              In order to protect your funds, you should revoke approval.
            </div>
            <Grid
              cols={{ xs: 1, sm: 2 }}
              gap={4}
              className="content-center align-center"
            >
              {tokenArr
                .filter(([obj, status]) => status)
                .map(([obj, status]) => (
                  <div>
                    <RevokeRow {...obj} />
                  </div>
                ))}
            </Grid>
          </Card>
        </div>
      </Grid>
    )
  } else {
    ;<></>
  }
}

function RevokeRow({ token, targetChainId, dappAddr, type }:{ token: any, targetChainId: any, dappAddr: any, type: any }) {
  const { chainId } = useActiveWeb3React()
  const revokeToken = useRevokeToken(token)
  const isCurrentChain = targetChainId == chainId

  let btnContent
  if (isCurrentChain) {
    if (type == 'native') {
      btnContent = `Revoke AnySwap approval for any${token.name}`
    } else {
      btnContent = `Revoke AnySwap Router approval for ${token.name}`
    }
  } else {
    btnContent = `Switch to ${CHAIN_PARAMS[targetChainId].chainName} to revoke`
  }

  return (
    <Grid gap={2} cols={{ xs: 1 }} className="align-middle">
      <div className="pb-2">
        <Button
          disabled={!isCurrentChain}
          className="rounded-lg"
          onClick={() => {
            revokeToken({ addrToRevoke: dappAddr })
          }}
        >
          {btnContent}
        </Button>
      </div>
    </Grid>
  )
}
