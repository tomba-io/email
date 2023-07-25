package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomba-io/email/pkg/output"
	"github.com/tomba-io/email/pkg/start"
	"github.com/tomba-io/email/pkg/util"
	_email "github.com/tomba-io/email/pkg/validation/email"
)

// verifyCmd represents the verify command
// see https://developer.tomba.io/#email-verifier
var verifyCmd = &cobra.Command{
	Use:     "verify",
	Aliases: []string{"t"},
	Short:   "Verify the deliverability of an email address.",
	Long:    Long,
	Run:     verifyRun,
	Example: verifyExample,
}

// verifyRun the actual work verify
func verifyRun(cmd *cobra.Command, args []string) {
	fmt.Println(Long)
	init := start.New(conn)
	if init.Key == "" || init.Secret == "" {
		fmt.Println(util.WarningIcon(), util.Yellow(start.ErrErrInvalidNoLogin.Error()))
		return
	}
	email := init.Target
	if !_email.IsValidEmail(email) {
		fmt.Println(util.ErrorIcon(), util.Red(start.ErrArgumentEmail.Error()))
		return
	}

	result, err := init.Tomba.EmailVerifier(email)
	if err != nil {
		fmt.Println(util.ErrorIcon(), util.Red(start.ErrErrInvalidLogin.Error()))
		return
	}
	if result.Data.Email.Email != "" {
		if result.Data.Email.Disposable {
			fmt.Println(util.WarningIcon(), util.Bold("The domain name is used by a disposable email addresses provider."))
			fmt.Println(util.WarningIcon(), util.Yellow("Tomba is designed to contact other professionals. This email is used to create personal email addresses so we don't the verification. 💡"))
			return
		}
		if result.Data.Email.Webmail {
			fmt.Println(util.WarningIcon(), util.Bold("The domain name  is webmail provider."))
			fmt.Println(util.WarningIcon(), util.Yellow("Tomba is designed to contact other professionals. This email is used to create personal email addresses so we don't the verification. 💡"))
			return
		}
		if init.JSON {
			raw, _ := result.Marshal()
			json, _ := output.DisplayJSON(string(raw))
			fmt.Println(json)
			return
		}
		if init.YAML {
			raw, _ := result.Marshal()
			yaml, _ := output.DisplayYAML(string(raw))
			fmt.Println(yaml)
			return
		}
		return
	}
	fmt.Println(util.WarningIcon(), util.Yellow("The Email Verification failed because of an unexpected response from the remote SMTP server. This failure is outside of our control. We advise you to retry later."))
}
