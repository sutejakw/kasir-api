package models

// ProdukTerlaris is the best-selling product in a period.
type ProdukTerlaris struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

// SalesSummaryResponse is the JSON response for report endpoints.
type SalesSummaryResponse struct {
	TotalRevenue    int             `json:"total_revenue"`
	TotalTransaksi  int             `json:"total_transaksi"`
	ProdukTerlaris *ProdukTerlaris `json:"produk_terlaris,omitempty"`
}
