package handlers

import (
    "database/sql"
    "html/template"
    "net/http"
)

func IndexHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT * FROM produk")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var products []map[string]string
        for rows.Next() {
            var id, nama_produk, harga, stok string
            err := rows.Scan(&id, &nama_produk, &harga, &stok)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            products = append(products, map[string]string{
                "id":          id,
                "nama_produk": nama_produk,
                "harga":       harga,
                "stok":        stok,
            })
        }

        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, products)
    }
}