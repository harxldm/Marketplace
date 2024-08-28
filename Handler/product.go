package handler

import (
	model "backend_en_go/Model"
	"backend_en_go/Storage"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func NewProduct(w http.ResponseWriter, r *http.Request) {
	setupCORS(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var product model.Product

	// Decodificar el JSON del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println("Error al decodificar el cuerpo de la solicitud:", err)
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Asignar la fecha de creación
	product.CreatedAt = time.Now()

	sellerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Println("Error al obtener el ID del vendedor desde el contexto")
		http.Error(w, "No se pudo obtener el ID del vendedor", http.StatusUnauthorized)
		return
	}

	product.SellerID = sellerID

	// Suponiendo que el SKU es generado por el sistema o ingresado por el usuario
	if product.SKU == "" {
		// Aquí podrías generar un SKU automáticamente, si no se proporciona
		product.SKU = generateSKU(product) // Debes implementar la función generateSKU si es necesario
	}

	// Obtener la conexión a la base de datos
	db := Storage.Pool()

	// Instanciar el repositorio de productos
	productRepo := Storage.NewPsqlProduct(db)

	// Crear el producto
	err = productRepo.CreateProduct(&product)
	if err != nil {
		log.Println("Error al crear el producto:", err)
		http.Error(w, "Error al crear el producto: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Producto creado exitosamente",
		"productID": product.ProductID,
	})
}

func generateSKU(product model.Product) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s-%d", product.Name[:3], rand.Intn(1000)) // Ejemplo: "PRO-123"
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	db := Storage.Pool()

	// Instanciar el repositorio de productos
	productRepo := Storage.NewPsqlProduct(db)

	data, err := productRepo.GetAll()
	if err != nil {
		http.Error(w, "Hubo un problema al obtener todos los productos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "ok",
		"products": data,
	})
}

func GetProductsBySellerID(w http.ResponseWriter, r *http.Request) {

	// Obtener el `sellerID` del contexto de la solicitud
	sellerID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Println("Error al obtener el ID del vendedor desde el contexto")
		http.Error(w, "No se pudo obtener el ID del vendedor", http.StatusUnauthorized)
		return
	}

	// Instanciar el repositorio de productos
	db := Storage.Pool()
	productRepo := Storage.NewPsqlProduct(db)

	// Obtener los productos por SellerID
	products, err := productRepo.GetProductsBySellerID(sellerID)
	if err != nil {
		http.Error(w, "Hubo un problema al obtener los productos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar la respuesta en formato JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "ok",
		"products": products,
	})
}
