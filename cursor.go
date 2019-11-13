package kintone

type Cursor struct {
	Id       	string		`json:"id"`
	Records   	[]*Record	`json:"records"`
	Next	 	bool		`json:"next"`
	TotalCount	int64		`json:"totalCount"`
}

type AddCursorObj struct {
	Id       	string		`json:"id"`
	TotalCount	int64		`json:"totalCount"`
}

type GetCursorObj struct {
	Records   	[]*Record	`json:"records"`
	Next	 	bool		`json:"next"`
}