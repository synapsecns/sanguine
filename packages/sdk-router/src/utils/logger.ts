export const logger = console

// Decorator to log the execution time of a function
export const logExecutionTime =
  () => (_target: any, propertyKey: string, descriptor: PropertyDescriptor) => {
    const originalMethod = descriptor.value

    descriptor.value = async function (...args: any[]) {
      const startTime = Date.now()
      const result = await originalMethod.apply(this, args)
      const elapsedTime = Date.now() - startTime

      const className = this.constructor.name
      const functionName = `${className}.${propertyKey}`

      logger.info(`⌛ ${functionName}: ${elapsedTime}ms`)
      return result
    }

    return descriptor
  }
