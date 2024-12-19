import { trace, context, SpanStatusCode } from '@opentelemetry/api';
import { Request, Response, NextFunction } from 'express';

export interface TracingMiddlewareConfig {
  serviceName?: string;
}

export function tracingMiddleware(config: TracingMiddlewareConfig = {}): (req: Request, res: Response, next: NextFunction) => void {
  const tracer = trace.getTracer(config.serviceName || 'rest-api');

  return (req: Request, res: Response, next: NextFunction) => {
    const span = tracer.startSpan(`${req.method} ${req.path}`, {
      attributes: {
        'http.method': req.method,
        'http.url': req.url,
        'http.route': req.path,
      },
    });

    // Set the current span in context
    const ctx = trace.setSpan(context.active(), span);
    return context.with(ctx, () => {
      // Handle response finish
      res.on('finish', () => {
        span.setAttributes({
          'http.status_code': res.statusCode,
        });

        if (res.statusCode >= 400) {
          span.setStatus({
            code: SpanStatusCode.ERROR,
            message: `HTTP ${res.statusCode}`,
          });
        }

        span.end();
      });

      next();
    });
  };
}
