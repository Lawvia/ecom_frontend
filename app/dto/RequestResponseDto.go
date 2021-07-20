package dto

type UserPanelDto struct {
	Token		 string           `json:"token" db:"token"`
	Username     string          `json:"username" db:"username"`
	Password     string          `json:"password" db:"password"`
}

type LoginPayload struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type AuthResponse struct {
	Error int `json:"error" db:"error"`
	Message string `json:"message" db:"message"`
	Status string `json:"status" db:"status"`
	Token string `json:"token" db:"token"`
}

type ItemsResponse struct {
	Error int `json:"error" db:"error"`
	Message string `json:"message" db:"message"`
	Status string `json:"status" db:"status"`
	Data	[]ItemsDetail `json:"data" db:"data"`
}

type ItemsDetail struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price string `json:"price" db:"price"`
	Stock int `json:"stock" db:"stock"`
}

type UpdateCart struct{
	Cart []DetailsCart `json:"cart"`
	Address string `json:"address"`
}

type DetailsCart struct {
	IdItems	 int  `json:"id_items"`
	Qty      int   `json:"qty"`
}

type UpdateResponse struct {
	Error int `json:"error" db:"error"`
	Message string `json:"message" db:"message"`
	Status string `json:"status" db:"status"`
	GrandTotal int `json:"grand_total" db:"grand_total"`
	Data	[]UpdateDetails `json:"data" db:"data"`
}

type UpdateDetails struct {
	Id int `json:"id" db:"id"`
	Name string `json:"item_name" db:"item_name"`
	IdItems string `json:"id_items" db:"id_items"`
	Price string `json:"price" db:"price"`
	Qty int `json:"qty" db:"qty"`
}

type ListTransac struct {
	Error int `json:"error" db:"error"`
	Message string `json:"message" db:"message"`
	Status string `json:"status" db:"status"`
	Data	[]ListDetails `json:"data" db:"data"`
}

type ListDetails struct {
	PurchaseKey string `json:"purchase_key" db:"purchase_key"`
	Address string `json:"address" db:"address"`
	IdItems string `json:"id_items" db:"id_items"`
	Total string `json:"total" db:"total"`
	Status int `json:"status" db:"status"`
	Detail 	[]DetailItem `json:"detail" db:"detail"`
}

type DetailItem struct {
	Id int `json:"item_id" db:"item_id"`
	Name string `json:"item_name" db:"item_name"`
	Qty int `json:"qty" db:"qty"`
}