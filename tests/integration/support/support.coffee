RequestBuilder = require('./request_builder')
redis = require('redis')
pg = require('pg')
r = null
p = null
u = null

module.exports =
  redis: -> r
  pg: -> p
  request: (url, query = {}) -> new RequestBuilder(url, query)
  last_header: (callback) -> r.rpop 'headers', (err, header) -> callback(JSON.parse(header))
  last_url: (callback) -> r.rpop 'urls', (err, url) -> callback(url)
  last_raw: (callback) -> r.rpop 'raw', (err, url) -> callback(url)
  all_urls: (callback) -> r.lrange 'urls', 0, -1, (err, urls) -> callback(urls)
  prepare: (done) ->
    purge = -> r.select 15, -> r.flushdb ->
      p.query "truncate table beatvectors;", done
    if r? && p? then purge()
    else
      p = new pg.Client("postgres://localhost/bs_test")
      r = redis.createClient(6379, 'localhost', {enable_offline_queue: false})
      r.on 'ready', -> p.connect (err) ->
        console.log(err) if err?
        purge()
