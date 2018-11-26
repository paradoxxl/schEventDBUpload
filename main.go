package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"

	"firebase.google.com/go"
	_ "firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func main() {


	sch_path := os.Getenv("SCH_FB")

	opt := option.WithCredentialsFile(sch_path +"\\sch01-59f65-firebase-adminsdk-nkb52-4ec6f84666.json")
	conf := &firebase.Config{ProjectID: "sch01-59f65"}

	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return
	}



	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()


	//ref := cli.NewRef("regattaCalendar")
	//dataset := ref.Child("MwFQqtscilDNbw3ZkmhT")


	testdata, err := ioutil.ReadFile(sch_path+"\\testdata.json")
	if err != nil {
		fmt.Println(err)
	}

	m := make(map[string]interface{})

	json.Unmarshal(testdata,&m)
	log.Println(m)


	res,err := client.Collection("regattaCalendar").Doc("2018").Create(context.Background(),m)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)


}
