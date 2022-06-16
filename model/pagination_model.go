package model

import (
	"math"

	"github.com/krobus00/technical-test-rest-api/constant"
)

type PaginationRequest struct {
	Page   int64  `query:"page"`
	Limit  int64  `query:"limit"`
	Search string `query:"search"`
}

func (p *PaginationRequest) BuildRequest() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = constant.DEFAULT_PAGE_LIMIT
	}
	if p.Limit > constant.MAX_PAGE_LIMIT {
		p.Limit = constant.MAX_PAGE_LIMIT
	}
}

type PaginationResponse struct {
	TotalItems  int64       `json:"totalItems"`
	TotalPages  int64       `json:"totalPages"`
	CurrentPage int64       `json:"currentPage"`
	Limit       int64       `json:"limit"`
	Items       interface{} `json:"items"`
}

func (p *PaginationResponse) BuildResponse(payload *PaginationRequest, items interface{}, totalItems int64) {
	p.TotalItems = totalItems
	p.CurrentPage = payload.Page
	p.Limit = payload.Limit
	p.Items = items
	if p.TotalItems == 0 || p.Limit == 0 {
		p.TotalPages = 0
	} else {
		p.TotalPages = int64(math.Ceil(float64(p.TotalItems) / float64(p.Limit)))
	}
}
