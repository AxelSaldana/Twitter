package main

func main() {

}
package main
import(
	"log"
	"github.com/Forest24/Twitter/handlers"
	"github.com/Forest24/Twitter/bd"
)
func main() {
	if bd.ChequeoConnection()==0{
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
	



}