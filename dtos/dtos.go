package dtos

import "time"

type Transaction struct {
	Id   string `json:"id"`
	Span Span   `json:"span"`
}

type Span struct {
	Parent Record `json:"parent"`
}

type Record struct {
	Name      string     `json:"name"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Child     []*Record  `json:"child"`
	IsParent  bool       `json:"-"`
}
