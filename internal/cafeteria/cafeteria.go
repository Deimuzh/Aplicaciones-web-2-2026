package cafeteria
import "errors"
var (
	ErrClienteNoEncontrado = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente = errors.New("stock insuficiente")
	ErrSaldoInsuficiente = errors.New("saldo insuficiente del cliente")
)

type Cliente struct {
	ID     int
	Nombre string
	Carrera string
	Saldo float64
}
type Categoria struct {
	ID     int
	Nombre string
}
type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Categoria Categoria 
}
type Pedido struct {
	ID       int
	Cliente  Cliente  
	Producto Producto 
	Cantidad int
	Total    float64
}

type Repository interface {
	GuardarCliente(c Cliente) error
	ObtenerCliente(id int) (Cliente, error)
	ListarClientes() []Cliente

	GuardarProducto(p Producto) error
	ObtenerProducto(id int) (Producto, error)
	ListarProductos() []Producto

	GuardarPedido(p Pedido) error
	ListarPedidos() []Pedido
}
type RepoMemoria struct {
	clientes  []Cliente
	productos []Producto
	pedidos   []Pedido
}
func (r *RepoMemoria) GuardarCliente(c Cliente) error {
	r.clientes = append(r.clientes, c)
	return nil
}

func (r *RepoMemoria) ObtenerCliente(id int) (Cliente, error) {
	for _, c := range r.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return Cliente{}, ErrClienteNoEncontrado
}
func (r *RepoMemoria) ListarClientes() []Cliente {
	return r.clientes
}
func (r *RepoMemoria) GuardarProducto(p Producto) error {
	r.productos = append(r.productos, p)
	return nil
}
func (r RepoMemoria) ObtenerProducto(id int) (Producto, error) {
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}

func (r *RepoMemoria) ListarProductos() []Producto {
	return r.productos
}
func (r *RepoMemoria) GuardarPedido(p Pedido) error {
	r.pedidos = append(r.pedidos, p)
	return nil
}

func (r *RepoMemoria) ListarPedidos() []Pedido {
	return r.pedidos
}
func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{}
}
// Verificación en tiempo de compilación:
// Si RepoMemoria NO cumple Repository, esto da error al compilar.
var _ Repository = (*RepoMemoria)(nil)

