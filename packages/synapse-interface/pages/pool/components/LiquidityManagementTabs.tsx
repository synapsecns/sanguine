import { useTranslations } from 'next-intl'

import Tabs from '@tw/Tabs'
import TabItem from '@tw/TabItem'

const LiquidityManagementTabs = ({ cardNav, setCardNav }) => {
  const t = useTranslations('Pools')

  return (
    <Tabs>
      <TabItem
        isActive={cardNav === 'addLiquidity'}
        onClick={() => {
          setCardNav('addLiquidity')
        }}
        className="rounded-tl-sm"
      >
        {t('Add Liquidity')}
      </TabItem>
      <TabItem
        isActive={cardNav === 'removeLiquidity'}
        onClick={() => {
          setCardNav('removeLiquidity')
        }}
        className="rounded-tr-sm"
      >
        {t('Remove Liquidity')}
      </TabItem>
    </Tabs>
  )
}
export default LiquidityManagementTabs
