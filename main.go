package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type CRM struct {
	ID                int    `pg:"id,pk"`
	PartnerUsers      string `pg:"partner_users"`
	ActiveLead        int    `pg:"active_lead"`
	UnassignedLeads   int    `pg:"unassigned_leads"`
	ActiveProspect    int    `pg:"active_prospect"`
	ActiveOpportunity int    `pg:"active_opportunity"`
}

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "kaif",
		Password: "kaif",
		Database: "mydb",
		Addr:     "localhost:5432",
	}

	db := pg.Connect(opts)
	if db == nil {
		fmt.Print("failed to connect to the database")
		os.Exit(100)
	} else {

		fmt.Println("Hello, you are connected to the database")
	}
	return db
}

func main() {
	db := Connect()
	server := gin.Default()
	server.Use(cors.Default())
	server.GET("/crm", func(ctx *gin.Context) {
		var crm []CRM
		err := db.Model(&crm).Select()
		if err != nil {
			fmt.Printf("Error fetching data: %v\n", err)
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, crm)
	})

	server.Run()

}
