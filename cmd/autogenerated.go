package cmd

import "github.com/spf13/cobra"

var (
	ZoneNameFilter string
)

func init() {
	var ListZones = &cobra.Command{
		Use:   "list-zones",
		Short: "Command for listing zones",
		Long:  `  This is a meaty description of the list-zones`,
		Run: func(cmd *cobra.Command, args []string) {
			Main(cmd, args, "ListZones")
		},
	}

	ListZones.Flags().StringVar(&ZoneNameFilter, "zone-name-filter", "", "string for filtering by name")

	var Zone = &cobra.Command{
		Use:   "zone",
		Short: "Commands for interacting with zones",
		Long:  `  This is a meaty description of the zone api.`,
	}
	Zone.AddCommand(ListZones)
	RootCmd.AddCommand(Zone)

}
