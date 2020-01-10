# gebal
Golang Library for European bank accounts

## Examples

### Spain

Lenght: 20 chars
Example: "11112222334444444444"
Content:
- 1111 - Bank identifier
- 2222 - Branch identifier
- 33 - Control digit
- 4444444444 - Account number

#### Generate Control Digit
```
package main

import (
	"github.com/aluvare/gebal/gebal-es"
	"log"
)

func main() {
	res, err := gebales.CCCCdGen("1111", "2222", "4444444444")
	log.Println(res, err)
}
```
#### Check 
```
package main

import (
	"github.com/aluvare/gebal/gebal-es"
	"log"
)

func main() {
	res, err := gebales.CCCCdCheck("1111", "2222", "33", "4444444444")
	log.Println(res, err)
}
```
