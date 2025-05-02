package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

type numberData struct {
	Number       uint64   `json:"number"`
	PrimeNumbers []uint64 `json:"prime_numbers"`
	Total        uint64   `json:"total"`
}

type resultData []numberData

func main() {
	http.HandleFunc("/", requestHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	reqNumber := r.FormValue("number")

	num, err := strconv.ParseUint(reqNumber, 10, 32)
	if err != nil {
		http.Error(w, "Error input type, must be a positive number", http.StatusBadRequest)
		return
	}

	result := processNumber(num)

	json.NewEncoder(w).Encode(result)
}

func processNumber(n uint64) resultData {
	res := resultData{}

	for i := 0; i < 10; i++ {
		cn := n + uint64(i) // current number

		ps := getNext7PrimeNumbers(cn) // get next 7 prime numbers
		ttl := getTotal(cn, ps)        // get sum of number and prime numbers

		res = append(res, numberData{
			Number:       cn,
			PrimeNumbers: ps,
			Total:        ttl,
		})
	}

	return res
}

func getNext7PrimeNumbers(n uint64) []uint64 {
	pn := []uint64{}
	nn := n

	for {
		if len(pn) == 7 {
			break
		}

		nn++ // next number

		if isPrime(nn) {
			pn = append(pn, nn)
		}
	}

	return pn
}

func isPrime(n uint64) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 3; i <= sqrt; i += 2 {
		if n%uint64(i) == 0 {
			return false
		}
	}

	return true
}

func getTotal(cn uint64, ps []uint64) uint64 {
	var total uint64 = 0

	for _, v := range ps {
		total += v
	}

	total += cn

	return uint64(total)
}
