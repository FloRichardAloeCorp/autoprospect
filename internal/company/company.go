package company

import (
	"time"
)

type Company struct {
	Siren                       string                 `json:"siren"`
	NomComplet                  string                 `json:"nom_complet"`
	NomRaisonSociale            string                 `json:"nom_raison_sociale"`
	Sigle                       *string                `json:"sigle"`
	NombreEtablissements        int                    `json:"nombre_etablissements"`
	NombreEtablissementsOuverts int                    `json:"nombre_etablissements_ouverts"`
	Siege                       Siege                  `json:"siege"`
	DateCreation                Date                   `json:"date_creation"`
	DateFermeture               string                 `json:"date_fermeture"`
	TrancheEffectifSalarie      TrancheEffectifSalarie `json:"tranche_effectif_salarie"`
	AnneeTrancheEffectifSalarie string                 `json:"annee_tranche_effectif_salarie"`
	DateMiseAJour               string                 `json:"date_mise_a_jour"`
	CategorieEntreprise         string                 `json:"categorie_entreprise"`
	CaractereEmployeur          string                 `json:"caractere_employeur"`
	AnneeCategorieEntreprise    string                 `json:"annee_categorie_entreprise"`
	EtatAdministratif           string                 `json:"etat_administratif"`
	NatureJuridique             string                 `json:"nature_juridique"`
	ActivitePrincipale          string                 `json:"activite_principale"`
	SectionActivitePrincipale   string                 `json:"section_activite_principale"`
	StatutDiffusion             string                 `json:"statut_diffusion"`
	MatchingEtablissements      []Etablissement        `json:"matching_etablissements"`
	Dirigeants                  []Dirigeant            `json:"dirigeants"`
	Finances                    map[string]Finance     `json:"finances"`
	Complements                 Complements            `json:"complements"`
}

// Structure représentant le siège social
type Siege struct {
	ActivitePrincipale               string    `json:"activite_principale"`
	ActivitePrincipaleRegistreMetier *string   `json:"activite_principale_registre_metier"`
	AnneeTrancheEffectifSalarie      string    `json:"annee_tranche_effectif_salarie"`
	Adresse                          string    `json:"adresse"`
	CaractereEmployeur               string    `json:"caractere_employeur"`
	Cedex                            *string   `json:"cedex"`
	CodePaysEtranger                 *string   `json:"code_pays_etranger"`
	CodePostal                       string    `json:"code_postal"`
	Commune                          string    `json:"commune"`
	ComplementAdresse                string    `json:"complement_adresse"`
	DateCreation                     string    `json:"date_creation"`
	DateFermeture                    string    `json:"date_fermeture"`
	DateDebutActivite                string    `json:"date_debut_activite"`
	DateMiseAJour                    time.Time `json:"date_mise_a_jour"`
	Departement                      string    `json:"departement"`
	DistributionSpeciale             *string   `json:"distribution_speciale"`
	EstSiege                         bool      `json:"est_siege"`
	EtatAdministratif                string    `json:"etat_administratif"`
	GeoID                            string    `json:"geo_id"`
	IndiceRepetition                 *string   `json:"indice_repetition"`
	Latitude                         string    `json:"latitude"`
	LibelleCedex                     *string   `json:"libelle_cedex"`
	LibelleCommune                   string    `json:"libelle_commune"`
	LibelleCommuneEtranger           *string   `json:"libelle_commune_etranger"`
	LibellePaysEtranger              *string   `json:"libelle_pays_etranger"`
	LibelleVoie                      string    `json:"libelle_voie"`
	ListeEnseignes                   []string  `json:"liste_enseignes"`
	ListeFiness                      []string  `json:"liste_finess"`
	ListeIDCC                        []string  `json:"liste_idcc"`
	ListeIDBio                       []string  `json:"liste_id_bio"`
	ListeRGE                         []string  `json:"liste_rge"`
	ListeUAI                         []string  `json:"liste_uai"`
	Longitude                        string    `json:"longitude"`
	NomCommercial                    *string   `json:"nom_commercial"`
	NumeroVoie                       string    `json:"numero_voie"`
	Region                           string    `json:"region"`
	EPCI                             string    `json:"epci"`
	Siret                            string    `json:"siret"`
	StatutDiffusionEtablissement     string    `json:"statut_diffusion_etablissement"`
	TrancheEffectifSalarie           string    `json:"tranche_effectif_salarie"`
	TypeVoie                         string    `json:"type_voie"`
}

// Structure représentant un établissement associé
type Etablissement struct {
	ActivitePrincipale           string   `json:"activite_principale"`
	Adresse                      string   `json:"adresse"`
	AnneeTrancheEffectifSalarie  string   `json:"annee_tranche_effectif_salarie"`
	AncienSiege                  bool     `json:"ancien_siege"`
	CaractereEmployeur           string   `json:"caractere_employeur"`
	CodePostal                   string   `json:"code_postal"`
	Commune                      string   `json:"commune"`
	DateCreation                 string   `json:"date_creation"`
	DateDebutActivite            string   `json:"date_debut_activite"`
	DateFermeture                string   `json:"date_fermeture"`
	EPCI                         string   `json:"epci"`
	EstSiege                     bool     `json:"est_siege"`
	EtatAdministratif            string   `json:"etat_administratif"`
	GeoID                        string   `json:"geo_id"`
	Latitude                     string   `json:"latitude"`
	LibelleCommune               string   `json:"libelle_commune"`
	ListeEnseignes               []string `json:"liste_enseignes"`
	ListeFiness                  []string `json:"liste_finess"`
	ListeIDCC                    []string `json:"liste_idcc"`
	ListeIDOrganismeFormation    []string `json:"liste_id_organisme_formation"`
	ListeIDBio                   []string `json:"liste_id_bio"`
	ListeRGE                     []string `json:"liste_rge"`
	ListeUAI                     []string `json:"liste_uai"`
	Longitude                    string   `json:"longitude"`
	NomCommercial                *string  `json:"nom_commercial"`
	Region                       string   `json:"region"`
	Siret                        string   `json:"siret"`
	StatutDiffusionEtablissement string   `json:"statut_diffusion_etablissement"`
	TrancheEffectifSalarie       string   `json:"tranche_effectif_salarie"`
}

// Structure représentant un dirigeant
type Dirigeant struct {
	Nom              string `json:"nom"`
	Prenoms          string `json:"prenoms"`
	AnneeDeNaissance string `json:"annee_de_naissance"`
	DateDeNaissance  string `json:"date_de_naissance"`
	Qualite          string `json:"qualite"`
	Nationalite      string `json:"nationalite"`
	TypeDirigeant    string `json:"type_dirigeant"`
}

// Structure représentant les informations financières
type Finance struct {
	CA          int `json:"ca"`
	ResultatNet int `json:"resultat_net"`
}

// Structure des compléments d'information
type Complements struct {
	CollectiviteTerritoriale       CollectiviteTerritoriale `json:"collectivite_territoriale"`
	ConventionCollectiveRenseignee bool                     `json:"convention_collective_renseignee"`
	ListeIDCC                      []string                 `json:"liste_idcc"`
	EgaproRenseignee               bool                     `json:"egapro_renseignee"`
	EstAssociation                 bool                     `json:"est_association"`
	EstBio                         bool                     `json:"est_bio"`
	EstEntrepreneurIndividuel      bool                     `json:"est_entrepreneur_individuel"`
	EstEntrepreneurSpectacle       bool                     `json:"est_entrepreneur_spectacle"`
	EstESS                         bool                     `json:"est_ess"`
	EstFiness                      bool                     `json:"est_finess"`
	EstOrganismeFormation          bool                     `json:"est_organisme_formation"`
	EstQualiopi                    bool                     `json:"est_qualiopi"`
	ListeIDOrganismeFormation      []string                 `json:"liste_id_organisme_formation"`
	EstRGE                         bool                     `json:"est_rge"`
	EstSIAE                        bool                     `json:"est_siae"`
	EstServicePublic               bool                     `json:"est_service_public"`
	EstSocieteMission              bool                     `json:"est_societe_mission"`
	EstUAI                         bool                     `json:"est_uai"`
	IdentifiantAssociation         *string                  `json:"identifiant_association"`
	StatutBio                      bool                     `json:"statut_bio"`
	StatutEntrepreneurSpectacle    string                   `json:"statut_entrepreneur_spectacle"`
	TypeSIAE                       string                   `json:"type_siae"`
}

// Structure représentant les collectivités territoriales
type CollectiviteTerritoriale struct {
	CodeINSEE string `json:"code_insee"`
	Code      string `json:"code"`
	Niveau    string `json:"niveau"`
	Elus      []Elu  `json:"elus"`
}

// Structure représentant un élu
type Elu struct {
	Nom              string `json:"nom"`
	Prenoms          string `json:"prenoms"`
	AnneeDeNaissance string `json:"annee_de_naissance"`
	Fonction         string `json:"fonction"`
	Sexe             string `json:"sexe"`
}

type TrancheEffectifSalarie string

func (t *TrancheEffectifSalarie) ToEmployeeRange() string {
	switch string(*t) {
	case "00":
		return "0"
	case "01":
		return "1 ou 2"
	case "02":
		return "3 à 5"
	case "03":
		return "6 à 9"
	case "11":
		return "10 à 19"
	case "12":
		return "20 à 49"
	case "21":
		return "50 à 99"
	case "22":
		return "100 à 199"
	case "31":
		return "200 à 249"
	case "32":
		return "250 à 499"
	case "41":
		return "500 à 999"
	case "42":
		return "1 000 à 1 999"
	case "51":
		return "2 000 à 4 999"
	case "52":
		return "5 000 à 9 999"
	case "53":
		return "10 000+"
	default:
		return "Inconnu"
	}
}
