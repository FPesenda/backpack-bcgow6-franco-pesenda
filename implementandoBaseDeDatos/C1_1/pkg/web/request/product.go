package request

type Product struct {
	Description    string  `json:"description" binding:"required"`
	ExpirationRate int     `json:"expiration_rate" binding:"required"`
	FreezingRate   int     `json:"freezing_rate" binding:"required"`
	Height         float32 `json:"height" binding:"required"`
	Length         float32 `json:"length" binding:"required"`
	Netweight      float32 `json:"netweight" binding:"required"`
	ProductCode    string  `json:"product_code" binding:"required"`
	RecomFreezTemp float32 `json:"recommended_freezing_temperature" binding:"required"`
	Width          float32 `json:"width" binding:"required"`
	ProductTypeID  int     `json:"product_type_id" binding:"required"`
	SellerID       int     `json:"seller_id"`
}

type ProductUpdate struct {
	Description    string  `json:"description" gorm:"size:250"`
	ExpirationRate int     `json:"expiration_rate"`
	FreezingRate   int     `json:"freezing_rate"`
	Height         float32 `json:"height"`
	Length         float32 `json:"length"`
	Netweight      float32 `json:"netweight"`
	ProductCode    string  `json:"product_code" gorm:"size:20"`
	RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
	Width          float32 `json:"width"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       int     `json:"seller_id"`
}
