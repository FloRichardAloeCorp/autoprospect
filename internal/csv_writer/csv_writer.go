package csvwriter

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FloRichardAloeCorp/autoprospect/internal/company"
)

var CSVHeader = []string{"nom", "SIREN", "date de création", "Département", "CA", "Résultat Net", "Salariés", "Lien Pappers"}

func Write(to string, companies []company.Company) error {
	csvData := [][]string{
		CSVHeader,
	}

	for _, c := range companies {
		turnover := -1
		netResults := -1
		if _, ok := c.Finances["2023"]; ok {
			turnover = c.Finances["2023"].CA
			netResults = c.Finances["2023"].ResultatNet
		}

		csvData = append(csvData, []string{
			c.NomRaisonSociale,
			c.Siren,
			c.DateCreation.Format("01/02/2006"),
			c.Siege.Departement,
			strconv.Itoa(turnover),
			strconv.Itoa(netResults),
			c.TrancheEffectifSalarie.ToEmployeeRange(),
			fmt.Sprintf("https://www.pappers.fr/entreprise/%s-%s", strings.ReplaceAll(strings.ToLower(c.NomRaisonSociale), " ", "-"), c.Siren),
		})
	}

	f, err := os.Create("out.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	if err := w.WriteAll(csvData); err != nil {
		return err
	}

	return nil
}
