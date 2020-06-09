package command

import (
	"path"
	"strings"

	"github.com/masterzen/dashlane-cli/pkg/dashlane"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"

	"github.com/alecthomas/kong"
	"github.com/sirupsen/logrus"
)

type debugFlag bool

// Context of the app
type Context struct {
	Dl *dashlane.Dashlane
}

type cli struct {
	Debug debugFlag `help:"Enable debug logging."`

	Uki     UkiCmd     `cmd help:"Manage computer registration."`
	Vault   VaultCmd   `cmd help:"Manage the vault."`
	Version VersionCmd `cmd help:"Displays dahslane-cli version."`
}

// Execute the commands
func Execute(command string) {
	cmds := strings.Split(command, " ")

	dir, err := homedir.Dir()
	if err != nil {
		logrus.WithError(err).Error("Can't get user home directory")
	}
	dashlaneDir := path.Join(dir, ".dashlane")

	fs := afero.NewOsFs()
	DashlaneConfig := path.Join(dashlaneDir, "config.json")
	DashlaneVault := path.Join(dashlaneDir, "vault.json")

	context := &Context{
		Dl: dashlane.New(fs, DashlaneVault, DashlaneConfig),
	}
	fs.MkdirAll(dashlaneDir, 0700)

	cli := new(cli)
	kconfig := kong.Configuration(kong.JSON, DashlaneConfig)

	parser, err := kong.New(cli, kconfig)
	if err != nil {
		panic(err)
	}
	kongctx, err := parser.Parse(cmds)
	parser.FatalIfErrorf(err)

	err = kongctx.Run(context)
	kongctx.FatalIfErrorf(err)
}

func (d debugFlag) BeforeApply() error {
	logrus.SetLevel(logrus.DebugLevel)
	return nil
}
