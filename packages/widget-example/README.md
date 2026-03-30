# Widget Example

Minimal Next.js consumer package for the Synapse bridge widget.

## Run locally

From the repository root:

```bash
yarn install
yarn workspace @synapsecns/widget build
yarn workspace @synapsecns/widget-example dev
```

Open `http://localhost:3000` to view the example.

## Iterate on the widget package

`@synapsecns/widget` resolves to the local workspace package and serves built files from `dist`, so rebuild it when the widget source changes.

Use a second terminal from the repository root:

```bash
yarn workspace @synapsecns/widget watch
```

If you switch branches or clear build artifacts, rerun:

```bash
yarn workspace @synapsecns/widget build
```

## Notes

- The example uses the Next `pages/` router to match the repo's standalone frontend packages.
- The widget only renders on the client to avoid SSR access to browser-only APIs.
- An injected Ethereum wallet is required to provide the browser provider.
- Account connection is optional: the example now renders the widget in read-only mode until a wallet account is connected.
