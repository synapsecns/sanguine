import React, { useState } from 'react'
import { useTranslations } from 'next-intl'
import { useAppDispatch } from '@/store/hooks'
import { updateBridgeParameters } from '@/slices/bridge/actions'
import Button from '@/components/ui/tailwind/Button'
import { toast } from 'react-hot-toast'
import { useAccount } from 'wagmi'
import PoweredByCx from '@/components/PoweredByCx'
export const ChatInputBox = () => {
  const [input, setInput] = useState('')
  const [isProcessing, setIsProcessing] = useState(false)
  const [isFocused, setIsFocused] = useState(false)
  const [lastProcessedInput, setLastProcessedInput] = useState('')
  const dispatch = useAppDispatch()
  const t = useTranslations('Bridge')
  const { isConnected } = useAccount()

  const handleSubmit = async () => {
    if (!input.trim()) return
    
    // Save current input for reference
    const currentInput = input.trim()
    
    // Don't process the same input twice in a row
    if (currentInput === lastProcessedInput) {
      toast.error('You already submitted this request. Try something different.')
      return
    }
    
    setIsProcessing(true)
    try {
      const result = await dispatch(updateBridgeParameters(currentInput))
      
      // Only clear the input if the request was successful
      // We determine success if the action returned a defined value
      if (result.payload) {
        setInput('') // Clear input field after processing
        setLastProcessedInput(currentInput)
      }
    } catch (error) {
      console.error('Error processing bridge request:', error)
      toast.error('An unexpected error occurred. Please try again.')
    } finally {
      setIsProcessing(false)
    }
  }

  const handleKeyDown = (e) => {
    if (e.key === 'Enter' && !isProcessing && input.trim() && isConnected) {
      handleSubmit()
    }
  }

  // Get button text based on conditions
  const getButtonText = () => {
    if (isProcessing) {
      return (
        <span className="flex items-center">
          <svg className="animate-spin -ml-1 mr-2 h-4 w-4 text-white" 
               style={{ animationDuration: '0.6s' }} // Make spinner faster
               xmlns="http://www.w3.org/2000/svg" 
               fill="none" 
               viewBox="0 0 24 24">
            <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
            <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          {t('Processing')}
        </span>
      )
    }
    
    if (!isConnected) {
      return t('Connect Wallet')
    }
    
    return t('Send')
  }

  // Determine if button should be disabled
  const isButtonDisabled = isProcessing || !input.trim() || !isConnected

  return (
    <div className="mt-4 p-3 transition-all duration-200">
      <div className="text-sm text-secondaryTextColor mb-2">
        Autofill with a simple command:
      </div>
      <div className="flex items-center gap-2">
        <div className={`flex items-center gap-4 p-2 rounded-md w-full
                        bg-white dark:bg-inherit
                        border border-zinc-200 dark:border-zinc-700
                        ${isFocused ? 'border-primary dark:border-primary' : ''}`}>
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={handleKeyDown}
            onFocus={() => setIsFocused(true)}
            onBlur={() => setIsFocused(false)}
            className="w-full bg-transparent border-none p-0 placeholder:text-zinc-500 placeholder:dark:text-zinc-400 text-black dark:text-white focus:outline-none"
            placeholder="e.g., Bridge 50 USDC from Base to Arb"
            disabled={isProcessing}
            autoComplete="off"
          />
        </div>
        <Button
          onClick={handleSubmit}
          disabled={isButtonDisabled}
          className={`whitespace-nowrap px-3 py-2 transition-all duration-200
                     ${isButtonDisabled 
                      ? 'opacity-50 bg-gray-600' 
                      : 'enabled:bg-gradient-to-r !from-fuchsia-500 !to-purple-500 dark:!to-purple-600 hover:opacity-80'}`}
        >
          {getButtonText()}
        </Button>
      </div>
      
      {/* Two-column layout for examples and branding */}
      <div className="mt-2 grid grid-cols-2">
        <div className="text-xs text-secondaryTextColor">
          Try: "100 ETH to Arb"<br />
          or "USDC from Op to Base"
        </div>
        <div className="flex justify-end">
          <PoweredByCx />
        </div>
      </div>
    </div>
  )
}

export default ChatInputBox