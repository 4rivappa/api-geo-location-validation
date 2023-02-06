package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var lat_bounds = []float32{24.6541753, 24.4635105}
var lng_bounds = []float32{73.7612527, 73.6436533}

func Validate(w http.ResponseWriter, r *http.Request) {
	req_body := make(map[string]string)
	json.NewDecoder(r.Body).Decode(&req_body)

	// DEBUG
	// fmt.Printf("=========================================================================\n")
	// fmt.Printf("%#v\n", req_body)

	x64, _ := strconv.ParseFloat(req_body["x"], 32)
	x := float32(x64)
	y64, _ := strconv.ParseFloat(req_body["y"], 32)
	y := float32(y64)
	pincode, _ := strconv.Atoi(req_body["pincode"])

	if pincode == 313001 {
		if x > lat_bounds[1] && x < lat_bounds[0] {
			if y > lng_bounds[1] && y < lng_bounds[0] {
				json.NewEncoder(w).Encode("correct")
				return
			}
		}
		json.NewEncoder(w).Encode("incorrect")
		return
	}
	json.NewEncoder(w).Encode("pincode did not match")
	return
}
