package models

// HTMLVersion is the html version of the page
type HTMLVersion string

const (
	HTML5   = "HTML5"
	HTML4   = "HTML4"
	XHTML   = "XHTML"
	UNKNOWN = "UNKNOWN"
)

// HeadingCount returns the count of each heading tag
type HeadingCount struct {
	H1Count int `json:"h1_count"`
	H2Count int `json:"h2_count"`
	H3Count int `json:"h3_count"`
	H4Count int `json:"h4_count"`
	H5Count int `json:"h5_count"`
	H6Count int `json:"h6_count"`
}

// LinkCount returns the details about the internal and external links in the page
type LinkCount struct {
	InternalLinkCount     int `json:"internal_link_count"`
	ExternalLinkCount     int `json:"external_link_count"`
	InaccessibleLinkCount int `json:"inaccessible_link_count"`
}

// HTMLPageDetails holds details extracted from a particular html page
type HTMLPageDetails struct {
	HTMLVersion  HTMLVersion `json:"html_version"`
	Title        string      `json:"title"`
	HeadingCount `json:"heading_count"`
	LinkCount    `json:"link_count"`
	IsLoginForm  bool `json:"is_login_form"`
}
