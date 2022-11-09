package request

type EmployeeCreate struct {
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseID  int    `json:"warehouse_id" binding:"required"`
}

type EmployeePatch struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	WarehouseID *int    `json:"warehouse_id"`
}
