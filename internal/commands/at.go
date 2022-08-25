package commands

import (
	"io/ioutil"

	"github.com/oidc-mytoken/api/v0"
	mytokenlib "github.com/oidc-mytoken/lib"
	"github.com/urfave/cli/v2"

	"github.com/oidc-mytoken/client/internal/config"
)

var atCommand = struct {
	PTOptions
	Scopes    cli.StringSlice
	Audiences cli.StringSlice
	Out       string
}{}

func init() {
	app.Commands = append(
		app.Commands, &cli.Command{
			Name: "AT",
			Aliases: []string{
				"at",
				"access-token",
			},
			Usage:  "Obtain an OIDC access token",
			Action: getAT,
			Flags: append(
				getPTFlags(),
				&cli.StringSliceFlag{
					Name:        "scope",
					Aliases:     []string{"s"},
					Usage:       "Request the passed scope.",
					DefaultText: "all scopes allowed for the used mytoken",
					Destination: &atCommand.Scopes,
					Placeholder: "SCOPE",
				},
				&cli.StringSliceFlag{
					Name:        "aud",
					Aliases:     []string{"audience"},
					Usage:       "Request the passed audience.",
					Destination: &atCommand.Audiences,
					Placeholder: "AUD",
				},
				&cli.StringFlag{
					Name:        "out",
					Aliases:     []string{"o"},
					Usage:       "The access token will be printed to this output",
					Value:       "/dev/stdout",
					Destination: &atCommand.Out,
					Placeholder: "FILE",
				},
			),
		},
	)
}

func getAT(context *cli.Context) error {
	atc := atCommand
	var comment string
	if context.Args().Len() > 0 {
		comment = context.Args().Get(0)
	}
	if ssh := atc.SSH(); ssh != "" {
		req := mytokenlib.NewAccessTokenRequest("", "", atc.Scopes.Value(), atc.Audiences.Value(), comment)
		return doSSH(ssh, api.SSHRequestAccessToken, req)
	}
	mytoken := config.Get().Mytoken
	provider, mToken := atc.Check()
	atRes, err := mytoken.AccessToken.APIGet(
		mToken, provider, atc.Scopes.Value(), atc.Audiences.Value(), comment,
	)
	if err != nil {
		return err
	}
	if atRes.TokenUpdate != nil {
		updateMytoken(atRes.TokenUpdate.Mytoken)
	}
	return ioutil.WriteFile(atc.Out, append([]byte(atRes.AccessToken), '\n'), 0600)
}
