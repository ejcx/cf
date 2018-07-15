package cmd

import "github.com/spf13/cobra"

func init() {
	var ListZones = &cobra.Command{
		Use:   "list-zones",
		Short: "Command for listing zones",
		Long: `This is a meaty description of the list-zones
`,
	}

	var zone = &cobra.Command{
		Use:   "zone",
		Short: "Commands for interacting zones",
		Long: `This is a meaty description of the zone api.
`,
	}
	zone.AddCommand(ListZones)
	RootCmd.AddCommand(zone)

}
