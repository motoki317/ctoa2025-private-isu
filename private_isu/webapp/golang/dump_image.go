package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
)

const imageBaseDir = "/home/isucon/private_isu/webapp/golang/image"

func dumpImage(db *sqlx.DB) {
	var ids []int
	err := db.Select(&ids, "SELECT `id` FROM `posts`")
	if err != nil {
		panic(err)
	}

	for _, id := range ids {
		var post Post
		err := db.Get(&post, "SELECT * FROM `posts` WHERE `id` = ?", id)
		if err != nil {
			panic(err)
		}

		var suffix string
		switch post.Mime {
		case "image/jpeg":
			suffix = ".jpg"
		case "image/png":
			suffix = ".png"
		case "image/gif":
			suffix = ".gif"
		default:
			slog.Error("unknown mime type", "mime", post.Mime)
			continue
		}
		filename := fmt.Sprintf("%s/%d%s", imageBaseDir, id, suffix)

		err = os.WriteFile(filename, post.Imgdata, 0644)
		if err != nil {
			panic(err)
		}
		slog.Info("dumped", "file", filename)
	}

	slog.Info("done", "files", len(ids))
}
