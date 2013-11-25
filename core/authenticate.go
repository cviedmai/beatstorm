package core

import (
  "fmt"
  "time"
  "regexp"
  "strconv"
  "crypto/hmac"
  "crypto/sha1"
  "github.com/cviedmai/auwfg"
)

var (
  secret []byte
  sigRegex, _ = regexp.Compile("&?sig=[^&]+")
  AuthenticationInvalidSignature = auwfg.Json(`{"code":1401,"error":"Unathorized, invalid signature"}`).Status(401)
)

func Authenticate(c *Context) auwfg.Response {
  if t := c.Query["t"]; len(t) == 0 || !validT(t) {
    return auwfg.Json(`{"code":1402,"error":"Unathorized, invalid timestamp", "current_timestamp":` + strconv.FormatInt(time.Now().Unix(), 10) + `}`).Status(401)
  }
  if sig := c.Query["sig"]; len(sig) == 0 || !validSig(c, sig) {
    return AuthenticationInvalidSignature
  }
  return nil
}

func validT(raw string) bool {
  t, err := strconv.ParseInt(raw, 10, 64)
  if err != nil { return false }
  return time.Now().Sub(time.Unix(t, 0)).Minutes() <= 30
}

func validSig(context *Context, sig string) bool {
  hasher := hmac.New(sha1.New, []byte(GetConfig().Secret))

  q := context.RawQuery
  q = sigRegex.ReplaceAllString(q, "")
  hasher.Write([]byte(context.RawPath + "?" + q))

  if context.RawBody != nil { hasher.Write(context.RawBody) }
  return fmt.Sprintf("%x", hasher.Sum(nil)) == sig
}
