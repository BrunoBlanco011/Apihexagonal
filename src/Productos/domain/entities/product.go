package entities	

type Product struct {
	Id       int             `json:"id"`
	Nombre   string          `json:"nombre"`
	Cantidad int             `json:"cantidad"`
	Precio  int 		   `json:"precio"`
}
