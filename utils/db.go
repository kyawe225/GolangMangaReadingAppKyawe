package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "file:sample.sqlite?cache=shared")

	if err != nil {
		panic(err)
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(5)

	initializeDatabase()
}

func initializeDatabase() {
	transaction, err := DB.Begin()
	if err != nil {
		panic(err)
	}

	createMangaTable := `
		create table if not exists mangas(
			id varchar(125) primary key, 
			name varchar(255) not null,
			description varchar(255) not null,
			publish_date datetime not null,
			is_published tinyint not null default 0,
			published_url varchar(255) not null,
			created_at datetime not null default CURRENT_TIMESTAMP ,
			updated_at datetime not null default CURRENT_TIMESTAMP
		)
	`

	_, err = transaction.Exec(createMangaTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	createCategoryTable := `
		create table if not exists category(
			id varchar(125) primary key, 
			name varchar(255) not null,
			description varchar(255) not null,
			created_at datetime not null default CURRENT_TIMESTAMP ,
			updated_at datetime not null default CURRENT_TIMESTAMP
		)
	`
	_, err = transaction.Exec(createCategoryTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	createMangaCategoryTable := `
		create table if not exists manga_category(
			id varchar(125) primary key, 
			manga_id varchar(125) not null,
			category_id varchar(125) not null,
			created_at datetime not null default CURRENT_TIMESTAMP ,
			updated_at datetime not null default CURRENT_TIMESTAMP,
			foreign key(manga_id) references mangas(id),
			foreign key(category_id) references category(id)
		)
	`
	_, err = transaction.Exec(createMangaCategoryTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	createChapterTable := `
		create table if not exists chapter(
			id varchar(125) primary key not null,
			manga_id varchar(125) not null,
			chapter_name varchar(255) not null,
			name varchar(225) not null,
			description varchar(255) not null,
			is_published tinyint default 0,
			publish_url varchar(125) not null,
			created_at datetime not null default CURRENT_TIMESTAMP,
			updated_at datetime not null default CURRENT_TIMESTAMP
		)
	`
	_, err = transaction.Exec(createChapterTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}
	createChapterPictureTable := `
		create table if not exists chapter_pictures(
			id varchar(125) primary key not null,
			picture_data text not null,
			chapter_id varchar(125) not null,
			serial int not null,
			created_at datetime not null default CURRENT_TIMESTAMP,
			updated_at datetime not null default CURRENT_TIMESTAMP
		)
	`

	_, err = transaction.Exec(createChapterPictureTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	createUserTable := `
		create table if not exists users(
			id varchar(125) primary key not null,
			name varchar(225) not null,
			email varchar(225) not null,
			password varchar(225) not null,
			birthdate datetime not null,
			role varchar(120) not null,
			created_at datetime not null default CURRENT_TIMESTAMP,
			updated_at datetime not null default CURRENT_TIMESTAMP
		)
	`

	_, err = transaction.Exec(createUserTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	transaction.Commit()
}
