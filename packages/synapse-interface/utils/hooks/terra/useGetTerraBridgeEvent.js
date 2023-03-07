import { useLCDClient } from "@terra-money/wallet-provider"

import { SYNAPSE_BRIDGE_ADDRESSES } from "@constants/bridge"
import { ChainId } from "@constants/networks"

import { formatTerraTx, formatTerraWithdrawTx } from "@utils/terra/formatTerraTx"


export function useGetTerraBridgeEvent() {
  // const otherStuff = useTerraWallet()
  const terra = useLCDClient()
  // console.log({terra, otherStuff})
  async function getTerraBridgeEvent({ kekTxSig }) {
    const terraResult = await terra.wasm.contractQuery(
      SYNAPSE_BRIDGE_ADDRESSES[ChainId.TERRA],
      {
        get_kappa: {
          kappa: kekTxSig.slice(2, kekTxSig.length)
        }
      }
    )

    if (terraResult?.withdraw) {
      return formatTerraWithdrawTx(terraResult?.withdraw)
    } else {
      return formatTerraTx(terraResult?.deposit ?? {})
    }

  }
  return getTerraBridgeEvent
}

//"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f" //
// "timestamp": "1644721887881054858",
    // "block_height": 6460544,
    // "message_index": "17",
    // "denom": "uusd",
    // "amount": "4000000",
    // "fee": "1000000",
    // "to_address": "terra1ua547gy82azz9ptl8rpg0yz5dwt9na62hmglan",
    // "kappa": "aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"

/*
{"GetKappa":{"kappa":"0f78ed9cd78e7307ceb6711c4901c64ed0b9cd902902dbb50e3add24c18223d4"}}
{"GetKappa":"0f78ed9cd78e7307ceb6711c4901c64ed0b9cd902902dbb50e3add24c18223d4"}
{"get_kappa":{"kappa":"0f78ed9cd78e7307ceb6711c4901c64ed0b9cd902902dbb50e3add24c18223d4"}}
{"get_kappa":"0f78ed9cd78e7307ceb6711c4901c64ed0b9cd902902dbb50e3add24c18223d4"}

{"GetKappa":{"kappa":"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"}}
{"GetKappa":"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"}
{"get_kappa":{"kappa":"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"}}
{"get_kappa":"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"}


{
    "get_kappa": {
        "kappa": "54a526d3a4c6dd19605aad4dc82d69bff4b272c4d58ee7e90e3a493d8e563a37"
    }
}

{
    "get_kappa": {
        "kappa": "27a45c50a0f7802adcd19957982776d3e9fb7fac5c09a76d4cc88ef8c838d5c3"
    }
}
0x27a45c50a0f7802adcd19957982776d3e9fb7fac5c09a76d4cc88ef8c838d5c3
{"get_deposits_by_address": {"recipient":"terra1ua547gy82azz9ptl8rpg0yz5dwt9na62hmglan"}}
{"get_deposits_by_recipient": {"recipient":"terra1ua547gy82azz9ptl8rpg0yz5dwt9na62hmglan"}}


*/