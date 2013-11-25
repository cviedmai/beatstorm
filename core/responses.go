package core

import (
  "github.com/cviedmai/auwfg"
)

var (
  NotFoundResponse = auwfg.Json(`{"code":1404,"error":"Not found"}`).Status(404)
  InvalidRequestResponse = auwfg.Json(`{"code":1400,"error":"invalid request"}`).Status(400)
  InternalErrorResponse = auwfg.Json(`{"code":1500, "error":"There was an error in the application"}`).Status(500)
  CreatedResponse = auwfg.Json(`{"ok":true}`).Status(201)
)
