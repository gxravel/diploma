package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/rav1L/book-machine/modules/db"
)

const (
	dbDriver = "postgres"
)

var (
	myDB   db.ISQL
	dbPath = os.Getenv("DATABASE_URL")
)

func main() {
	myDB = &db.Handler{}
	err := myDB.Init(dbDriver, dbPath)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer myDB.Disconnect()
	// err = updateUsers()
	// if err != nil {
	// 	log.Printf("%+v", err)
	// 	return
	// }
	// err = updateGenres()
	// if err != nil {
	// 	log.Printf("%+v", err)
	// 	return
	// }
	// err = updateAuthors()
	// if err != nil {
	// 	log.Printf("%+v", err)
	// 	return
	// }
	// err = updateBookGenres()
	// if err != nil {
	// 	log.Printf("%+v", err)
	// 	return
	// }
	// err = updateBooks()
	// if err != nil {
	// 	log.Printf("%+v", err)
	// 	return
	// }
	// err = updateMyCollections()
	// if err != nil {
	// 	log.Printf("%+v", err)
	// 	return
	// }
	// err = updateReviews()
	// if err != nil {
	// 	log.Printf("%+v", err)
	// 	return
	// }
	err = updateBookGenres()
	if err != nil {
		log.Printf("%+v", err)
		return
	}
}

func updateUsers() (err error) {
	err = myDB.Users().DropTable()
	if err != nil {
		return
	}
	err = myDB.Users().CreateTable()
	if err != nil {
		return
	}
	err = myDB.Users().Init()
	return
}

func updateAuthors() (err error) {
	err = myDB.Authors().DropTable()
	if err != nil {
		return
	}
	err = myDB.Authors().CreateTable()
	if err != nil {
		return
	}
	err = myDB.Authors().Init()
	return
}

func updateBooks() (err error) {
	err = myDB.Books().DropTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.Books().CreateTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.Books().Init()
	return errors.WithStack(err)

}

func updateMyCollections() (err error) {
	err = myDB.Collections().DropTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.Collections().CreateTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.Collections().Init()
	return errors.WithStack(err)

}

func updateGenres() (err error) {
	err = myDB.Genres().DropTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.Genres().CreateTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.Genres().Init()
	return errors.WithStack(err)

}

func updateBookGenres() (err error) {
	err = myDB.BookGenres().DropTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.BookGenres().CreateTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.BookGenres().Init()
	return errors.WithStack(err)

}

func updateReviews() (err error) {
	err = myDB.Reviews().DropTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.Reviews().CreateTable()
	if err != nil {
		return errors.WithStack(err)

	}
	return errors.WithStack(err)

}

func updateRating() (err error) {
	err = myDB.Rating().DropTable()
	if err != nil {
		return errors.WithStack(err)

	}
	err = myDB.Rating().CreateTable()
	if err != nil {
		return errors.WithStack(err)

	}
	return errors.WithStack(err)

}
