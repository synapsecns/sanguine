import { useState, useRef, useEffect } from 'react'
import { useTranslations } from 'next-intl'

import Button from '@tw/Button'

export default function ImageUploader() {
  const [uploadedImage, setUploadedImage] = useState(null)
  const [processedImage, setProcessedImage] = useState(null)

  const t = useTranslations('ReturnToMonke')

  const fileInput = useRef(null)
  const imgRef = useRef(null)

  useEffect(() => {
    if (uploadedImage) {
      applyImageOverlay()
    }
  }, [uploadedImage])

  function handleFile(file) {
    setUploadedImage(file)
  }

  function handleDragOver(event) {
    event.preventDefault()
  }

  function handleOnDrop(event) {
    event.preventDefault()
    event.stopPropagation()
    handleFile(event.dataTransfer.files[0])
    console.log('File drop was a success')
  }

  const applyImageOverlay = () => {
    const canvas = imgRef.current
    const ctx = canvas.getContext('2d')

    // Clear the canvas
    ctx.clearRect(0, 0, canvas.width, canvas.height)

    // Load the uploaded image
    const img = new Image()
    img.src = URL.createObjectURL(uploadedImage)
    img.onload = () => {
      if (img.width >= img.height) {
        canvas.width = img.width
        canvas.height = img.width
      } else {
        canvas.width = img.height
        canvas.height = img.height
      }

      // Calculate the start positions to center the image
      const startX = (canvas.width - img.width) / 2
      const startY = (canvas.height - img.height) / 2

      // Draw the uploaded image centered
      ctx.drawImage(img, startX, startY)

      // Load the overlay image
      const overlay = new Image()
      overlay.src = './synpfpborder.png'
      overlay.onload = () => {
        // Draw the overlay image on top of the uploaded image
        ctx.drawImage(overlay, 0, 0, canvas.width, canvas.height)

        // Convert the canvas to a Blob and update the state
        canvas.toBlob((blob) => {
          setProcessedImage(blob)
        }, 'image/jpeg')
      }
    }
  }

  return (
    <div>
      <div
        className={`
          flex ${
            processedImage ? '' : 'h-[200px]'
          }  w-full  rounded-xl cursor-pointer
          border-dashed border-4 border-white border-opacity-50 mb-2
          text-white text-opacity-50 hover:text-opacity-100
          hover:border-white
          transform transition-all duration-200
        `}
        style={{
          backgroundRepeat: 'no-repeat',
          backgroundSize: 100,
          backgroundPosition: 'center',
        }}
        onDragOver={handleDragOver}
        onDrop={handleOnDrop}
        onClick={() => fileInput.current.click()}
      >
        <p className="p-3 m-auto text-lg text-center center">
          {!processedImage && (
            <>
              {t('Click to Upload or')}
              <br />
              {t('Drag and drop image here')}...
            </>
          )}
          {processedImage && <>{t('Select different image?')}</>}
        </p>
        <input
          type="file"
          accept="image/*"
          ref={fileInput}
          hidden
          onChange={(e) => handleFile(e.target.files[0])}
        />
      </div>
      {processedImage && (
        <img src={URL.createObjectURL(processedImage)} alt="preview" />
      )}
      <canvas ref={imgRef} style={{ display: 'none' }} />
      {uploadedImage && (
        <Button
          className={`
            w-full rounded-xl my-2 px-4 py-3 tracking-wide
            text-white text-opacity-100
            hover:opacity-80 disabled:opacity-100 disabled:text-[#88818C]
            disabled:from-bgLight disabled:to-bgLight
            bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
          `}
          onClick={() => {
            downloadBase64File(
              URL.createObjectURL(processedImage),
              `synape_${uploadedImage.name?.split('.')[0] ?? 'random'}.jpg`
            )
          }}
        >
          {t('Download')}
        </Button>
      )}
    </div>
  )
}

function downloadBase64File(dataStr, fileName) {
  const downloadLink = document.createElement('a')
  downloadLink.href = dataStr
  downloadLink.download = fileName
  downloadLink.click()
  console.log('evolving into a synape')
}
