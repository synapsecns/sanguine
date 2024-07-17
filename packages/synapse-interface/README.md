This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app)

## Getting Started

First, run the development server:

```bash
yarn dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `pages/index.tsx`. The page auto-updates as you edit and save the file.

---

# Local Development Setup Guide

This guide is for running `@synapsecns/synapse-interface` and `@synapsecns/sdk-router` simultaneously for local development, with continuous recompilation on changes.

## Prerequisites

Ensure you've installed Node.js (version 18.17.0) and Yarn on your machine. This setup assumes you're using Yarn Workspaces and Lerna to manage your project, with `@synapsecns/sdk-router` and `@synapsecns/synapse-interface` as part of the same workspace.

## Steps

1. **Install dependencies**
   From the root directory of your workspace, run:

```shell
yarn install
```

This will handle dependency installation and local package linking.

2. **Watch for changes in `@synapsecns/sdk-router`**
   Open a terminal, navigate to the workspace root, and run:

```shell
lerna run --scope @synapsecns/sdk-router start --stream
```

This triggers TSDX in watch mode for `@synapsecns/sdk-router`, triggering rebuilds on file changes.

3. **Run the Next.js application in development mode**
   In a separate terminal window, navigate to the `synapse-interface` directory and start the dev server:

```shell
yarn dev
```

This command watches for file changes and automatically rebuilds the application, including updated dependencies.

After completing these steps, any changes to `@synapsecns/sdk-router` will be automatically detected and rebuilt. The `@synapsecns/synapse-interface` application will then pick up and incorporate these updates.

Make sure the `@synapsecns/sdk-router` dependency in `synapse-interface`'s `package.json` is declared by name and version (like `"@synapsecns/sdk-router": "0.1.0"`), matching `sdk-router`'s `package.json` version.

---
