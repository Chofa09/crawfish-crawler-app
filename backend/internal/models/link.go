package models

type LINK struct {
	ID					int			`json:"id"`
	URLID 			int 		`json:"url_id"`
	Href 				string 	`json:"href"`
	IsInternal 	bool 		`json:"is_internal"`
	StatusCode 	int 		`json:"status_code"`
	IsBroken 		bool		`json:"is_broken"`
}

