package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rshulabs/micro-frame/pkg/errorx"
	"github.com/rshulabs/micro-frame/pkg/logx"
	"github.com/rshulabs/micro-frame/pkg/utils/term"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type App struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	// noVersion   bool
	noConfig bool
	commands []*Command
	args     cobra.PositionalArgs // 位置参数校验
	cmd      *cobra.Command
}

type RunFunc func(basename string) error

type Option func(app *App)

func WithOptions(opt CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

func WithDescription(desc string) Option {
	return func(a *App) {
		a.description = desc
	}
}

func WithNoConfig() Option {
	return func(a *App) {
		a.noConfig = true
	}
}

func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(a *App) {
		a.args = args
	}
}

func WithDefaultValidArgs() Option {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}

func NewApp(name string, basename string, opts ...Option) *App {
	app := &App{
		name:     name,
		basename: basename,
	}
	for _, opt := range opts {
		opt(app)
	}
	app.buildCommand()
	return app
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:   FormatBaseName(a.basename),
		Short: a.name,
		Long:  a.description,
		// 出现错误 停止打印usage
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          a.args,
	}
	// cmd.SetUsageTemplate(usageTemplate)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	// 参数选项
	cmd.Flags().SortFlags = true

	if len(a.commands) > 0 {
		for _, command := range a.commands {
			cmd.AddCommand(command.cobraCommand())
		}
		cmd.SetHelpCommand(helpCommand(FormatBaseName(a.basename)))
	}
	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	var flagSets FlagSets
	if a.options != nil {
		flagSets = a.options.Flags()
		fs := cmd.Flags()
		for _, f := range flagSets.Sets {
			fs.AddFlagSet(f)
		}
	}
	if !a.noConfig {
		addConfigFlag(flagSets.FlagSet("global"))
	}

	cmd.Flags().AddFlagSet(flagSets.FlagSet("global"))
	addHelpFlag(cmd.Name(), flagSets.FlagSet("global"))
	addCmdTemplate(&cmd, flagSets)
	a.cmd = &cmd
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	printWorkingDir()

	// 配置
	if !a.noConfig {
		// 绑定所有命令行参数
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		if err := viper.Unmarshal(a.options); err != nil {
			return err
		}
	}

	if a.options != nil {
		if err := a.applyOptionRules(); err != nil {
			return err
		}
	}

	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}

	return nil
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func printWorkingDir() {
	wd, _ := os.Getwd()
	logx.Infof("WorkingDir: %s", wd)
}

func (a *App) applyOptionRules() error {
	// if completeableOptions, ok := a.options.(CompleteableOptions); ok {
	// 	if err := completeableOptions.Complete(); err != nil {
	// 		return err
	// 	}
	// }

	if errs := a.options.Validate(); len(errs) != 0 {
		return errorx.NewAggregate(errs)
	}

	return nil
}

func addCmdTemplate(cmd *cobra.Command, namedFlagSets FlagSets) {
	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)

		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})
}
