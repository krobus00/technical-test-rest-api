package model

type DateColumn struct {
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

type BasicResponse struct {
	Message string `json:"message"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
}
