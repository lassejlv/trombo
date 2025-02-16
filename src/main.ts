import { Trombo, type Command } from './client'
import { log } from './lib/logger'
import { resolve } from 'path'

const trombo = new Trombo()

const search_commands = new Bun.Glob('./src/commands/**/*.ts')

for await (const command of search_commands.scan('.')) {
  const full_path = resolve(command)
  const imported_command = await import(full_path).then((module) => module.default as Command)
  await trombo.add_command(imported_command)
}

trombo.on('messageCreate', (message) => trombo.on_message(message))

trombo
  .run()
  .then(() => log.info('Trombo is running'))
  .catch((error) => log.error(error))
