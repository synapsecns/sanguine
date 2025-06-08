export const logger = console

// Decorator to log the execution time of a function
export const logExecutionTime =
  (explicitName?: string) =>
  (_target: any, propertyKey: string, descriptor: PropertyDescriptor) => {
    const originalMethod = descriptor.value

    descriptor.value = async function (...args: any[]) {
      const startTime = Date.now()
      try {
        return await originalMethod.apply(this, args)
      } finally {
        const elapsedTime = Date.now() - startTime
        const functionName =
          explicitName ?? `${this.constructor.name}.${propertyKey}`
        logger.info(`⌛ ${functionName}: ${elapsedTime}ms`)
      }
    }

    return descriptor
  }
