import Card from '@tw/Card'

import ImageUploader from './ImageUploader'

export default function PfpGeneratorCard() {
  return (
    <Card
      title={
        <>
          Choose Synapse
        </>
      }
      divider={false}
      className="rounded-xl min-w-[380px]"
      titleClassName="text-center text-white text-opacity-50 font-normal py-2"
    >
      <ImageUploader />
    </Card>
  )
}
