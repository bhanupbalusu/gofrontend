package model

type PQ struct {
	Vol     string `json:"vol,omitempty"`
	PQUnits string `json:"pq_units,omitempty"`
}

type PPrice struct {
	PAmount   string `json:"p_amount,omitempty"`
	PCurrency string `json:"p_currency,omitempty"`
	PPerUnit  string `json:"p_per_unit,omitempty"`
	PPUnits   string `json:"pp_units,omitempty"`
}

type ProdDetails struct {
	PName  string `json:"p_name,omitempty"`
	PDesc  string `json:"p_desc,omitempty"`
	PQ     PQ     `json:"p_quantity,omitempty"`
	PPrice PPrice `json:"p_price,omitempty"`
}

type ProdDetList struct {
	ProdDL []ProdDetails
}
