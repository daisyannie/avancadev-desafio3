package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) Check(code string) string {

	for _, item := range c.Coupon {
		if code == item.Code {
			return "valid"
		}
	}

	return "invalid"

}

type Result struct {
	Status string
}

var coupons Coupons

func main() {
	coupon := Coupon {
		Code: "abc",
	}

	coupons.Coupon = append(coupons.Coupon, coupon)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	isValid := coupons.Check(coupon)

	result := Result{Status: isValid}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
}