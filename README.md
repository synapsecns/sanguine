# Installing for local development and testing

## Widget

1. Clone repo

```
git clone git@github.com:abtestingalpha/widget.git
```

2. Enter widget directory

```
cd widget
```

3. Install dependencies for widget

```
yarn
```

4. Use right node version

```
nvm use
```

5. Create a link to the widget package to allow user in consumer app

```
yarn link
```

6. Create a build to be used by consumer app. This is live reload (through Rollup's `watch` flag), which will need to remain running so changes to the widget code are updated in the build (which is referenced by the consumer app).

```
yarn watch
```

## Example React consumer

1. Open up another terminal window

2. Enter React app directory

```
cd examples/with-react
```

3. Install dependencies for React app

```
yarn
```

4. Link widget to allow use in consumer app

```
yarn link @synapsecns/widget
```

5. Start consumer app

```
yarn start
```
