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
    if (e.key === 'Enter' && !isProcessing && input.trim()) {
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
          Processing...
        </span>
      )
    }
    
    return "Send"
  }

  // Determine if button should be disabled
  const isButtonDisabled = isProcessing || !input.trim()

  return (
    <div className="transition-all duration-200">
      <div className="flex items-center">
        <div className={`flex items-center gap-4 p-2 pr-1 rounded-md w-full
                        bg-white dark:bg-inherit
                        border ${isFocused 
                          ? 'border-fuchsia-500 dark:border-fuchsia-400' 
                          : 'border-zinc-200 dark:border-zinc-700'}`}>
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyDown={handleKeyDown}
            onFocus={() => setIsFocused(true)}
            onBlur={() => setIsFocused(false)}
            className="w-full bg-transparent border-none p-0 placeholder:text-zinc-500 placeholder:dark:text-zinc-400 text-black dark:text-white focus:outline-none focus:ring-0"
            placeholder="Autofill with AI"
            disabled={isProcessing}
            autoComplete="off"
          />
          <button
            onClick={handleSubmit}
            disabled={isButtonDisabled}
            className={`flex items-center justify-center rounded-full w-8 h-8 ml-1 transition-all duration-200
                      ${isButtonDisabled 
                        ? 'opacity-50 text-gray-400 cursor-not-allowed' 
                        : 'text-fuchsia-500 hover:bg-fuchsia-100 dark:hover:bg-fuchsia-900/30'}`}
          >
            {isProcessing ? (
              <svg className="animate-spin h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
            ) : (
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M10.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L12.586 11H5a1 1 0 110-2h7.586l-2.293-2.293a1 1 0 010-1.414z" clipRule="evenodd" />
              </svg>
            )}
          </button>
        </div>
      </div>
      
      {/* Two-column layout for examples and branding */}
      <div className="mt-2 grid grid-cols-3">
        <div className="text-xs text-secondaryTextColor col-span-2">
          Try: "100 ETH to Arb" or "USDC from Op to Base"
        </div>
        <div className="flex justify-end">
          <PoweredByCx />
        </div>
      </div>
    </div>
  )
}

export default ChatInputBox