package main

import (
	"encoding/json"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/jesand/crowds/amt"
	xsdt "github.com/metaleap/go-xsd/types"
	"os"
	"reflect"
	"sort"
	"strconv"
)

const (
	USAGE = `hiclusterd - Web service hosting for hicluster

Usage:
  amtadmin balance --amt=<path> [--sandbox]
  amtadmin show [--hit=<id>] [--assn=<id>] --amt=<path> [--sandbox]
  amtadmin hits [--sort=<field>] [--desc] [--page=<num>] [--pageSize=<num>] ` +
		`--amt=<path> [--sandbox]
  amtadmin assns --hit=<id> [--status=<str>] [--sort=<field>] [--desc] ` +
		`[--page=<num>] [--pageSize=<num>] --amt=<path> [--sandbox]
  amtadmin -h | --help
  amtadmin --version

Options:
  balance           Get the account balance
  show              Display the status of a HIT or Assignment
  hits              Find matching HITs
  assns             Find assignments for a HIT
  --amt=<path>      The path to a file containing AMT credentials
  --sandbox         Address the AMT sandbox instead of the production site
  --hit=<id>        The ID of the HIT you want to view
  --assn=<id>       The ID of the assignment you want to view
  --sort=<field>    The field to sort by. For hits, one of: CreationTime,
                    Enumeration, Expiration, Reward, or Title. For assns, one
                    of: AcceptTime, SubmitTime, or AssignmentStatus.
  --status=<str>    The assignment status to search for. Can be:
                    Submitted, Approved, or Rejected.
  --desc            Sort results in descending order
  --page=<num>      The page number of results to display [default: 1]
  --pageSize=<num>  The number of results to display per page [default: 10]
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
		client   amt.AmtClient
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

	case args["show"].(bool):
		hitId, _ := args["--hit"].(string)
		assnId, _ := args["--assn"].(string)
		RunShow(client, hitId, assnId)

	case args["hits"].(bool):
		var (
			sort, _               = args["--sort"].(string)
			desc                  = args["--desc"].(bool)
			page, pageErr         = strconv.Atoi(args["--page"].(string))
			pageSize, pageSizeErr = strconv.Atoi(args["--pageSize"].(string))
		)
		if sort == "" {
			sort = "CreationTime"
		}
		if pageErr != nil {
			fmt.Printf("Invalid --page argument\n")
		} else if pageSizeErr != nil {
			fmt.Printf("Invald --pageSize argument\n")
		} else {
			RunHits(client, sort, desc, page, pageSize)
		}

	case args["assns"].(bool):
		var (
			hitId, _              = args["--hit"].(string)
			status, _             = args["--status"].(string)
			sort, _               = args["--sort"].(string)
			desc                  = args["--desc"].(bool)
			page, pageErr         = strconv.Atoi(args["--page"].(string))
			pageSize, pageSizeErr = strconv.Atoi(args["--pageSize"].(string))
			statuses              []string
		)
		if sort == "" {
			sort = "AcceptTime"
		}
		if status != "" {
			statuses = append(statuses, status)
		}
		if pageErr != nil {
			fmt.Printf("Invalid --page argument\n")
		} else if pageSizeErr != nil {
			fmt.Printf("Invald --pageSize argument\n")
		} else {
			RunAssns(client, hitId, statuses, sort, desc, page, pageSize)
		}
	}
}

func getObjectFields(object interface{}, vals map[string]string) {
	v := reflect.Indirect(reflect.ValueOf(object))
	if !v.IsValid() {
		return
	}
	t := v.Type()
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			getObjectFields(v.Index(i).Interface(), vals)
		}
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			switch field.Type.Kind() {
			case reflect.Struct, reflect.Ptr, reflect.Slice:
				getObjectFields(v.Field(i).Interface(), vals)
			default:
				if field.Type == reflect.TypeOf(xsdt.Int(0)) {
					vals[field.Name] = fmt.Sprintf("%d", v.Field(i).Interface())
				} else if field.Type == reflect.TypeOf(xsdt.Long(0)) {
					vals[field.Name] = fmt.Sprintf("%d", v.Field(i).Interface())
				} else {
					vals[field.Name] = v.Field(i).String()
				}
			}
		}
	}
}

func printObject(object interface{}) {
	var (
		fields   []string
		vals     = make(map[string]string)
		fieldLen int
	)
	getObjectFields(object, vals)
	for name, _ := range vals {
		fields = append(fields, name)
		if len(name) > fieldLen {
			fieldLen = len(name)
		}
	}
	sort.Strings(fields)
	format := fmt.Sprintf("%%%ds: %%s\n", fieldLen)
	for _, name := range fields {
		fmt.Printf(format, name, vals[name])
	}
}

func RunBalance(client amt.AmtClient) {
	balance, err := client.GetAccountBalance()
	if err != nil {
		fmt.Printf("Error: The AMT request failed: %v\n", err)
		return
	}
	printObject(balance)
}

func RunShow(client amt.AmtClient, hitId, assnId string) {
	switch {
	case hitId != "":
		if resp, err := client.GetHIT(hitId); err != nil {
			fmt.Printf("Error: The AMT request failed: %v\n", err)
			return
		} else if len(resp.Hits) > 0 && resp.Hits[0].Request != nil &&
			resp.Hits[0].Request.Errors != nil {

			printObject(resp.Hits[0].Request)
		} else {
			printObject(resp)
		}

	case assnId != "":
		if resp, err := client.GetAssignment(assnId); err != nil {
			fmt.Printf("Error: The AMT request failed: %v\n", err)
			return
		} else if len(resp.GetAssignmentResults) > 0 &&
			resp.GetAssignmentResults[0].Request != nil &&
			resp.GetAssignmentResults[0].Request.Errors != nil {

			printObject(resp.GetAssignmentResults[0].Request)
		} else {
			printObject(resp)
		}

	default:
		fmt.Println("You must provide a value for either --hit or --assn")
	}
}

func RunHits(client amt.AmtClient, sort string, desc bool, page, pageSize int) {
	if resp, err := client.SearchHITs(sort, !desc, pageSize, page); err != nil {
		fmt.Printf("Error: The AMT request failed: %v\n", err)
		return
	} else if len(resp.SearchHITsResults) > 0 &&
		resp.SearchHITsResults[0].Request != nil &&
		resp.SearchHITsResults[0].Request.Errors != nil {

		printObject(resp.SearchHITsResults[0].Request)
	} else if len(resp.SearchHITsResults[0].Hits) == 0 {
		fmt.Println("Found no HITs for this account")
	} else {
		for i, hit := range resp.SearchHITsResults[0].Hits {
			fmt.Printf("HIT %d/%d:\n", i+1, len(resp.SearchHITsResults))
			printObject(hit)
			fmt.Println()
		}
	}
}

func RunAssns(client amt.AmtClient, hitId string, statuses []string, sort string, desc bool, page, pageSize int) {
	if resp, err := client.GetAssignmentsForHIT(hitId, statuses, sort, !desc,
		pageSize, page); err != nil {

		fmt.Printf("Error: The AMT request failed: %v\n", err)
		return
	} else if len(resp.GetAssignmentsForHITResults) > 0 &&
		resp.GetAssignmentsForHITResults[0].Request != nil &&
		resp.GetAssignmentsForHITResults[0].Request.Errors != nil {

		printObject(resp.GetAssignmentsForHITResults[0].Request)
	} else if len(resp.GetAssignmentsForHITResults[0].Assignments) == 0 {
		fmt.Println("Found no assignments for this HIT")
	} else {
		for i, assn := range resp.GetAssignmentsForHITResults[0].Assignments {
			fmt.Printf("Assignment %d/%d:\n", i+1, len(resp.GetAssignmentsForHITResults))
			printObject(assn)
			fmt.Println()
		}
	}
}
