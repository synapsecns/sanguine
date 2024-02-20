import { useEffect } from 'react'
import { GlobalEventEmitter } from '@/utils/globalEventEmitter'

const modalContainerStyle = `position: fixed; top: 0; left: 0; right: 0; bottom: 0; display: flex; justify-content: center; align-items: center; background-color: rgba(0, 0, 0, 0.7); z-index: 1000`

export const useRiskEvent = () => {
  useEffect(() => {
    const handleRiskDetected = (event) => {
      const newBody = document.createElement('body')
      const modalContainer = document.createElement('div')

      modalContainer.setAttribute('style', modalContainerStyle)

      modalContainer.innerHTML = `
        <div style="background-color: #fff; padding: 20px; border-radius: 5px; box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); max-width: 500px; width: 90%; margin: auto;">
          <p>${event.detail.message}</p>
        </div>
      `

      newBody.appendChild(modalContainer)
      document.documentElement.replaceChild(newBody, document.body)
    }

    GlobalEventEmitter.addEventListener('riskDetected', handleRiskDetected)

    return () => {
      GlobalEventEmitter.removeEventListener('riskDetected', handleRiskDetected)
    }
  }, [])
}
