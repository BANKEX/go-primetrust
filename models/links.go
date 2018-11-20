package models

type Links struct {
	Self    string `json:"self,omitempty"`
	First   string `json:"first,omitempty"`
	Related string `json:"related,omitempty"`
}
