exec = require('child_process').exec

start = ->
  jasmine = boundExec("jasmine-node --color --coffee --forceexit #{__dirname}/specs/")
  jasmine.on 'exit', (code) ->
    beatstorm.kill() if beatstorm?
    exec('killall beatstorm')
    code = -1 unless code?
    process.exit(code)

boundExec = (cmd) ->
  c = exec(cmd, {timeout: 10000})
  c.stdout.pipe(process.stdout)
  c.stderr.pipe(process.stderr)
  return c

exec('killall beatstorm')
beatstorm = boundExec('BS_ENV=test go run app/beatstorm.go config.json')
setTimeout(start, 4000)
