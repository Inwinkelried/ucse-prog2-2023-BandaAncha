package model

type PedidoProducto struct {
	CodigoProducto string `bson:"codigo_producto"`
	Nombre         string `bson:"nombre_producto"`
	PesoUnitario   int    `bson:"peso_unitario"`
	PrecioUnitario int    `bson:"precio_unitario"`
	Cantidad       int    `bson:"cantidad"`
}

func (productoPedido PedidoProducto) ObtenerPesoProductoPedido() int {
	return productoPedido.PesoUnitario * productoPedido.Cantidad
}
