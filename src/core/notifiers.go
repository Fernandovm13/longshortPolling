package core

import "fmt"

var (
	ProductNotifier  = make(chan struct{})
	CategoryNotifier = make(chan struct{})
)

// notifica un cambio en productos y muestra un mensaje en la terminal.
func NotifyProductUpdate() {
	fmt.Println("[NOTIFICACIÓN] Se actualizó un producto")
	close(ProductNotifier)
	ProductNotifier = make(chan struct{})
}

// notifica un cambio en categorías y muestra un mensaje en la terminal.
func NotifyCategoryUpdate() {
	fmt.Println("[NOTIFICACIÓN] Se actualizó una categoría")
	close(CategoryNotifier)
	CategoryNotifier = make(chan struct{})
}
