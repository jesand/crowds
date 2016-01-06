package main

import (
	"encoding/json"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/jesand/crowds/amt"
	"os"
)

const (
	USAGE = `hiclusterd - Web service hosting for hicluster

Usage:
  amtadmin balance --amt=<path> [--sandbox]
  amtadmin -h | --help
  amtadmin --version

Options:
  balance       Get the account balance
  --amt=<path>  The path to a file containing AMT credentials
  --sandbox     Address the AMT sandbox instead of the production site
`
)

type AmtCred struct {
	AccessKey, SecretKey string
}

func main() {

	// Parse the command line
	args, _ := docopt.Parse(USAGE, nil, true, "1.0", false)

	// Initialize the AMT client
	var (
		credPath = args["--amt"].(string)
		sandbox  = args["--sandbox"].(bool)
		amtCred  AmtCred
		client   *amt.AmtClient
	)
	if f, err := os.Open(credPath); err != nil {
		fmt.Printf("Error: Could not open %s - %v", credPath, err)
		return
	} else if err = json.NewDecoder(f).Decode(&amtCred); err != nil {
		fmt.Printf("Error: Could not parse %s - %v", credPath, err)
		return
	} else {
		client = amt.NewClient(amtCred.AccessKey, amtCred.SecretKey, sandbox)
	}

	switch {
	case args["balance"].(bool):
		RunBalance(client)
	}
}

func RunBalance(client *amt.AmtClient) {
	balance, err := client.GetAccountBalance()
	if err != nil {
		fmt.Printf("Error: Could not check account balance: %v\n", err)
		return
	}
	fmt.Println("Current balance:",
		balance.GetAccountBalanceResults[0].AvailableBalance.FormattedPrice)
}
