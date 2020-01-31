package main

import (
  //"os"
)

func main() {
	router := newRouter()
  // 本番環境はこっち (os を import する)
  //router.Logger.Fatal(router.Start(":"+os.Getenv("PORT")))

	router.Logger.Fatal(router.Start(":8080"))
}
