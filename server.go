package main

import (
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"time"
	"strconv"
)

type Post struct {
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `form:"Title"`
	Body    string `form:"Body" binding:"required"`
}

type Ph_spot struct {
	ID      string `db:"ID"`
	COUNTRY string `form:"COUNTRY"`
	PROVINCE   string `form:"PROVINCE"`
	CITY    string `form:"CITY"`
	COUNTY   string `form:"COUNTY"`
	NAME    string `form:"NAME"`
	LEVEL   string `form:"LEVEL"`
	LABEL    string `form:"LABEL"`
	PRICE   string `form:"PRICE"`
	STATUS    int `form:"STATUS"`
}

//type Ph_spot struct {
//	Id      int64 `db:"post_id"`
//	COUNTRY string `form:"COUNTRY" binding:"required"`
//	PROVINCE   string `form:"PROVINCE" binding:"required"`
//	CITY    string `form:"CITY" binding:"required"`
//	COUNTY   string `form:"COUNTY" binding:"required"`
//	NAME    string `form:"NAME" binding:"required"`
//	LEVEL   string `form:"LEVEL" binding:"required"`
//	LABEL    string `form:"LABEL" binding:"required"`
//	PRICE   string `form:"PRICE" binding:"required"`
//	STATUS    int64 `form:"STATUS"`
//}

func (bp Post) Validate(errors *binding.Errors, req *http.Request) {
	//custom validation
	if len(bp.Title) == 0 {
		errors.Fields["title"] = "Title cannot be empty"
	}
}

func (bp Ph_spot) Validate1(errors *binding.Errors, req *http.Request) {
	//custom validation
	if len(bp.NAME) == 0 {
		errors.Fields["title"] = "Name cannot be empty"
	}
}


func main() {
	// initialize the DbMap
	dbmap := initDb()
	defer dbmap.Db.Close()
	// setup some of the database

	// lets start martini and the real code
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
		Funcs: []template.FuncMap{
			{
				"formatTime": func(args ...interface{}) string {
					t1 := time.Unix(args[0].(int64), 0)
					return t1.Format(time.Stamp)
				},
				"unescaped": func(args ...interface{}) template.HTML {
					return template.HTML(args[0].(string))
				},
			},
		},
	}))

	m.Get("/", func(r render.Render) {
		//fetch all rows
		var posts []Ph_spot
		_, err:= dbmap.Select(&posts, "select * from PH_VIEW_SPOTS order by ID")
		checkErr(err, "Select failed")
		//log.Printf(phViewSpots)
		newmap := map[string]interface{}{"metatitle": "this is my custom title", "posts": posts}
		r.HTML(200, "posts", newmap)
//		r.HTML(200, "layout", "Kanghui")
	})

	m.Get("/:id", func(args martini.Params, r render.Render) {
		var post Ph_spot

		err:= dbmap.SelectOne(&post, "select * from ph_view_spots where id=?", args["id"])

		//simple error check
		if err != nil {
			newmap := map[string]interface{}{"metatitle": "404 Error", "message": "This is not found"}
			r.HTML(404, "error", newmap)
		} else {
			newmap := map[string]interface{}{"metatitle": post.NAME + " more custom", "post": post}
			r.HTML(200, "post", newmap, render.HTMLOptions{

			})
		}
	})

	//shows how to create with binding params
	m.Post("/", binding.Bind(Ph_spot{}), func(post Ph_spot, r render.Render) {

		p1 := newPost(post.COUNTRY, post.PROVINCE,  post.CITY,  post.COUNTY, post.NAME, post.LEVEL, post.LABEL, post.PRICE,post.STATUS)

		log.Println(p1.ID)
		log.Println(p1.COUNTRY)
		log.Println(p1.PROVINCE)
		log.Println(p1.CITY)
		log.Println(p1.NAME)

		err:= dbmap.Insert(&p1)
		checkErr(err, "Insert failed")

		newmap := map[string]interface{}{"metatitle": "created post", "post": p1}
		r.HTML(200, "post", newmap, render.HTMLOptions{

		})
	})
//	//delete the spot
//	m.Post("/Deletespot", binding.Bind(Post{}), func(post Post, r render.Render) {
//
//		p1 := newPost(post.Title, post.Body)
//
//		log.Println(p1)
//
//		err:= dbmap.Insert(&p1)
//		checkErr(err, "Insert failed")
//
//		newmap := map[string]interface{}{"metatitle": "created post", "post": p1}
//		r.HTML(200, "post", newmap, render.HTMLOptions{
//
//		})
//	})
//	jump to admin/viewspots page
	m.Get("/admin/viewSpots", func(r render.Render) {
//		r.HTML(200, "viewSpots", "这是老何的测试数据-viewspots", render.HTMLOptions{
//			Layout: "viewSpots",
//		})

		//fetch all rows
		var posts []Ph_spot
		_, err:= dbmap.Select(&posts, "select * from ph_view_spots")
		checkErr(err, "Select failed")

		newmap := map[string]interface{}{"metatitle": "this is my custom title", "posts": posts}

		r.HTML(200, "posts", newmap, render.HTMLOptions{

		})

	})


	//jump to login page
	m.Get("/user/login/",func(r render.Render){
		r.HTML(200, "login","",render.HTMLOptions{

		})
	})


	m.Run()

}

//func newPost(title, body string) Ph_spot {
//	return Ph_spot{
//		//Created: time.Now().UnixNano(),
//		Created: time.Now().Unix(),
//		Title:   title,
//		Body:    body,
//	}
//}

func newPost(COUNTRY, PROVINCE, CITY, COUNTY,NAME,LEVEL,LABEL,PRICE string ,Status int) Ph_spot {
	return Ph_spot{
		ID: strconv.FormatInt(time.Now().UnixNano(), 2),
		COUNTRY:  COUNTRY,
		PROVINCE:   PROVINCE,
		CITY:    CITY,
		COUNTY:COUNTY,
		NAME:NAME,
		LEVEL:LEVEL,
		LABEL:LABEL,
		PRICE:PRICE,
		STATUS:Status,
	}
}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish

	//db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	//db, err := sql.Open("mysql", "USERNAME:PASSWORD@unix(/var/run/mysqld/mysqld.sock)/sample")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/kanghui")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	// dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
//	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
	dbmap.AddTableWithName(Ph_spot{}, "ph_view_spots")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}


