package seeddb

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/lsortudo/blog-golang/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Leonardo Santos",
		Email:    "LeonardoTeste@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Carlos Souza",
		Email:    "CarlosSouza@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Primeiro post",
		Content: "Hello world",
	},
	models.Post{
		Title:   "Segundo post",
		Content: "Nesse segundo post irei testar o limite de caracteres pra ver se encontro alguma restricao antes do arquivo de teste ficar pronto ---- Nesse segundo post irei testar o limite de caracteres pra ver se encontro alguma restricao antes do arquivo de teste ficar pronto",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
