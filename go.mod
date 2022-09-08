module github.com/oidc-mytoken/client

go 1.16

require (
	github.com/Songmu/prompter v0.5.1
	github.com/evanphx/json-patch v0.5.2
	github.com/gliderlabs/ssh v0.3.4
	github.com/oidc-mytoken/api v0.8.0
	github.com/oidc-mytoken/lib v0.5.0
	github.com/oidc-mytoken/server v0.5.5-0.20220907095920-beac501d4181
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.0
	github.com/urfave/cli/v2 v2.3.1-0.20211205195634-e8d81738896c
	golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b
	golang.org/x/term v0.0.0-20220526004731-065cf7ba2467
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/urfave/cli/v2 => github.com/zachmann/cli/v2 v2.3.1-0.20211220102037-d619fd40a704
