# Env

Simple approach to environment variables.

`go get -u github.com/nubunto/env`

## Examples

```
package main

import (
    "fmt"
    "github.com/nubunto/env"
)

func main() {
    port := env.Get("PORT", env.Default(":9090"))
    fmt.Println("port is:", port)

    env.Set("PORT", ":9083")
    fmt.Println("port is:", env.Get("PORT")) // 9083

    hosts := map[string]string{
        "HOST1": ":9087",
        "HOST2": ":9082",
    }
    h, _ := os.Hostname()
    hosts[h] = ":7777"

    port = env.Get("PORT", env.Transform(value string, existsInSystem bool) string{
        h, err := os.Hostname()
        if err != nil {
            return value
        }
        return hosts[h]
    })
    fmt.Println("port is:", port) // 7777
}
```

# License

MIT
