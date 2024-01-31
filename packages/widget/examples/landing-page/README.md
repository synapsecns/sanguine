# Installing for local development and testing

## Widget

1. Clone repo

```
git clone git@github.com:synapsecns/sanguine.git
```

2. Enter widget directory

```
cd sanguine/packages/widget
```

3. Use right node version

```
nvm use
```

4. Install dependencies for widget

```
yarn
```

5. Ensure local singular build works

```
yarn build
```

5. Create a link to the widget package to allow user in consumer app

```
yarn link
```

6. To use live reload for changes made in widget, we can use through Rollup's `watch` flag, we'll need a live builder, which updates the build based on changes made in the `widget` package

```
yarn watch
```

## Example Landing Page Next.js consumer

1. Open up another terminal window

2. Enter Landing Page app directory

```
cd examples/landing-page
```

3. Install dependencies for app

```
yarn
```

4. Link widget to allow use in consumer app

```
yarn link @synapsecns/widget
```

5. Start consumer app

```
yarn dev
```

### Deploying

1. When ready to deploy, make sure to add the `@synapsecns/widget` dependency to the `package.json` file of the Landing page app.
