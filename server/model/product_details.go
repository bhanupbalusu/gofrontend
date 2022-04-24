package model

type BulkQuantity struct {
	Volume  string `json:"volume,omitempty"`
	BQUnits string `json:"bq_units,omitempty"`
}

type Price struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	PerUnit  string `json:"per_unit,omitempty"`
	PUnits   string `json:"p_units,omitempty"`
}

type ProductDetails struct {
	ProductName  string       `json:"product_name,omitempty"`
	Description  string       `json:"description,omitempty"`
	ImageURL     string       `json:"image_url,omitempty"`
	BulkQuantity BulkQuantity `json:"bulk_quantity,omitempty"`
	Price        Price        `json:"price,omitempty"`
}

type Schedular struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty" bson:"end_date,omitempty"`
}

type ProductDetailsModel struct {
	ProductId          string           `json:"product_id,omitempty"`
	PreOrderRequestId  string           `json:"pre_order_request_id,omitempty"`
	CustomerId         string           `json:"customer_id,omitempty"`
	ProductDetailsList []ProductDetails `json:"product_details_list,omitempty"`
	Schedular          Schedular        `json:"schedular,omitempty"`
}
