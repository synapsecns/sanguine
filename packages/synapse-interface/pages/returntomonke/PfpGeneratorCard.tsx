import { useTranslations } from 'next-intl'

import Card from '@tw/Card'
import ImageUploader from './ImageUploader'

export default function PfpGeneratorCard() {
  const t = useTranslations('ReturnToMonke')

  return (
    <Card
      title={t('Choose Synapse')}
      divider={false}
      className="rounded-xl min-w-[380px]"
      titleClassName="text-center text-white text-opacity-50 font-normal py-2"
    >
      <ImageUploader />
    </Card>
  )
}
