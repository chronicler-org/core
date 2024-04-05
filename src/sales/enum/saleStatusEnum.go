package saleEnum

type SaleStatus string

const (
	AWAITING_PAYMENT   SaleStatus = "Aguardando pagamento"
	PURCHASE_CONFIRMED SaleStatus = "Compra confirmada"
	PURCHASE_COMPLETED SaleStatus = "Compra concluida"
	CANCELLED_PURCHASE SaleStatus = "Compra cancelada"
)
