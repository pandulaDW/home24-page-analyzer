package models

// HeadingCount defines the count of heading tags
type HeadingCount struct {
	H1Count int `json:"h1_count"`
	H2Count int `json:"h2_count"`
	H3Count int `json:"h3_count"`
	H4Count int `json:"h4_count"`
	H5Count int `json:"h5_count"`
	H6Count int `json:"h6_count"`
}

// HTMLPageDetails holds details extracted from a particular html page
type HTMLPageDetails struct {
	Title        string `json:"title"`
	HeadingCount `json:"heading_count"`
}
