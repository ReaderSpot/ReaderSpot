package models

type Compras struct {
	ID       uint `json:"id"`
	Cantidad int  `json:"cantidad"`
}

type CarritoCompras []Compras

type DetallesLibro struct {
	ID       uint
	Cantidad int
	Titulo   string
	Precio   float64
	Subtotal float64
}
