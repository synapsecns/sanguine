# Synapse Constants
[![npm](https://img.shields.io/npm/v/synapse-constants?style=flat-square)](https://www.npmjs.com/package/synapse-constants)

This package contains the Synapse Protocol Token and Chain Constants


#



## Installation

```bash
npm install synapse-constants
```

With Yarn:

```bash
yarn add synapse-constants
```

## Usage


To restrict the assets and chains that are imported, you can create a "custom bridge list". From the set of all tokens imported from "bridgeable.ts" you can import specific tokens and use that as the custom list you use in your application. The same can be done for chains

## Usage
For maintenance, when new tokens are added to the bridge the following steps should be taken.

1. Regenerate bridgeMaps.ts

```bash
yarn maps:generate
```

2. Update Bridgeable.ts with the new token addresses (check all other variables like decimals/ symbols etc. )

3. Repackage and webpack all of the data

```bash
yarn compile
```

4. Republish the npm package (make sure to update the version)

```bash
npm publish
```


TODO:
- add the basic structure of the token type and the chain type to show accessibility for token logos, chain logos, and any additional information.
