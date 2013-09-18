package gotcdata

import (
	"encoding/xml"
)

type Coder struct {
	CoderId       string `xml:"coder_id"`
	Handle        string `xml:"handle"`
	CountryName   string `xml:"country_name"`
	AlgRating     string `xml:"alg_rating"`
	AlgVol        string `xml:"alg_vol"`
	AlgNumRatings string `xml:"alg_num_ratings"`
	DesRating     string `xml:"des_rating"`
	DesVol        string `xml:"des_vol"`
	DesNumRatings string `xml:"des_num_ratings"`
	MarRating     string `xml:"mar_rating"`
	MarVol        string `xml:"mar_vol"`
	MarNumRatings string `xml:"mar_num_ratings"`
	School        string `xml:"school"`
}

type CoderList struct {
	XMLName xml.Name `xml:"dd_coder_list"`
	Coders  []Coder  `xml:"row"`
}
