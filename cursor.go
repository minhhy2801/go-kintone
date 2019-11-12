package kintone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Cursor struct {
	id       	uint64		`json:"id"`
	records   	[]*Record	`json:"records"`
	next	 	bool		`json:"next"`
	totalCount	int64		`json:"totalCount"`
}

type AddCursorObj struct {
	id       	uint64		`json:"id"`
	totalCount	int64		`json:"totalCount"`
}

type GetCursorObj struct {
	records   	[]*Record	`json:"records"`
	next	 	bool		`json:"next"`
}