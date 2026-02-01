package main

import (
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"PORT"`
	DBConn 	string `mapstructure:"DB_CONN"`
}


type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}



// var produk = []Produk{
// 	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
// 	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
// 	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
// }

// type Category struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// var categories = []Category{
// 	{ID: 1, Name: "Books", Description: "All kinds of books"},
// 	{ID: 2, Name: "Electronics", Description: "Gadgets and devices"},
// 	{ID: 3, Name: "Clothing", Description: "Apparel and accessories"},
// }

// func getProdukByID(w http.ResponseWriter, r *http.Request) {
// 	// Parse ID dari URL path
// 	// URL: /api/produk/123 -> ID = 123
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Cari produk dengan ID tersebut
// 	for _, p := range produk {
// 		if p.ID == id {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(p)
// 			return
// 		}
// 	}

// 	// Kalau tidak found
// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

// // PUT localhost:8080/api/produk/{id}
// func updateProduk(w http.ResponseWriter, r *http.Request) {
// 	// get id dari request
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

// 	// ganti int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}

// 	// get data dari request
// 	var updateProduk Produk
// 	err = json.NewDecoder(r.Body).Decode(&updateProduk)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	// loop produk, cari id, ganti sesuai data dari request
// 	for i := range produk {
// 		if produk[i].ID == id {
// 			updateProduk.ID = id
// 			produk[i] = updateProduk

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(updateProduk)
// 			return
// 		}
// 	}
	
// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

// func deleteProduk(w http.ResponseWriter, r *http.Request) {
// 	// get id
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	
// 	// ganti id int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}
	
// 	// loop produk cari ID, dapet index yang mau dihapus
// 	for i, p := range produk {
// 		if p.ID == id {
// 			// bikin slice baru dengan data sebelum dan sesudah index
// 			produk = append(produk[:i], produk[i+1:]...)
			
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "sukses delete",
// 			})
// 			return
// 		}
// 	}

// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

// func getCategoryByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	for _, category := range categories {
// 		if category.ID == id {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(category)
// 			return
// 		}
// 	}

// 	http.Error(w, "Category not found", http.StatusNotFound)
// }

// func updateCategory(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	var updateCategory Category
// 	err = json.NewDecoder(r.Body).Decode(&updateCategory)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	for i := range categories {
// 		if categories[i].ID == id {
// 			updateCategory.ID = id
// 			categories[i] = updateCategory

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(updateCategory)
// 			return
// 		}
// 	}

// 	http.Error(w, "Category not found", http.StatusNotFound)
// }

// func deleteCategory(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	for i, p := range categories {
// 		if p.ID == id {
// 			categories = append(categories[:i], categories[i+1:]...)

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "Category deleted",
// 			})
// 			return
// 		}
// 	}

// 	http.Error(w, "Category not found", http.StatusNotFound)
// }


func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}


   // Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)


	// Setup routes
	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)


	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}