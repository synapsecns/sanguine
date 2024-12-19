import pyroscope from '@pyroscope/nodejs';

export interface PyroscopeConfig {
  applicationName: string;
  serverAddress?: string;
  tags?: Record<string, string>;
}

export function startPyroscope(config: PyroscopeConfig): void {
  if (!process.env.PYROSCOPE_ENDPOINT) {
    return;
  }

  pyroscope.init({
    serverAddress: process.env.PYROSCOPE_ENDPOINT,
    appName: config.applicationName,
    tags: {
      ...config.tags,
      hostname: process.env.HOSTNAME || 'unknown',
    },
    // @ts-expect-error: pyroscope types are incomplete
    profilers: {
      cpu: true,
      heap: true,
    },
  });
}
