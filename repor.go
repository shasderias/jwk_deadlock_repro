package main

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/httprc/v3"
	"github.com/lestrrat-go/jwx/v3/jwk"
)

func main() {
	ctx := context.Background()

	c, err := jwk.NewCache(ctx, httprc.NewClient())
	if err != nil {
		panic(err)
	}

	for i := 0; i < 50; i++ {
		u := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/certs/%d", i)
		fmt.Println("registering", u)
		if err := c.Register(ctx, u, jwk.WithWaitReady(false)); err != nil {
			panic(err)
		}
	}
}
