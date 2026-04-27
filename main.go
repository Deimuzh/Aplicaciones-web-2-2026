package main

import (
	"errors"
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {

	// Crear un repo usando la INTERFAZ (var repo cafeteria.Repository = cafeteria.NewRepoMemoria()).
	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()
	// 2. Guardar clientes
	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "James Molina", Carrera: "TI", Saldo: 20})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Luis alcedo", Carrera: "Sistemas", Saldo: 15})
	// 3. Guardar productos
	cat := cafeteria.Categoria{ID: 1, Nombre: "Bebidas"}
	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Inka cola", Precio: 1.5, Categoria: cat})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "Agua de limon", Precio: 1.0, Categoria: cat})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Jugo de fresa", Precio: 2.0, Categoria: cat})
	// 4. Obtener cliente existente
	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}
	// 5. Obtener cliente NO existente
	_, err = repo.ObtenerCliente(99)
	if err != nil {
		fmt.Println("Error:", err)
		if errors.Is(err, cafeteria.ErrClienteNoEncontrado) {
			fmt.Println("El cliente no existe en la bese de datos")
		}
	}
	// . Listar productos
	fmt.Println("\nLista de productos:")
	for _, p := range repo.ListarProductos() {
		fmt.Println("-", p.Nombre, "$", p.Precio)
	}
	//7. Mostrar que un Pedido contiene el Cliente y Producto completos (no solo IDs).
	cliente, _ := repo.ObtenerCliente(1)
	producto, _ := repo.ObtenerProducto(1)
	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  cliente,
		Producto: producto,
		Cantidad: 2,
		Total:    producto.Precio * 2,
	}
	repo.GuardarPedido(pedido)
	fmt.Println("\nPedido creado:")
	fmt.Println(pedido)
}

