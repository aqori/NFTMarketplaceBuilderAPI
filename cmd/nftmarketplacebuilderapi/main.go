// cmd/nftmarketplacebuilderapi/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacebuilderapi/internal/nftmarketplacebuilderapi"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacebuilderapi.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
