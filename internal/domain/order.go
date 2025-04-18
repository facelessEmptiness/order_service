package domain

type OrderItem struct {
	ProductID string
	Quantity  int32
}

type Order struct {
	ID            string
	UserID        string
	Items         []OrderItem
	PaymentMethod string
	Status        string
}
