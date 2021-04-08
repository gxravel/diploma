package db

import (
	"log"

	uuid "github.com/satori/go.uuid"
)

var gList []*Genre

func init() {
	id, err := uuid.NewV4()
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	fictionID := id.String()
	id, err = uuid.NewV4()
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	storyID := id.String()
	id, err = uuid.NewV4()
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	novelID := id.String()
	gList = []*Genre{
		&Genre{
			ID:   fictionID,
			Name: "художественная литература",
		},
		&Genre{
			ID:     novelID,
			Name:   "роман",
			Parent: fictionID,
		},
		&Genre{
			Name:   "роман-эпопея",
			Parent: novelID,
		},
		&Genre{
			Name:   "исторический роман",
			Parent: novelID,
		},
		&Genre{
			Name:   "трагедия",
			Parent: fictionID,
		},
		&Genre{
			ID:     storyID,
			Name:   "повесть",
			Parent: fictionID,
		},
		&Genre{
			Name:   "юмористическая повесть",
			Parent: storyID,
		},
		&Genre{
			Name:   "комедия",
			Parent: fictionID,
		},
		&Genre{
			Name:   "драма",
			Parent: fictionID,
		},
	}
}
