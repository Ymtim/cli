package actions

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewCmdActions(f *cmdutil.Factory) *cobra.Command {
	cs := f.IOStreams.ColorScheme()

	cmd := &cobra.Command{
		Use:    "actions",
		Short:  "Learn about working with GitHub Actions",
		Long:   actionsExplainer(cs),
		Hidden: true,
	}

	cmdutil.DisableAuthCheck(cmd)

	return cmd
}

func actionsExplainer(cs *iostreams.ColorScheme) string {
	header := cs.Bold("Welcome to GitHub Actions on the command line.")
	runHeader := cs.Bold("Interacting with workflow runs")
	workflowHeader := cs.Bold("Interacting with workflow files")
	cacheHeader := cs.Bold("Interacting with the Actions cache")

	return heredoc.Docf(`
			%s

			GitHub CLI integrates with Actions to help you manage runs and workflows.

			%s  
			gh run list:      List recent workflow runs  
			gh run view:      View details for a workflow run or one of its jobs  
			gh run watch:     Watch a workflow run while it executes  
			gh run rerun:     Rerun a failed workflow run  
			gh run download:  Download artifacts generated by runs

			To see more help, run 'gh help run <subcommand>'

			%s  
			gh workflow list:     List all the workflow files in your repository  
			gh workflow view:     View details for a workflow file  
			gh workflow enable:   Enable a workflow file  
			gh workflow disable:  Disable a workflow file  
			gh workflow run:      Trigger a workflow_dispatch run for a workflow file

			To see more help, run 'gh help workflow <subcommand>'

			%s
			gh cache list:    List all the caches saved in Actions for a repository
			gh cache delete:  Delete one or all saved caches in Actions for a repository

			To see more help, run 'gh help cache <subcommand>'

		`, header, runHeader, workflowHeader, cacheHeader)
}
