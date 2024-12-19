const pyroscope = require('@pyroscope/nodejs');

export interface PyroscopeConfig {
  applicationName: string;
  serverAddress: string;
  tags?: Record<string, string>;
}

export function startPyroscope(config: PyroscopeConfig): any {
  if (!process.env.PYROSCOPE_ENDPOINT) {
    return null;
  }

  return pyroscope.init({
    serverAddress: process.env.PYROSCOPE_ENDPOINT,
    applicationName: config.applicationName,
    tags: {
      ...config.tags,
      hostname: process.env.HOSTNAME,
    },
    profilers: {
      cpu: true,
      heap: true,
    },
  });
}
