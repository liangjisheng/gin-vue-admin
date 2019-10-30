package modelinterface

// PageInfo ...
type PageInfo struct {
	Page     int
	PageSize int
}

// Paging 分页接口
type Paging interface {
	GetInfoList(PageInfo) (list interface{}, total int, err error)
}
