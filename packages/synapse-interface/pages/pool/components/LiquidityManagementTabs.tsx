import Tabs from '@tw/Tabs'
import TabItem from '@tw/TabItem'

const LiquidityManagementTabs = ({ cardNav, setCardNav }) => {
  return (
    <Tabs>
      <TabItem
        isActive={cardNav === 'addLiquidity'}
        onClick={() => {
          setCardNav('addLiquidity')
        }}
        className="rounded-tl-md"
      >
        Add Liquidity
      </TabItem>
      <TabItem
        isActive={cardNav === 'removeLiquidity'}
        onClick={() => {
          setCardNav('removeLiquidity')
        }}
        className="rounded-tr-md"
      >
        Remove Liquidity
      </TabItem>
    </Tabs>
  )
}
export default LiquidityManagementTabs
