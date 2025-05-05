# OpBot

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/opbot.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/opbot)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/opbot)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/opbot)

![icon.png](icon.png)
<!-- apoligies, this one was all chatgpt.-->

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

- `slack_bot_token`: The [bot token](https://api.slack.com/concepts/token-types#bot) for your Slack bot.
- `slack_app_token`: The [app token](https://api.slack.com/concepts/token-types#app-level) for your Slack app.
- `signoz_email`: The email address used to log in to [Signoz](https://signoz.io/docs/userguide/authentication/).
- `signoz_password`: The password used to log in to Signoz.
- `signoz_base_url`: The base URL for the Signoz API instance (example: http://mysignoz )

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

Certainly! I'll provide a step-by-step guide on how to add a new command to OpBot based on the information available in the provided code files.

### Adding a Command

1. **Create a new command function**
   In the `botmd/botmd.go` file, add a new method to the `Bot` struct. This method should return a `*slacker.CommandDefinition`. For example:

   ```go
   func (b *Bot) newCommand() *slacker.CommandDefinition {
       return &slacker.CommandDefinition{
           Command: "your-command <argument>",
           Description: "Description of your command",
           Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
               // Command logic goes here
           },
       }
   }
   ```

2. **Add the command to the bot**
   In the `NewBot` function within `botmd/botmd.go`, add your new command to the `addCommands` call:

   ```go
   bot.addCommands(
       bot.traceCommand(),
       bot.rfqLookupCommand(),
       bot.rfqRefund(),
       bot.newCommand(), // Add your new command here
   )
   ```

3. **Implement command logic**
   In the `Handler` function of your command, implement the logic for your command. You can access bot resources like `b.signozClient`, `b.rpcClient`, etc., to interact with different services.

4. **Add any necessary configuration**
   If your command requires additional configuration, add the necessary fields to the `config.Config` struct in the `config/config.go` file.

5. **Update the README**
   Add information about your new command to the README.md file, including its usage and any new configuration options.

6. **Test your command**
   Rebuild the bot and test your new command in Slack to ensure it works as expected.
