package models

type UzivatelInfo struct {
	Jmeno          string `json:"JMENO"`
	UzivJmeno      string `json:"UZIV_JMENO"`
	DatabaseID     string `json:"DATABASE_ID"`
	OrganizaceID   string `json:"ORGANIZACE_ID"`
	OsobaID        string `json:"OSOBA_ID"`
	KategorieIDCSV string `json:"KATEGORIE_ID_CSV"`
	TridaID        string `json:"TRIDA_ID"`
	TridaNazev     string `json:"TRIDA_NAZEV"`
}

type RozvrhovaUdalost struct {
	UdalostID               string             `json:"UDALOST_ID"`
	Datum                   string             `json:"DATUM"`
	Poradi                  int                `json:"PORADI"`
	ObdobiDneOdID           string             `json:"OBDOBI_DNE_OD_ID"`
	ObdobiDneDoID           string             `json:"OBDOBI_DNE_DO_ID"`
	ObdobiDneOdNazev        string             `json:"OBDOBI_DNE_OD_NAZEV"`
	ObdobiDneDoNazev        string             `json:"OBDOBI_DNE_DO_NAZEV"`
	CasOd                   string             `json:"CAS_OD"`
	CasDo                   string             `json:"CAS_DO"`
	DelkaPocetHodin         int                `json:"DELKA_POCET_HODIN"`
	Nazev                   string             `json:"NAZEV"`
	Popis                   string             `json:"POPIS"`
	TypUdalosti             *TypUdalosti       `json:"TYP_UDALOSTI"`
	DruhUdalosti            *DruhUdalosti      `json:"DRUH_UDALOSTI"`
	Cyklus                  string             `json:"CYKLUS"`
	Predmet                 *UdalostPredmet    `json:"PREDMET"`
	Barva                   string             `json:"BARVA"`
	BarvaPisma              string             `json:"BARVA_PISMA"`
	PovolenZapisDochazky    bool               `json:"POVOLEN_ZAPIS_DOCHAZKY"`
	PovolenZapisHodnoceni   bool               `json:"POVOLEN_ZAPIS_HODNOCENI"`
	SkupinyUdalosti         []SkupinaUdalosti  `json:"SKUPINY_UDALOSTI"`
	TridyUdalosti           interface{}        `json:"TRIDY_UDALOSTI"`
	MistnostiUdalosti       []MistnostUdalosti `json:"MISTNOSTI_UDALOSTI"`
	UciteleUdalosti         []UcitelUdalosti   `json:"UCITELE_UDALOSTI"`
	Poznamka                string             `json:"POZNAMKA"`
	ProbraneUcivo           string             `json:"PROBRANE_UCIVO"`
	NahrazujeHodiny         bool               `json:"NAHRAZUJE_HODINY"`
	JeSuplovanaHodinami     bool               `json:"JE_SUPLOVANA_HODINAMI"`
	NahrazujeHodinyText     string             `json:"NAHRAZUJE_HODINY_TEXT"`
	JeSuplovanaHodinamiText string             `json:"JE_SUPLOVANA_HODINAMI_TEXT"`
	PocetOducenychHodin     int                `json:"POCET_ODUCENYCH_HODIN"`
}

type TypUdalosti struct {
	TypUdalostiID string `json:"TYP_UDALOSTI_ID"`
	Popis         string `json:"POPIS"`
}

type DruhUdalosti struct {
	TypUdalostiID  string `json:"TYP_UDALOSTI_ID"`
	DruhUdalostiID string `json:"DRUH_UDALOSTI_ID"`
	Nazev          string `json:"NAZEV"`
	Popis          string `json:"POPIS"`
}

type UdalostPredmet struct {
	SkolniRokID     string `json:"SKOLNI_ROK_ID"`
	PredmetID       string `json:"PREDMET_ID"`
	Zkratka         string `json:"ZKRATKA"`
	Nazev           string `json:"NAZEV"`
	PoradiZobrazeni int    `json:"PORADI_ZOBRAZENI"`
}

type SkupinaUdalosti struct {
	SkupinaID          string `json:"SKUPINA_ID"`
	SkupinaNazev       string `json:"SKUPINA_NAZEV"`
	PriznakDruhSkupiny string `json:"PRIZNAK_DRUH_SKUPINY"`
	TridaID            string `json:"TRIDA_ID"`
	TridaNazev         string `json:"TRIDA_NAZEV"`
	PriznakAbsence     bool   `json:"PRIZNAK_ABSENCE"`
}

type MistnostUdalosti struct {
	MistnostID     string `json:"MISTNOST_ID"`
	Nazev          string `json:"NAZEV"`
	Popis          string `json:"POPIS"`
	PriznakAbsence bool   `json:"PRIZNAK_ABSENCE"`
}

type UcitelUdalosti struct {
	UcitelID       string `json:"UCITEL_ID"`
	Prijmeni       string `json:"PRIJMENI"`
	Jmeno          string `json:"JMENO"`
	Zkratka        string `json:"ZKRATKA"`
	PriznakAbsence bool   `json:"PRIZNAK_ABSENCE"`
}
