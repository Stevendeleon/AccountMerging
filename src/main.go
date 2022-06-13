package main

import (
	"github.com/stevendeleon/accountmerging/src/accounts"
)

func main() {
	accs := accounts.LoadJSON("accounts.json")
	accounts.Merge(accs)
}
