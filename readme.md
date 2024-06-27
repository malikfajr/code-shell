# Golang Telegram BOT: Code Shell Execution

This project is a Telegram bot that allows users to execute shell commands via Telegram. It is built using Golang version 1.22 and leverages the `github.com/spf13/viper` library for configuration management and `github.com/go-telegram/bot` for interacting with the Telegram Bot API.

## Features

- Execute shell commands via Telegram messages.
- Secure and configurable using environment variables.
- Easy to deploy and run.

## Requirements

- Golang 1.22
- Telegram Bot API token
- Shell environment

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/malikfajr/code-shell.git
    cd code-shell
    ```

2. **Install dependencies:**

    ```bash
    go mod download
    ```

3. **Set up environment variables:**

    Create a `.env` file in the project root and add the following variables:

    ```env
    TELEGRAM_BOT_TOKEN=your-telegram-bot-token
    ALLOWED_USERS=user1,user2,user3
    ```

    - `TELEGRAM_BOT_TOKEN`: Your Telegram bot token obtained from BotFather.
    - `ALLOWED_USERS`: Comma-separated list of Telegram usernames allowed to execute commands.

## Usage

1. **Run the bot:**

    ```bash
    go run main.go
    ```

2. **Interact with the bot:**

    Open your Telegram app and send a message to your bot to execute shell commands. Only users specified in the `ALLOWED_USERS` environment variable will be able to execute commands.

## Configuration

This project uses `viper` for configuration management. You can configure the bot using environment variables.

### Example `.env` file:

```env
TELEGRAM_BOT_TOKEN=your-telegram-bot-token
ALLOWED_USERS=user1,user2,user3
```

Viper will automatically read the variables from the `.env` file.

## Libraries Used

- [github.com/spf13/viper](https://github.com/spf13/viper): Viper is a complete configuration solution for Go applications.
- [github.com/go-telegram/bot](https://github.com/go-telegram/bot): A Telegram Bot API library for Golang.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.