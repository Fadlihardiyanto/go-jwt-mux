package productcontroller

import (
	"net/http"

	"github.com/Fadlihardiyanto/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "kemeja",
			"harga":        100000,
			"stok":         10,
		},
		{
			"id":           2,
			"nama_product": "celana",
			"harga":        100000,
			"stok":         10,
		},
		{
			"id":           3,
			"nama_product": "topi",
			"harga":        100000,
			"stok":         10,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
