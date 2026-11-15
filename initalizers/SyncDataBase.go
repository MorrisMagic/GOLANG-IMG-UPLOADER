package initalizers

import (
	"fmt"
	"img-uploader/models"
)

func SyncData() {

	err := DB.AutoMigrate(&models.IMG{})
	if err != nil {
		panic(fmt.Sprintf("Error migrating database: %v", err))
	}

	fmt.Println("Database migrated successfully!")

}
