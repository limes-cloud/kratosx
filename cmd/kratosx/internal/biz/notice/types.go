package notice

type ListNoticeRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"pageSize"`
	Order      *string `json:"order"`
	OrderBy    *string `json:"orderBy"`
	Id         *uint32 `json:"id"`
	ClassifyId *uint32 `json:"classifyId"`
	Title      *string `json:"title"`
}

type ExportNoticeRequest struct {
	Id         *uint32 `json:"id"`
	ClassifyId *uint32 `json:"classifyId"`
	Title      *string `json:"title"`
}
