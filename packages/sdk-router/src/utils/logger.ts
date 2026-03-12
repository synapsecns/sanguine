export const logger = console

// Check once at module load time (not on every call)
const ENABLE_TIMING_LOGS = process.env.LOG_EXECUTION_TIME === 'true'

// Decorator to log the execution time of a function
// Only active when LOG_EXECUTION_TIME=true
export const logExecutionTime =
  (explicitName?: string) =>
  (_target: any, propertyKey: string, descriptor: PropertyDescriptor) => {
    // When disabled, return original descriptor unchanged (no-op)
    if (!ENABLE_TIMING_LOGS) {
      return descriptor
    }

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
