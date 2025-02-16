import type { Command } from '../client'

export default {
  name: 'ping',
  description: 'Replies with pong!',
  aliases: ['p'],
  run: async (message) => {
    await message.reply('pong!')
  },
} satisfies Command
