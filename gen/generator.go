package gen

import (
	"fmt"
	"go-minikube/api/models"
)

var owners = []models.Owner{
	models.Owner{FirstName: "Bruce", LastName: "Wayne", Email: "brucewayne@email.com", Password: "123456", Gender: "M", Status: 0},
	models.Owner{FirstName: "Barbara", LastName: "Gordon", Email: "barbaragordon@email.com", Password: "123456", Gender: "F", Status: 1},
	models.Owner{FirstName: "Peter", LastName: "Parker", Email: "peterparker@email.com", Password: "123456", Gender: "M", Status: 1},
	models.Owner{FirstName: "Gwen", LastName: "Stacy", Email: "gwenstacy@email.com", Password: "123456", Gender: "F", Status: 1},
	models.Owner{FirstName: "Clark", LastName: "Kent", Email: "clarkkent@email.com", Password: "123456", Gender: "M", Status: 1},
	models.Owner{FirstName: "Susan", LastName: "Storm", Email: "susanstorm@email.com", Password: "123456", Gender: "F", Status: 1},
	models.Owner{FirstName: "Rocket", LastName: "Raccoon", Email: "rocketraccoon@email.com", Password: "123456", Gender: "M", Status: 1},
	models.Owner{FirstName: "Jessica", LastName: "Jones", Email: "jessicajones@email.com", Password: "123456", Gender: "F", Status: 1},
	models.Owner{FirstName: "Matt", LastName: "Murdock", Email: "mattmurdock@email.com", Password: "123456", Gender: "M", Status: 1},
	models.Owner{FirstName: "Jean", LastName: "Grey", Email: "jeangrey@email.com", Password: "123456", Gender: "F", Status: 1},
	models.Owner{FirstName: "Frank", LastName: "Castle", Email: "frankcastle@email.com", Password: "123456", Gender: "M", Status: 0},
}

func GenerateData() {
	for _, owner := range owners {
		models.NewOwner(owner)
	}
	fmt.Println("\r\n", len(owners), "objetos, foram gerados com sucesso...")
}
