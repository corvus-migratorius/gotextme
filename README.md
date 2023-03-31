# gotextme

Send text messages to yourself via CLI. Primary use case: as a notifier for completion of long-running jobs.

## Usage

```bash
sleep 30; gotextme "I slept for 30 seconds"
```

If you want different messages on exit statuses != 0, try the following:

```bash
sleep 30 && gotextme "I slept for 30 seconds" || gotextme "Sleep has failed, somehow"
```

## Installation

### Use a pre-compiled binary

1. Grab a binary from the [Releases](https://github.com/corvus-sapiens/gotextme/releases) page and put it somewhere under you `$PATH`;
2. Set up a Telegram bot - it's free!
3. Create a `.gotextme.ini` file with the appropriate bot credentials under you `$HOME`;

```ini
[telegram]
botApiToken = your_token
chatId = your_group_chat_id
```

## Telegram setup

In order to use this tool, you need a Telegram bot. Follow these steps to set up one (this is free):

1. Create a bot by contacting @BotFather 
2. Get the bot's API token from @BotFather
3. Create a new Telegram group chat and add your bot. You'll be sending messages to that chat.
4. Fetch bot updates and look for the chat id:

```bash
curl "https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/getUpdates" | jq .result[0].my_chat_member.chat.id
```
