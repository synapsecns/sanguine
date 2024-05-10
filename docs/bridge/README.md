# Synapse Bridge Docs

Note, because of [this](https://github.com/PaloAltoNetworks/docusaurus-openapi-docs/issues/580#issuecomment-2103047228) issue, docs currently live outside of the yarn workspace. This will be fixed in a future version.

## Generating API Docs

`yarn docusaurus gen-api-docs all`: <!--todo: needs to be done from ci to ensure regenration is done-->

### Installation

```
$ yarn
```

### Local Development

```
$ yarn start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

### Build

```
$ yarn build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.
