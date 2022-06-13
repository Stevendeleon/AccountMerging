package accounts

import (
	a "github.com/stevendeleon/accountmerging/src/accounts"
	"testing"
)

// TestLoadJSON tests the LoadJSON function
func TestLoadJSON(t *testing.T) {
	accs := a.LoadJSON("../../src/accounts.json")
	if len(accs) != 4 {
		t.Errorf("Expected 4 accounts, got %d", len(accs))
	}
}

// TestMerge tests the Merge function
func TestMerge(t *testing.T) {
	// load the accounts through the LoadJSON function
	accs := a.LoadJSON("../../src/accounts.json")

	// merge the accounts
	a.Merge(accs)

	if len(a.Output) != 2 {
		t.Errorf("Expected 2 accounts, got %d", len(a.Output))
	}

	if len(a.Output[0].Applications) != 3 {
		t.Errorf("Expected 3 applications, got %d", len(a.Output[0].Applications))
	}

	if len(a.Output[0].Emails) != 5 {
		t.Errorf("Expected 5 emails, got %d", len(a.Output[0].Emails))
	}

}
