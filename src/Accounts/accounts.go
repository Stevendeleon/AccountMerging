package accounts

import (
	"encoding/json"
	"fmt"
	"os"
)

type Accounts []struct {
	Application string   `json:"application"`
	Emails      []string `json:"emails"`
	Name        string   `json:"name"`
}

var (
	Accs Accounts
)

func LoadJSON(path string) Accounts {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Accs)
	if err != nil {
		fmt.Println(err)
	}

	return Accs
}

func Merge(a Accounts) {
	var mergedAccounts []struct {
		Applications []string
		Emails       []string
		Name         string
	}

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(mergedAccounts); j++ {
			if mergedAccounts[j].Name == a[i].Name {
				found = true
				break
			}
		}
		if !found {
			mergedAccounts = append(mergedAccounts, struct {
				Applications []string
				Emails       []string
				Name         string
			}{
				Applications: []string{a[i].Application},
				Emails:       a[i].Emails,
				Name:         a[i].Name,
			})
		} else {
			for j := 0; j < len(mergedAccounts); j++ {
				if mergedAccounts[j].Name == a[i].Name {
					mergedAccounts[j].Emails = append(mergedAccounts[j].Emails, a[i].Emails...)
					mergedAccounts[j].Applications = append(mergedAccounts[j].Applications, a[i].Application)
				}
			}
		}

	}

	fmt.Println(mergedAccounts)
}
