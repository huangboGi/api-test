package api

// 用户端接口
const (
	UserSubscribe    = "/api/v1/wth/order/subscribe"
	UserRedeem       = "/api/v1/wth/order/redeem"
	UserOrderList    = "/api/v1/wth/order/page"
	UserOrderDetail  = "/api/v1/wth/order/detail"
	UserProductPage  = "/api/v1/wth/product/page"
	UserHoldPosition = "/api/v1/wth/order/holdPosition"
	UserInterestPage = "/api/v1/wth/order/interestPage"
	UserHis          = "/api/v1/wth/order/his"
	UserPeriodDetail = "/api/v1/wth/order/periodDetail"
	UserOpenSub      = "/api/v1/wth/order/openSub"
)

// 管理端-币种配置接口
const (
	AdminCoinAdd        = "/api/v1/admin/wth/coin/add"
	AdminCoinUpdate     = "/api/v1/admin/wth/coin/update"
	AdminCoinShelves    = "/api/v1/admin/wth/coin/updateShelves"
	AdminCoinPage       = "/api/v1/admin/wth/coin/page"
	AdminCoinDetail     = "/api/v1/admin/wth/coin/detail"
	AdminCoinSelectCoin = "/api/v1/admin/wth/coin/selectCoin"
)

// 管理端-规格管理接口
const (
	AdminSpecAdd     = "/api/v1/admin/wth/spec/add"
	AdminSpecUpdate  = "/api/v1/admin/wth/spec/update"
	AdminSpecShelves = "/api/v1/admin/wth/spec/updateShelves"
	AdminSpecPage    = "/api/v1/admin/wth/spec/page"
	AdminSpecList    = "/api/v1/admin/wth/spec/list"
	AdminSpecDetail  = "/api/v1/admin/wth/spec/detail"
)

// 管理端-产品管理接口
const (
	AdminProductAdd     = "/api/v1/admin/wth/product/add"
	AdminProductShelves = "/api/v1/admin/wth/product/updateShelves"
	AdminProductUpdate  = "/api/v1/admin/wth/product/update"
	AdminProductPage    = "/api/v1/admin/wth/product/page"
	AdminProductDetail  = "/api/v1/admin/wth/product/detail"
)

// 用户端-产品管理接口
const (
	AppProductList   = "/api/v1/wth/product/list"
	AppProductDetail = "/api/v1/wth/product/detail"
)
