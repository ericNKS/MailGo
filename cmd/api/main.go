package main

import (
	"sendMail/cmd/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Routes(r)

	r.Run(":3000")
}

// func main() {
// 	email, err := model.CreateMail("ek.silva.santos@gmail.com", "!59753eric", "smtp.gmail.com", 587)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	msg, err := model.CreateMessage("Primeiro email de teste", []byte("Hello tudo bem"))
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	c, err := campaign.CreateCampaign(email, []string{"meusaopaulo@gmail.com"}, msg)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go c.Execute(&wg)
// 	wg.Wait()

// 	fmt.Println(email)
// 	fmt.Println(msg)
// 	fmt.Println(c)

// }
