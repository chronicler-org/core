package saleEnum

type Status string

const (
	AGUARDANDO_PAGAMENTO Status = "Aguardando pagamento"
	COMPRA_CONFIRMADA    Status = "Compra confirmada"
	COMPRA_CONCLUIDA     Status = "Compra concluida"
	COMPRA_CANCELADA     Status = "Compra cancelada"
)
