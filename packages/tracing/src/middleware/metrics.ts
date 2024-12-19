import { metrics } from '@opentelemetry/api';
import { Request, Response, NextFunction } from 'express';

const HTTP_REQUEST_DURATION = 'http.server.duration';
const HTTP_REQUEST_ACTIVE = 'http.server.active_requests';

export interface MetricsMiddlewareConfig {
  serviceName: string;
}

export function metricsMiddleware(config: MetricsMiddlewareConfig): (req: Request, res: Response, next: NextFunction) => void {
  const meter = metrics.getMeter(config.serviceName);

  const requestDuration = meter.createHistogram(HTTP_REQUEST_DURATION, {
    description: 'Duration of HTTP requests',
    unit: 'ms',
  });

  const activeRequests = meter.createUpDownCounter(HTTP_REQUEST_ACTIVE, {
    description: 'Number of concurrent HTTP requests',
  });

  return (req: Request, res: Response, next: NextFunction) => {
    const startTime = Date.now();
    const attributes = {
      method: req.method,
      route: req.path,
    };

    activeRequests.add(1, attributes);

    res.on('finish', () => {
      const duration = Date.now() - startTime;
      requestDuration.record(duration, {
        ...attributes,
        status_code: res.statusCode.toString(),
      });
      activeRequests.add(-1, attributes);
    });

    next();
  };
}
