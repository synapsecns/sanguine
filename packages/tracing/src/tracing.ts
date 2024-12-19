import { NodeSDK } from '@opentelemetry/sdk-node';
import { ExpressInstrumentation } from '@opentelemetry/instrumentation-express';
import { Resource } from '@opentelemetry/resources';
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions';

export interface TracingConfig {
  serviceName: string;
  version: string;
  environment?: string;
}

export function initializeTracing(config: TracingConfig): NodeSDK {
  const sdk = new NodeSDK({
    resource: new Resource({
      [SemanticResourceAttributes.SERVICE_NAME]: config.serviceName,
      [SemanticResourceAttributes.SERVICE_VERSION]: config.version,
      environment: config.environment || 'development',
    }),
    instrumentations: [
      new ExpressInstrumentation(),
    ],
  });

  sdk.start();
  return sdk;
}
