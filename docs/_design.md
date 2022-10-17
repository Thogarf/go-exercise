# MVP 

## Architecture

**Data Storage** will be satisfied by using sqlite to store relational data on exercise models, routines combining exercises, and with stretch implementations involving user statistics. This choice is motivated by desiring a lightweight storage solution, something easily containerizable, and one that has a readily available implementation (multiple in this case) in Go. [See here](https://pkg.go.dev/gorm.io/driver/sqlite#section-readme).

Use of the **Discord API** is required to make this a bot. Users will create a Discord App, turn it into a bot, and register its key into this code. Developers will use the [DiscordGo](https://github.com/bwmarrin/discordgo) library to implement calls to Discord.

Between storing data and making calls to the Discord API, the rest of the application is expected to run as a native Go application. 

## Data

**Exercises** contain a unique identifier, a name, a list of muscle areas, and a customizeable list of tags. 
- The list of muscle areas should be a best approximation of which muscles are most-worked in the exercise. People adding exercises should be generous in expanding this list to include more muscles, as opposed to "rounding down".
- The list of tags should be customizeable by the user. Tags are human-readable strings that allows users to recognize the benefits or process of an exercise at a glance.

**Routines** are sequences of *exercises* that have a pointer to the current day's exercise, the calendar date corresponding to that pointer/exercise, and knowledge of what exercise comes next. By default, *routines* are tied to a calendar, where one *exercise* corresponds to one *calendar day*, and moving the pointer to the next *exercise* also assumes that it will be notified on the next *calendar day*. 

*Routines* drive the cadence and content of notifications by the bot.

For MVP, we do not expect to store any other unique or crucial data tables or models. Future efforts might see additional tables for streaks, stats, and may contain links or descriptions.

## Cadence

For MVP, we expect that the bot will only send notifications to the registered channel **once when the bot starts up** and is functionally healthy, and **once per command issued by the user**. 

To begin with, it will be feasible that the bot could be offline for a period of time and, once started, will use its database and the passage of time to correctly identify where it is in a given routine. 