enum SynapseErrorPrefix {
  SDK = 'SynapseSDK',
  WIDGET = 'Synapse Widget',
}

export const suppressSynapseConsoleErrors = () => {
  console.log('Supressing Synapse console errors.')

  // Store the original console.error function
  const originalConsoleError = console.error

  // Redefine console.error to filter out specific messages
  console.error = (...args) => {
    const message = args.join(' ')

    // Suppress Synapse console.error messages
    if (
      message.includes(SynapseErrorPrefix.SDK) ||
      message.includes(SynapseErrorPrefix.WIDGET)
    ) {
      return
    }

    // Call the original console.error for other messages
    originalConsoleError(...args)
  }
}
