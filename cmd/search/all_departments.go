package search

import (
	"fmt"
	"strconv"

	"github.com/FloRichardAloeCorp/autoprospect/internal/company"
	companyapiclient "github.com/FloRichardAloeCorp/autoprospect/internal/company_api_client"
	csvwriter "github.com/FloRichardAloeCorp/autoprospect/internal/csv_writer"
	"github.com/spf13/cobra"
)

var allDepartments = &cobra.Command{
	Use:   "all_deps",
	Short: "Search in all departments of France",
	RunE: func(cmd *cobra.Command, args []string) error {
		companies := []company.Company{}
		api := companyapiclient.New()

		for i := 1; i < 96; i++ {
			if i == 20 {
				continue
			}

			department := fmt.Sprintf("%02s", strconv.Itoa(i))
			fmt.Println("Searching companies in", department)

			params := companyapiclient.SearchRequestQueryParams{
				AdministrativeState: "A",
				CompanyCategory:     category,
				MainActivity:        activity,
				Department:          department,
			}

			currentPage := 0
			pageCount := 1

			for currentPage < pageCount {
				currentPage++
				res, err := api.Search(&params, currentPage)
				if err != nil {
					return err
				}

				if currentPage == 1 {
					pageCount = res.TotalPages
				}

				for _, company := range res.Results {
					if company.Siege.Departement != department {
						continue
					}

					companies = append(companies, company)
				}
			}
		}

		if err := csvwriter.Write("out.csv", companies); err != nil {
			return err
		}
		return nil
	},
}

func initAllDepartmentsCmd(parent *cobra.Command) {
	parent.AddCommand(allDepartments)
}
