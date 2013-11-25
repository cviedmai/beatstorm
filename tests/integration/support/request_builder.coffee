http = require('http')
crypto = require('crypto')

class RequestBuilder
  constructor: (@url, @query) ->
   @set_host("bs.beatstorm.net")
   @parse = true

  get: (callback) -> @make_request('get', callback)
  post: (data, callback) -> @make_request('post', data, callback)
  purge: (callback) -> @make_request('purge', callback)

  no_parse: ->
    @parse = false
    @

  sign: ->
    @query.t = parseInt(new Date().getTime()/1000, 10)
    hmac = crypto.createHmac('sha1', 'grapes')
    hmac.update(new Buffer(@url + @build_qs(), 'utf-8'))
    @query.sig = hmac.digest('hex')
    @

  header_sign: ->
    @sign()
    @add_header('signature', @query.sig)
    @add_header('timestamp', @query.t)
    delete @query.sig
    delete @query.t
    @

  add_header: (key, value) ->
    @headers = {} unless @headers?
    @headers[key] = value
    @

  set_host: (host) ->
    @add_header("Host", host)
    @

  make_request: (method, data, callback) =>
    if data instanceof Function
      callback = data
      data = null
    options =
      hostname: '127.0.0.1'
      port: 9095
      path: @url + @build_qs()
      method: method
    options.headers = @headers if @headers?

    response = ''
    req = http.request options, (res) =>
      res.on 'data', (d) -> response += d
      res.on 'error', (err) -> callback(res.statusCode, err)
      res.on 'end', =>
        r = if @parse then JSON.parse(response) else response
        callback(res.statusCode, r, res.headers)
    req.on 'error', (err) -> callback(0, err)
    req.end(JSON.stringify(data))

  build_qs: ->
    parts = ("#{k}=#{encodeURIComponent(v)}" for k, v of @query)
    if parts.length == 0 then '' else '?' + parts.join('&')

module.exports = RequestBuilder
