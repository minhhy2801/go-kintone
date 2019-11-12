package kintone

type Cursor struct {
	Id       	uint64		`json:"id"`
	Records   	[]*Record	`json:"records"`
	Next	 	bool		`json:"next"`
	TotalCount	int64		`json:"totalCount"`
}

type AddCursorObj struct {
	Id       	uint64		`json:"id"`
	TotalCount	int64		`json:"totalCount"`
}

type GetCursorObj struct {
	Records   	[]*Record	`json:"records"`
	Next	 	bool		`json:"next"`
}