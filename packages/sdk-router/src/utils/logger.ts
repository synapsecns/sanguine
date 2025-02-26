export const logger = console

// Decorator to log the execution time of a function
export const logExecutionTime =
  (functionName: string) =>
  (_target: any, _propertyKey: string, descriptor: PropertyDescriptor) => {
    const originalMethod = descriptor.value

    descriptor.value = async function (...args: any[]) {
      const startTime = Date.now()
      const result = await originalMethod.apply(this, args)
      const elapsedTime = Date.now() - startTime
      logger.info({ args }, `${functionName} execution time: ${elapsedTime}ms`)
      return result
    }

    return descriptor
  }
