import _ from 'lodash'

import { useLCDClient } from "@terra-money/wallet-provider"

import { SYNAPSE_BRIDGE_ADDRESSES } from "@constants/bridge"
import { ChainId } from "@constants/networks"

import { formatTerraTx } from '@utils/terra/formatTerraTx'


export function useGetTerraBridgeEvents() {
  // const otherStuff = useTerraWallet()
  const lcd = useLCDClient()
  // console.log({lcd, otherStuff})
  async function getTerraBridgeEvents({ account, terraAddress }) {
    const terraDeposits = await lcd.wasm.contractQuery(
      SYNAPSE_BRIDGE_ADDRESSES[ChainId.TERRA],
      {
        get_deposits_by_sender: {
          sender: terraAddress
        }
      }
    )
    const terraDepositsResult = terraDeposits?.deposits ?? []
    // console.log({terraDepositsResult})
    const terraDepositTxns = (
      _.zip(
        (await Promise.all(
          terraDepositsResult.map(terraEvent => lcd.tx.txInfosByHeight(terraEvent.block_height))
        )),
        terraDepositsResult,
      ).map(([txInfos, terraEvent]) => {
        // console.log({ txInfos, terraEvent })
        const matchedTxInfo = _.find(txInfos, (txInfo) => {
          const message = txInfo.tx.body.messages[0]
          const isMatched =
            (message.contract == terraEvent.contract_address)
            && (message.sender == terraEvent.sender)
          return isMatched
        })

        return ({ ...terraEvent, transactionHash: matchedTxInfo?.txhash })
      })
    ).map(formatTerraTx)
    return terraDepositTxns
  }
  return getTerraBridgeEvents
}


        //"0xec254f185ce3e73a20266f15af87d771941d15269c3b733daedf6c03da789eae"
    // const [terraRecipientDepositsResult, terraSenderDepositsResult] = await Promise.all( //terraWithdrawsResult
    //   terra.wasm.contractQuery(
    //     SYNAPSE_BRIDGE_ADDRESSES[ChainId.TERRA],
    //     {
    //       get_deposits_by_recipient: {
    //         recipient: account
    //       }
    //     }
    //   ),
    //   terra.wasm.contractQuery(
    //     SYNAPSE_BRIDGE_ADDRESSES[ChainId.TERRA],
    //     {
    //       get_deposits_by_sender: {
    //         sender: terraAddress
    //       }
    //     }
    //   ),
    //   // terra.wasm.contractQuery(
    //   //   SYNAPSE_BRIDGE_ADDRESSES[ChainId.TERRA],
    //   //   {
    //   //     get_withdraws_by_recipient: {
    //   //       recipient: terraAddress
    //   //     }
    //   //   }
    //   // )
    // )
    // console.log([terraRecipientDepositsResult, terraSenderDepositsResult,]) // terraWithdrawsResult
    // const terraDepositsRecipient = terraRecipientDepositsResult?.deposits ?? []
    // const terraDepositsSender = terraSenderDepositsResult?.deposits ?? []
    // // const terraWithdraws = terraWithdrawsResult?.withdraws ?? []

    // const terraDepositRecipientTxns = terraDepositsRecipient.map(formatTerraTx)
    // const terraDepositSenderTxns = terraDepositsSender.map(formatTerraTx)
    // // const terraWithdrawTxns = terraWithdraws.map(formatTerraTx)
    // console.log({ terraDepositRecipientTxns, terraDepositSenderTxns, }) //terraWithdrawTxns
    // // return _.uniqBy([...terraDepositRecipientTxns, ...terraDepositSenderTxns, ...terraWithdrawTxns], (i) => i.sender) ?? []
    // return []
/*
{"GetKappa":{"kappa":"0f78ed9cd78e7307ceb6711c4901c64ed0b9cd902902dbb50e3add24c18223d4"}}
{"GetKappa":"0f78ed9cd78e7307ceb6711c4901c64ed0b9cd902902dbb50e3add24c18223d4"}
{"get_kappa":{"kappa":"0f78ed9cd78e7307ceb6711c4901c64ed0b9cd902902dbb50e3add24c18223d4"}}
{"get_kappa":"0f78ed9cd78e7307ceb6711c4901c64ed0b9cd902902dbb50e3add24c18223d4"}

{"GetKappa":{"kappa":"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"}}
{"GetKappa":"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"}
{"get_kappa":{"kappa":"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"}}
{"get_kappa":"aedbe41608a7649fc6f8b2c0111e1eecd1bb18c00ecde13584e05c946dbad25f"}



{"get_deposits_by_address": {"recipient":"terra1ua547gy82azz9ptl8rpg0yz5dwt9na62hmglan"}}
{"get_deposits_by_recipient": {"recipient":"terra1ua547gy82azz9ptl8rpg0yz5dwt9na62hmglan"}}
{"get_withdraws_by_address":{"recipient":"terra15q8axd0dvepk3cs2va8vvchsus9jhg35q0qspd"}}
{"get_deposits_by_sender": {"sender":"terra1dgj4554t5ewgw8kclh3qkeqwdav3fqqp5zjfg0"} }

{"get_deposits_by_sender": {"sender":"terra1ua547gy82azz9ptl8rpg0yz5dwt9na62hmglan"}}
*/