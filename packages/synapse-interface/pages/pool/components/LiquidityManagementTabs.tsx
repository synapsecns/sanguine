import Tabs from '@tw/Tabs'
import TabItem from '@tw/TabItem'

export default function LiquidityManagementTabs({ cardNav, setCardNav }) {
  return (
    <Tabs>
      <TabItem
        isActive={cardNav === 'addLiquidity'}
        onClick={() => {
          setCardNav('addLiquidity')
        }}
      >
        Add Liquidity
      </TabItem>
      <TabItem
        isActive={cardNav === 'removeLiquidity'}
        onClick={() => {
          setCardNav('removeLiquidity')
        }}
      >
        Remove Liquidity
      </TabItem>
    </Tabs>
  )
}
