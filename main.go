package main

import (
  service "./service"
)

func main() {
  port := "8080"
  server := service.NewServer()
  server.Run(":" + port)

}
