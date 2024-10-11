package search

import (
	"fmt"
	"time"

	"github.com/FloRichardAloeCorp/autoprospect/internal/company"
	companyapiclient "github.com/FloRichardAloeCorp/autoprospect/internal/company_api_client"
	csvwriter "github.com/FloRichardAloeCorp/autoprospect/internal/csv_writer"
	"github.com/spf13/cobra"
)

var (
	category   string
	activity   string
	department string
	from       string
	zipCode    string

	search = &cobra.Command{
		Use:   "search",
		Short: "Search companies with diverse parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			api := companyapiclient.New()
			dateLimit, err := time.Parse("01/02/2006", from)
			if err != nil {
				dateLimit = time.Time{}
			}

			fmt.Println(category, activity, department)

			params := companyapiclient.SearchRequestQueryParams{
				AdministrativeState: "A",
				CompanyCategory:     category,
				MainActivity:        activity,
				Department:          department,
			}

			companies := []company.Company{}

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
					if company.DateCreation.Before(dateLimit) {
						continue
					}

					if department != "" && company.Siege.Departement != department {
						continue
					}

					companies = append(companies, company)
				}

				fmt.Printf("%d/%d\n", res.Page, res.TotalPages)
			}

			if err := csvwriter.Write("out.csv", companies); err != nil {
				return err
			}

			return nil
		},
	}
)

func Init(parent *cobra.Command) {
	search.PersistentFlags().StringVarP(&category, "category", "c", "", "PME, ETI or GE. Can be combined by separating values with a coma.")
	search.PersistentFlags().StringVarP(&activity, "activity", "a", "", "NAF codes splitted by coma.")
	search.PersistentFlags().StringVarP(&from, "from", "f", "", "Minimum creation date of the company in dd/mm/yyyy format")

	search.Flags().StringVarP(&department, "department", "d", "", "Departement number.")
	search.Flags().StringVarP(&zipCode, "zip_code", "z", "", "Zip code of the company")

	initAllDepartmentsCmd(search)

	parent.AddCommand(search)
}
