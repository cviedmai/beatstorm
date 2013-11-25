#BeatStorm API (v1)
##Get BeatVector
`GET /v1/beatvectors/URI.json`

Will return a json file with the beatvector data. E.g. `{"uri":"abc", "data":"dce"}`

##Post BeatVector
`POST /v1/beatvectors.json` (requires signature)

The body should be a json file with the following attributes:

 - `Uri`: The Uri of the track
 - `Data`: The actual beatvector data

e.g.: `{"uri":"abc", "data":"dce"}`

##Signatures
Making requests
An SHA1-HMAC signature and timestamp should be included with every request in order to authenticate it.

The timestamp is the current time (unix timestamp) assigned to the t parameter. The signature is the url (including querystring) + body text (if any).

###Example
For example, given a secret of 'caracoles' as well as the following POST request:

`/v1/beatvectors.json?arriba=abajo`

The first thing we would do is add the t paremeter:

`/v1/beatvectors.json?arriba=abajo&t=1234567890`

We would then get the SHA1-HMAC value of the above. In ruby, this looks like:

`require 'openssl'
key = "caracoles"
data = "/v1/beatvectors/abcde.json?arriba=abajo&t=1234567890"
sig = OpenSSL::HMAC.hexdigest(OpenSSL::Digest::Digest.new('sha1'), key, data)`

And append it to our querystring:

`/v1/beatvectors/abcde.json?arriba=abajo&t=1234567890&sig=asdklsajdasldkj`

###Invalid responses
In the case of an invalid signature, the following error will be returned (HTTP status 401):

`{"code":1401,"error":"Unathorized, invalid signature"}`

In the case of an invalid timestamp (if the client/server are out of sync), the following error will be returned (HTTP status 401):

`{"code":1402,"error":"Unathorized, invalid timestamp", "current_timestamp": SERVER_TIMESTAMP}`

The command should be re-issued using the returned SERVER_TIMESTAMP.
