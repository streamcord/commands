A microservice command parser and handler for Discord, written in Golang.

Part of a larger, yet-to-be-released stack using Nats and Redis to communicate between different parts of a Discord bot.

This microservice was intended to act as a processor for `MESSAGE_CREATE` events from the Discord gateway,
but since the [recent announcement](https://support-dev.discord.com/hc/en-us/articles/4404772028055)
from the Discord devs outlining changes that will make message contents a privileged intent,
the code was scrapped in favor of a specific handler for slash commands.
