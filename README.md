# golang-random-key
golang project generate random key defined by length


how use randomkey
in this example generate a random key of lenght 30 
```golang
package main

import (
        "github.com/swloopit/golang-random-key"
        "fmt"
)

func main(){
     fmt.Println(randomkey.CreateRandomKey(30))
}
```
