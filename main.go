package main

import (
	"fmt"

	"github.com/funayoseyoshito/yakiniku-image-id/db"
	"github.com/funayoseyoshito/yakiniku-image-id/lib"
)

func main() {
	dbSet := db.NewDatabaseSet(
		lib.Config.Database.User,
		lib.Config.Database.Password,
		lib.Config.Database.Host,
		lib.Config.Database.Port,
		lib.Config.Database.Name)

	var image db.Images
	defer dbSet.Connection().Close()

	var cnt int = 0
	for i := 0; ; {
		rows := dbSet.SelectProcessingRows(i)

		if !rows.Next() {
			break
		}

		for {
			dbSet.Connection().ScanRows(rows, &image)
			fmt.Println(image.ID, image.Kind, image.OriginID)
			image.OriginID = image.ID
			dbSet.Connection().Save(&image)
			cnt++

			if !rows.Next() {
				break
			}
		}
		i += db.SelectLimit
	}
	fmt.Println("処理件数", cnt)
}
