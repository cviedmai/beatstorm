package beatvectors

import(
  "beatstorm/core"
  "github.com/cviedmai/auwfg"
  "beatstorm/models/beatvector"
  "github.com/cviedmai/auwfg/validation"
)

func init() {
  validation.Define("beatvector.uri", "uri", "uri is required", validation.Required())
  validation.Define("beatvector.data", "data", "data is required", validation.Required())
}

type CreateInput struct {
  Uri string `json:"uri"`
  Data string `json:"data"`
}

var authenticate = func(c *core.Context) auwfg.Response { return core.Authenticate(c) }

func Create(context *core.Context) auwfg.Response {
  if res := authenticate(context); res != nil { return res }
  input := context.Body.(*CreateInput)
  if res, valid := validateCreate(input); valid == false { return res }
  bv := beatvector.New()
  bv.Uri = input.Uri
  bv.Data = input.Data
  if err := bv.Save(); err != nil {
    return auwfg.Fatal(err)
  }
  return core.CreatedResponse
}

func validateCreate(input *CreateInput) (auwfg.Response, bool) {
  validator := auwfg.Validator()
  validator.Validate(input.Uri, "beatvector.uri")
  validator.Validate(input.Data, "beatvector.data")
  return validator.Response()
}
