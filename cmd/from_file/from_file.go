package fromfile

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FloRichardAloeCorp/autoprospect/internal/company"
	companyapiclient "github.com/FloRichardAloeCorp/autoprospect/internal/company_api_client"
	csvwriter "github.com/FloRichardAloeCorp/autoprospect/internal/csv_writer"
	"github.com/spf13/cobra"
)

var (
	path     string
	fromFile = &cobra.Command{
		Use:   "from_file",
		Short: "Retrieve companies data from a list of company names",
		RunE: func(cmd *cobra.Command, args []string) error {
			list, err := os.Open(path)
			if err != nil {
				return err
			}

			api := companyapiclient.New()

			companies := []company.Company{}

			scanner := bufio.NewScanner(list)
			for scanner.Scan() {
				params := companyapiclient.SearchRequestQueryParams{
					AdministrativeState: "A",
					Q:                   scanner.Text(),
				}

				res, err := api.Search(&params, 1)
				if err != nil {
					return err
				}

				if len(res.Results) == 0 {
					fmt.Println("No results for ", scanner.Text())
					continue
				}

				companies = append(companies, res.Results[0])
			}

			if err := csvwriter.Write("out.csv", companies); err != nil {
				return err
			}

			return nil
		},
	}
)

func Init(parent *cobra.Command) {
	fromFile.PersistentFlags().StringVarP(&path, "path", "p", "", "File that list companies")

	parent.AddCommand(fromFile)
}
