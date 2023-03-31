# gotextme

Send text messages to yourself via CLI. Primary use case: as a notifier for completion of long-running jobs.

## Usage

```bash
sleep 30 && gotextme "I slept for 30 seconds"
```

## Installation

1. Clone the repository 
2. Check if Go is available on your system: `go version`
3. Compile and install: `make install`.

NOTE: At present, this requires `root`-level priviledges, since the destination directory is `/usr/bin`. (TODO: make the app installable with `go isntall`, like the normal people do).

4. Make sure that a properly formatted INI file named `.gotextme.ini` is under your home directory root.

Configuration file example:

```ini
[telegram]
botApiToken = your_token
chatId = your_group_chat_id
```

## Messengers

### Telegram

In order to use this tool, you need a Telegram bot. Follow these steps to set up one (this is free):

1. Create a bot by contacting @BotFather 
2. Get the bot's API token from @BotFather
3. Create a new Telegram group chat and add your bot. You'll be sending messages to that chat.
4. Fetch bot updates and look for the chat id:
```bash
curl "https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/getUpdates" | jq .result[0].my_chat_member.chat.id
```
