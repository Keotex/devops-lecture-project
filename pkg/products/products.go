package products

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

// Static data for three products
var Products = []Product{
    {ID: 1, Name: "Product 1", Price: 19.99},
    {ID: 2, Name: "Product 2", Price: 29.99},
    {ID: 3, Name: "Product 3", Price: 39.99},
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
