// Copyright 2020 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package flags

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/params"
	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v2"
)

// usecolor defines whether the CLI help should use colored output or normal dumb
// colorless terminal formatting.
var usecolor = (isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())) && os.Getenv("TERM") != "dumb"

// NewApp creates an app with sane defaults.
func NewApp(gitCommit, gitDate, usage string) *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = params.VersionWithCommit(gitCommit, gitDate)
	app.Usage = usage
	app.Copyright = "Copyright 2013-2022 The go-ethereum Authors"
	app.Before = func(ctx *cli.Context) error {
		MigrateGlobalFlags(ctx)
		return nil
	}
	return app
}

var migrationApplied = map[*cli.Command]struct{}{}

// MigrateGlobalFlags makes all global flag values available in the
// context. This should be called as early as possible in app.Before.
//
// Example:
//
//	geth account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
//	geth --keystore /tmp/mykeystore --lightkdf account new
//
// i.e. in the subcommand Action function of 'account new', ctx.Bool("lightkdf)
// will return true even if --lightkdf is set as a global option.
//
// This function may become unnecessary when https://github.com/urfave/cli/pull/1245 is merged.
func MigrateGlobalFlags(ctx *cli.Context) {
	var iterate func(cs []*cli.Command, fn func(*cli.Command))
	iterate = func(cs []*cli.Command, fn func(*cli.Command)) {
		for _, cmd := range cs {
			if _, ok := migrationApplied[cmd]; ok {
				continue
			}
			migrationApplied[cmd] = struct{}{}
			fn(cmd)
			iterate(cmd.Subcommands, fn)
		}
	}

	// This iterates over all commands and wraps their action function.
	iterate(ctx.App.Commands, func(cmd *cli.Command) {
		action := cmd.Action
		cmd.Action = func(ctx *cli.Context) error {
			doMigrateFlags(ctx)
			return action(ctx)
		}
	})
}

func doMigrateFlags(ctx *cli.Context) {
	for _, name := range ctx.FlagNames() {
		for _, parent := range ctx.Lineage()[1:] {
			if parent.IsSet(name) {
				ctx.Set(name, parent.String(name))
				break
			}
		}
	}
}

func init() {
	if usecolor {
		// Annotate all help categories with colors
		cli.AppHelpTemplate = regexp.MustCompile("[A-Z ]+:").ReplaceAllString(cli.AppHelpTemplate, "\u001B[33m$0\u001B[0m")

		// Annotate flag categories with colors (private template, so need to
		// copy-paste the entire thing here...)
		cli.AppHelpTemplate = strings.ReplaceAll(cli.AppHelpTemplate, "{{template \"visibleFlagCategoryTemplate\" .}}", "{{range .VisibleFlagCategories}}\n   {{if .Name}}\u001B[33m{{.Name}}\u001B[0m\n\n   {{end}}{{$flglen := len .Flags}}{{range $i, $e := .Flags}}{{if eq (subtract $flglen $i) 1}}{{$e}}\n{{else}}{{$e}}\n   {{end}}{{end}}{{end}}")
	}
	cli.FlagStringer = FlagString
}

// FlagString prints a single flag in help.
func FlagString(f cli.Flag) string {
	df, ok := f.(cli.DocGenerationFlag)
	if !ok {
		return ""
	}

	needsPlaceholder := df.TakesValue()
	placeholder := ""
	if needsPlaceholder {
		placeholder = "value"
	}

	namesText := pad(cli.FlagNamePrefixer(df.Names(), placeholder), 30)

	defaultValueString := ""
	if s := df.GetDefaultText(); s != "" {
		defaultValueString = " (default: " + s + ")"
	}

	usage := strings.TrimSpace(df.GetUsage())
	envHint := strings.TrimSpace(cli.FlagEnvHinter(df.GetEnvVars(), ""))
	if len(envHint) > 0 {
		usage += " " + envHint
	}

	usage = wordWrap(usage, 80)
	usage = indent(usage, 10)

	if usecolor {
		return fmt.Sprintf("\n    \u001B[32m%-35s%-35s\u001B[0m%s\n%s", namesText, defaultValueString, envHint, usage)
	} else {
		return fmt.Sprintf("\n    %-35s%-35s%s\n%s", namesText, defaultValueString, envHint, usage)
	}
}

func pad(s string, length int) string {
	if len(s) < length {
		s += strings.Repeat(" ", length-len(s))
	}
	return s
}

func indent(s string, nspace int) string {
	ind := strings.Repeat(" ", nspace)
	return ind + strings.ReplaceAll(s, "\n", "\n"+ind)
}

func wordWrap(s string, width int) string {
	var (
		output     strings.Builder
		lineLength = 0
	)

	for {
		sp := strings.IndexByte(s, ' ')
		var word string
		if sp == -1 {
			word = s
		} else {
			word = s[:sp]
		}
		wlen := len(word)
		over := lineLength+wlen >= width
		if over {
			output.WriteByte('\n')
			lineLength = 0
		} else {
			if lineLength != 0 {
				output.WriteByte(' ')
				lineLength++
			}
		}

		output.WriteString(word)
		lineLength += wlen

		if sp == -1 {
			break
		}
		s = s[wlen+1:]
	}

	return output.String()
}
