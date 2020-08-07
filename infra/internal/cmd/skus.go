package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/MemeLabs/go-ppspp/infra/pkg/node"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(skusCmd)
}

var skusCmd = &cobra.Command{
	Use:               "skus [provider]",
	Short:             "List provider SKUs",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: providerValidArgsFunc,
	RunE: func(cmd *cobra.Command, args []string) error {
		header := []string{
			"Name",
			"CPUs (cores)",
			"Memory (MB)",
			"Network Cap (GB)",
			"Network Speed (Mbps)",
			"Price Monthly",
			"Price Hourly",
		}
		alignment := []int{
			tablewriter.ALIGN_DEFAULT,
			tablewriter.ALIGN_DEFAULT,
			tablewriter.ALIGN_DEFAULT,
			tablewriter.ALIGN_DEFAULT,
			tablewriter.ALIGN_DEFAULT,
			tablewriter.ALIGN_RIGHT,
			tablewriter.ALIGN_RIGHT,
		}
		var data [][]string

		if len(args) == 1 {
			provider := args[0]
			driver, ok := backend.NodeDrivers[provider]
			if !ok {
				return fmt.Errorf("Unsupported provider: %s", provider)
			}

			rows, err := formatProviderSKUs(cmd.Context(), driver)
			if err != nil {
				return err
			}
			data = rows
		} else {
			header = append([]string{"Provider"}, header...)
			alignment = append([]int{tablewriter.ALIGN_DEFAULT}, alignment...)
			for _, driver := range backend.NodeDrivers {
				rows, err := formatProviderSKUs(cmd.Context(), driver)
				if err != nil {
					return err
				}
				data = append(data, prependProviderColumn(rows, driver)...)
			}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(header)
		table.SetColumnAlignment(alignment)
		table.AppendBulk(data)
		table.Render()

		return nil
	},
}

func formatProviderSKUs(ctx context.Context, driver node.Driver) ([][]string, error) {
	skus, err := driver.SKUs(ctx, &node.SKUsRequest{})
	if err != nil {
		return nil, fmt.Errorf("Loading SKUs failed: %w", err)
	}

	rows := [][]string{}
	for _, r := range skus {
		var networkCap, priceHourly, priceMonthly string
		if r.NetworkCap != 0 {
			networkCap = strconv.Itoa(r.NetworkCap)
		}
		if r.PriceHourly != nil {
			priceHourly = fmt.Sprintf("%.4f %s", r.PriceHourly.Value, r.PriceHourly.Currency)
		}
		if r.PriceMonthly != nil {
			priceMonthly = fmt.Sprintf("%.4f %s", r.PriceMonthly.Value, r.PriceMonthly.Currency)
		}

		rows = append(rows, []string{
			r.Name,
			strconv.Itoa(r.CPUs),
			strconv.Itoa(r.Memory),
			networkCap,
			strconv.Itoa(r.NetworkSpeed),
			priceMonthly,
			priceHourly,
		})
	}
	return rows, nil
}
