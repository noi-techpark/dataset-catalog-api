// © 2024 NOI Techpark <digital@noi.bz.it>
// SPDX-License-Identifier: AGPL-3.0-or-later

package transformers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var BaseURL string

func init() {
	// Load environment variables from .env if available.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "https://data-catalog.opendatahub.com/"
	}
	BaseURL = baseURL
}

var (
	ContactEmail       = "help@opendatahub.com"
	ContactPhoneNumber = "+390471066600"
	ContactWebsite     = "https://opendatahub.com"
	OrganizationName   = "Noi Spa"
	OrganizationURL    = "https://noi.bz.it"
	BrandSlogan        = "Develop digital solutions based on real data"
	PostalCode         = "39100"
	StreetAddress      = "Via Volta 13/A"
	AddressLocality    = "Bolzano"
	AddressRegion      = "Alto Adige"
	VatID              = "IT02595720216"
	TaxID              = "IT02595720216"
)

// Dataset represents the internal dataset structure.
type Dataset struct {
	ID             string              `json:"Id"`
	Self           string              `json:"Self"`
	Type           string              `json:"Type"`
	Meta           MetaData            `json:"_Meta"`
	ApiUrl         string              `json:"ApiUrl"`
	Output         interface{}         `json:"Output"`
	ApiType        string              `json:"ApiType"`
	BaseUrl        string              `json:"BaseUrl"`
	ODHTags        []interface{}       `json:"ODHTags"`
	OdhType        interface{}         `json:"OdhType"`
	Sources        interface{}         `json:"Sources"`
	Category       []string            `json:"Category"`
	ApiAccess      interface{}         `json:"ApiAccess"`
	ApiFilter      []string            `json:"ApiFilter"`
	Dataspace      string              `json:"Dataspace"`
	OdhTagIds      interface{}         `json:"OdhTagIds"`
	PathParam      []string            `json:"PathParam"`
	Shortname      string              `json:"Shortname"`
	Deprecated     bool                `json:"Deprecated"`
	LastChange     string              `json:"LastChange"`
	SwaggerUrl     string              `json:"SwaggerUrl"`
	FirstImport    string              `json:"FirstImport"`
	LicenseInfo    LicenseInfo         `json:"LicenseInfo"`
	PublishedOn    []interface{}       `json:"PublishedOn"`
	RecordCount    interface{}         `json:"RecordCount"`
	DataProvider   []string            `json:"DataProvider"`
	ImageGallery   []ImageGalleryItem  `json:"ImageGallery"`
	ApiDescription map[string]string   `json:"ApiDescription"`
}

// MetaData represents metadata information.
type MetaData struct {
	ID         string `json:"Id"`
	Type       string `json:"Type"`
	Source     string `json:"Source"`
	Reduced    bool   `json:"Reduced"`
	LastUpdate string `json:"LastUpdate"`
	UpdateInfo struct {
		UpdatedBy    string `json:"UpdatedBy"`
		UpdateSource string `json:"UpdateSource"`
	} `json:"UpdateInfo"`
}

// LicenseInfo represents licensing information.
type LicenseInfo struct {
	Author        string `json:"Author"`
	License       string `json:"License"`
	ClosedData    bool   `json:"ClosedData"`
	LicenseHolder string `json:"LicenseHolder"`
}

// ImageGalleryItem represents an image item.
type ImageGalleryItem struct {
	Width         interface{}            `json:"Width"`
	Height        interface{}            `json:"Height"`
	License       string                 `json:"License"`
	ValidTo       interface{}            `json:"ValidTo"`
	ImageUrl      string                 `json:"ImageUrl"`
	CopyRight     interface{}            `json:"CopyRight"`
	ImageDesc     map[string]interface{} `json:"ImageDesc"`
	ImageName     interface{}            `json:"ImageName"`
	ImageTags     interface{}            `json:"ImageTags"`
	ValidFrom     interface{}            `json:"ValidFrom"`
	ImageTitle    map[string]interface{} `json:"ImageTitle"`
	ImageSource   string                 `json:"ImageSource"`
	IsInGallery   interface{}            `json:"IsInGallery"`
	ImageAltText  map[string]interface{} `json:"ImageAltText"`
	ListPosition  interface{}            `json:"ListPosition"`
	LicenseHolder interface{}            `json:"LicenseHolder"`
}

