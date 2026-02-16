package products

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

// Static data for three model sportscar products
var Products = []Product{
    {ID: 1, Name: "AMG GT 63 S 1:18", Price: 49.99},
    {ID: 2, Name: "Porsche 992.1 GT3 1:18", Price: 79.99},
    {ID: 3, Name: "Fiat Multipla 1:18", Price: 59.99},
}

// FindByID returns a pointer to the product with the given id or nil if not found
func FindByID(id int) *Product {
    for i := range Products {
        if Products[i].ID == id {
            return &Products[i]
        }
    }
    return nil
}
