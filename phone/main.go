package main

import (
	"io/ioutil"
	"log"
	"phone/db"
	"regexp"
	"strings"
)

func main() {
	res, err := ioutil.ReadFile("phone_numbers.txt")
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	rawPhones := strings.Split(string(res), "\r\n")

	if err := db.CreatePhoneNumberBatch(rawPhones); err != nil {
		log.Fatal("error inserting:", err)
	}

	//query and update
	normalPhones := make(map[string]struct{}, 0)
	for _, phone := range rawPhones {
		_, id, err := db.QueryPhoneNumberFirst(phone)
		if err != nil {
			log.Fatal("error query:", err)
		}

		r := regexp.MustCompile("\\D")
		newPhone := r.ReplaceAllString(phone, "")
		if _, ok := normalPhones[newPhone]; !ok {
			normalPhones[newPhone] = struct{}{}
		}

		if err := db.UpdatePhoneNumber(newPhone, id); err != nil {
			log.Fatal("error updating new phone number:", err)
		}
	}

	//Distinct
	for normal, _ := range normalPhones {
		ids, err := db.QueryPhoneNumbers(normal)
		if err != nil {
			log.Fatal("query multi error:", err)
		}
		for i := 1; i < len(ids); i++ {
			db.DeletePhoneNumber(ids[i])
		}
	}

	//Verify
	db.PrintAll()
}
