package db

import (
	"log"
	"time"
)

const authorsNumber = 4

var aList []*Author

func init() {
	location, err := time.LoadLocation("Local")
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	aList = []*Author{
		&Author{
			Name:         "Лев Николаевич Толстой",
			OriginalName: "Лев Николаевич Толстой",
			BirthDate:    time.Date(1828, time.August, 28, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1910, time.November, 10, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Джек Лондон",
			OriginalName: "Jack London",
			BirthDate:    time.Date(1876, time.January, 12, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1916, time.November, 22, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Уильям Сомерсет Моэм",
			OriginalName: "William Somerset Maugham",
			BirthDate:    time.Date(1874, time.January, 25, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1965, time.December, 16, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Уильям Шекспир",
			OriginalName: "William Shakespeare",
			BirthDate:    time.Date(1564, time.April, 26, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1616, time.April, 23, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Джордж Бернард Шоу",
			OriginalName: "George Bernard Shaw",
			BirthDate:    time.Date(1856, time.July, 26, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1950, time.November, 2, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Оскар Уайльд",
			OriginalName: "Oscar Wilde",
			BirthDate:    time.Date(1564, time.April, 26, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1616, time.April, 23, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Джером Клапка Джером",
			OriginalName: "Jerome Klapka Jerome",
			BirthDate:    time.Date(1859, time.May, 2, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1927, time.June, 14, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Александр Дюма",
			OriginalName: "Alexandre Dumas",
			BirthDate:    time.Date(1802, time.July, 24, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1870, time.December, 5, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Александр Иванович Куприн",
			OriginalName: "Александр Иванович Куприн",
			BirthDate:    time.Date(1870, time.August, 26, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1938, time.August, 25, 0, 0, 0, 0, location),
		},
		&Author{
			Name:         "Иоганн Вольфганг фон Гёте",
			OriginalName: "Johann Wolfgang von Goethe",
			BirthDate:    time.Date(1749, time.August, 28, 0, 0, 0, 0, location),
			DeathDate:    time.Date(1832, time.March, 22, 0, 0, 0, 0, location),
		},
	}
}
