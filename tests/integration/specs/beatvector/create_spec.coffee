support = require('../../support/support')

describe 'beatvector#create', ->
  beforeEach support.prepare

  it "works when valid data is passed", (done) ->
    support.request('/v1/beatvectors.json').post {uri:"super_uri", data:"super_data"}, (status, res) ->
      expect(res.ok).toEqual(true)
      done()
