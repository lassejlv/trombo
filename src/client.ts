import { Client, Message } from 'discord.js'
import { log } from './lib/logger'

export interface Command {
  name: string
  description: string
  aliases: string[]
  run: (message: Message, args: string[], client: Client<true>) => Promise<void>
}

const commands = new Map<string, Command>()

export class Trombo extends Client {
  constructor() {
    super({ intents: ['Guilds', 'GuildMessages', 'MessageContent', 'GuildMessageReactions'] })
  }

  async add_command(command: Command) {
    commands.set(command.name, command)
    log.info(`Added command ${command.name}`)
  }

  async on_message(message: Message) {
    if (message.author.bot) return

    const prefix = '+'

    if (message.content.startsWith(prefix)) {
      const args = message.content.slice(prefix.length).trim().split(/ +/)
      const commandName = args.shift()?.toLowerCase()
      if (!commandName) return

      if (commands.has(commandName)) {
        const command = commands.get(commandName)
        if (!command) return
        await command.run(message, args, message.client)
      }
    }
  }

  async run() {
    await this.login(Bun.env.DISCORD_TOKEN)
  }

  async stop() {
    await this.destroy()
  }
}
