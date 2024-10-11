package companyapiclient

import (
	"net/url"
	"strconv"
)

type SearchRequestQueryParams struct {
	Q                         string
	IsAssociation             bool
	IsBio                     bool
	IsIndividualEntrepreneur  bool
	IsShowEntrepreneur        bool
	IsTerritorialCollectiviry bool

	AdministrativeState string
	CompanyCategory     string

	MainActivity string

	Department string
	ZipCode    string
}

func (q *SearchRequestQueryParams) build() url.Values {
	values := url.Values{}

	if q.Q != "" {
		values.Add("q", q.Q)
	}

	values.Add("est_association", strconv.FormatBool(q.IsAssociation))
	values.Add("est_bio", strconv.FormatBool(q.IsBio))
	values.Add("est_entrepreneur_individuel", strconv.FormatBool(q.IsIndividualEntrepreneur))
	values.Add("est_entrepreneur_spectacle", strconv.FormatBool(q.IsShowEntrepreneur))
	values.Add("est_collectivite_territoriale", strconv.FormatBool(q.IsTerritorialCollectiviry))

	values.Add("etat_administratif", q.AdministrativeState)

	if q.CompanyCategory != "" {
		values.Add("categorie_entreprise", q.CompanyCategory)
	}

	if q.MainActivity != "" {
		values.Add("activite_principale", q.MainActivity)
	}

	if q.Department != "" {
		values.Add("departement", q.Department)
	}

	if q.ZipCode != "" {
		values.Add("code_postal", q.ZipCode)
	}

	return values
}
