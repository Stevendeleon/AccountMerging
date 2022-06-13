package main

import (
	"github.com/stevendeleon/accountmerging/src/accounts"
)

func main() {
	accs := accounts.LoadJSON("src/accounts.json")
	accounts.Merge(accs)
}
