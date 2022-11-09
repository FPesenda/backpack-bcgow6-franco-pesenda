package request

type SectionPost struct {
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseID        int `json:"warehouse_id" binding:"required"`
	ProductTypeID      int `json:"product_type_id" binding:"required"`
}

type SectionPatch struct {
	SectionNumber      *int `json:"section_number" `
	CurrentTemperature *int `json:"current_temperature" `
	MinimumTemperature *int `json:"minimum_temperature" `
	CurrentCapacity    *int `json:"current_capacity" `
	MinimumCapacity    *int `json:"minimum_capacity" `
	MaximumCapacity    *int `json:"maximum_capacity" `
	WarehouseID        *int `json:"warehouse_id" `
	ProductTypeID      *int `json:"product_type_id" `
}
