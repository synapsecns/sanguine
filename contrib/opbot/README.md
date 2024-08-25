# OpBot

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/opbot.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/opbot)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/opbot)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/opbot)

OpBot is a Slack bot written in Go that interacts with the Signoz trace API to provide various functionalities, including searching for transactions based on user-provided tags. This bot is designed to help teams monitor and manage their operations more effectively by integrating with Slack and Signoz.

## Features

- **Slack Integration**: Interact with the bot directly from Slack.
- **Signoz Integration**: Search for transactions and traces using the Signoz API.
- **Configuration Management**: Easily manage configuration through YAML files.
- **Metrics Handling**: Integrated with metrics handling for better monitoring.

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/synapsecns/sanguine.git
    cd sanguine/contrib/opbot
    ```

2. **Install dependencies**:
   Ensure you have Go installed (version 1.22.4 or later). Then, run:
    ```sh
    go mod tidy
    ```

3. **Build the bot**:
    ```sh
    go build -o opbot main.go
    ```

## Configuration

OpBot uses a YAML configuration file to manage its settings. The configuration file should be named `config.yml` and placed in the same directory as the executable.

### Example `config.yml`

```yaml
slack_bot_token: "your-slack-bot-token"
slack_app_token: "your-slack-app-token"
signoz_email: "your-signoz-email"
signoz_password: "your-signoz-password"
signoz_base_url: "https://signoz.example.com"
```

Tokens can be obtained [here](https://api.slack.com/tutorials/tracks/getting-a-token). When creating an app, you can copy and paste the [manifest](manifest.json) file to configure the app automatically.

### Configuration Fields

- `slack_bot_token`: The token for your Slack bot.
- `slack_app_token`: The token for your Slack app.
- `signoz_email`: The email address used to log in to Signoz.
- `signoz_password`: The password used to log in to Signoz.
- `signoz_base_url`: The base URL for the Signoz API.

## Usage

1. **Start the bot**:
    ```sh
    ./opbot start --config config.yml
    ```

2. **Interact with the bot in Slack**:
  - Use commands to search for transactions in Signoz.
  - Example command: `/opbot search --tag key:value`

## Development

### Directory Structure

- **`cmd`**: Contains the command line interface for the bot.
- **`config`**: Provides functionality to read and write configuration files.
- **`botmd`**: Contains the main bot server implementation.
- **`metadata`**: Provides metadata services for the bot.
- **`signoz`**: Contains the Signoz client for interacting with the Signoz API.

Feel free to reach out if you have any questions or need further assistance!
