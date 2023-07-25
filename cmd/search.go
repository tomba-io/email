package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomba-io/email/pkg/output"
	"github.com/tomba-io/email/pkg/start"
	"github.com/tomba-io/email/pkg/util"
	_domain "github.com/tomba-io/email/pkg/validation/domain"
	_key "github.com/tomba-io/email/pkg/validation/key"
)

// searchCmd represents the search command
// see https://developer.tomba.io/#domain-search
var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Instantly locate email addresses from any company name or website.",
	Run:     searchRun,
	Example: searchExample,
}

// searchRun the actual work search
func searchRun(cmd *cobra.Command, args []string) {
	fmt.Println(Long)
	init := start.New(conn)
	if init.Key == "" || init.Secret == "" {
		fmt.Println(util.WarningIcon(), util.Yellow(start.ErrErrInvalidNoLogin.Error()))
		return
	}
	if !_key.IsValidAPI(init.Key) && !_key.IsValidAPI(init.Secret) {
		fmt.Println(util.WarningIcon(), util.Yellow(start.ErrErrInvalidLogin.Error()))
		return
	}
	domain := init.Target
	if !_domain.IsValidDomain(domain) {
		fmt.Println(util.ErrorIcon(), util.Red(start.ErrArgumentsDomain.Error()))
		return

	}
	result, err := init.Tomba.DomainSearch(domain)
	if err != nil {
		fmt.Println(util.ErrorIcon(), util.Red(start.ErrErrInvalidLogin.Error()))
	}
	if init.JSON {
		raw, _ := result.Marshal()
		json, _ := output.DisplayJSON(string(raw))
		fmt.Println(json)
		return
	}
	if init.YAML {
		raw, _ := result.Marshal()
		json, _ := output.DisplayYAML(string(raw))
		fmt.Println(json)
		return
	}
}
