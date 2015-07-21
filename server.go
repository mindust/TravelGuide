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
	"fmt"
	"os"
	"io"
)

type Post struct {
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `form:"Title"`
	Body    string `form:"Body" binding:"required"`
}

//type Ph_spot struct {
//	ID      string `db:"ID"`
//	COUNTRY string `form:"COUNTRY"`
//	PROVINCE   string `form:"PROVINCE"`
//	CITY    string `form:"CITY"`
//	COUNTY   string `form:"COUNTY"`
//	NAME    string `form:"NAME"`
//	LEVEL   string `form:"LEVEL"`
//	LABEL    string `form:"LABEL"`
//	PRICE   string `form:"PRICE"`
//	STATUS    int `form:"STATUS"`
//}

type Ph_users struct {
	ID      int `db:"ID"`
	USER_NAME string `form:"COUNTRY" binding:"required"`
	PASSWORD   string `form:"PROVINCE" binding:"required"`
	DESCRIPTION    string `form:"CITY" binding:"required"`
	QQ_NO   string `form:"COUNTY" binding:"required"`
	MOBILE    string `form:"NAME" binding:"required"`
	EMAIL   string `form:"LEVEL" binding:"required"`
	LEVEL    string `form:"LABEL" binding:"required"`
	LOGO_PATH   string `form:"PRICE" binding:"required"`
}

type Ph_spot struct {
	ID      int `db:"ID"`
	COUNTRY string `form:"COUNTRY" binding:"required"`
	PROVINCE   string `form:"PROVINCE" binding:"required"`
	CITY    string `form:"CITY" binding:"required"`
	COUNTY   string `form:"COUNTY" binding:"required"`
	NAME    string `form:"NAME" binding:"required"`
	LEVEL   string `form:"LEVEL" binding:"required"`
	LABEL    string `form:"LABEL" binding:"required"`
	PRICE   string `form:"PRICE" binding:"required"`
	STATUS    int `form:"STATUS"`
}

type Image struct {
	Version    string `form:"version"`
	Email      string `form:"user_email"`
	PrivateKey string `form:"user_private_key"`
	Host       string `form:"file_host"`
	Album      string `form:"file_album"`
	Name       string `form:"file_name"`
	Mime       string `form:"file_mime"`
	unexported string `form:"-"` // skip binding of unexported fields
}

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
		Charset:   "UTF-8",
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
				Layout: "admin_layout",
			})
		}
	})

	//shows how to create with binding params
	m.Post("/admin/viewSpots", binding.Bind(Ph_spot{}), func(post Ph_spot, r render.Render) {

		p1 := newPost(post.COUNTRY, post.PROVINCE,  post.CITY,  post.COUNTY, post.NAME, post.LEVEL, post.LABEL, post.PRICE,post.STATUS)
//		log.Println(p1.ID)
//		log.Println(p1.COUNTRY)
//		log.Println(p1.PROVINCE)
//		log.Println(p1.CITY)
//		log.Println(p1.NAME)
		log.Println(p1)
		err:= dbmap.Insert(&p1)
		checkErr(err, "Insert failed")

		newmap := map[string]interface{}{"metatitle": "created post", "post": p1}
		r.HTML(200, "post", newmap, render.HTMLOptions{
			Layout: "admin_layout",
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
		//fetch all rows
		var posts []Ph_spot
		_, err:= dbmap.Select(&posts, "select * from ph_view_spots")
		checkErr(err, "Select failed")
		newmap := map[string]interface{}{"metatitle": "this is my custom title", "posts": posts}
		r.HTML(200, "posts", newmap, render.HTMLOptions{
			Layout: "admin_layout",
		})
	})
	//jump to login page
	m.Get("/user/login/",func(r render.Render){
		r.HTML(200, "login","",render.HTMLOptions{
			Layout: "admin_layout",
		})
	})
	m.Get("/admin/create/",func(r render.Render){
		r.HTML(200, "createpost","",render.HTMLOptions{
			Layout: "admin_layout",
		})
	})


	m.Get("/admin/upload",func(r render.Render){
		r.HTML(200, "upload","",render.HTMLOptions{
			Layout: "admin_layout",
		})
	})

	m.Post("/upload", binding.Bind(Image{}), func(image Image, args martini.Params, w http.ResponseWriter, r *http.Request) (int, string){
		err := r.ParseMultipartForm(100000)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		files := r.MultipartForm.File["file_album"]
		for i, _ := range files {
			log.Println("getting handle to file")
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			log.Println("creating destination file")
			dst, err := os.Create("./uploads/" + files[i].Filename)
			defer dst.Close()
			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			log.Println("copying the uploaded file to the destination file")
			if _, err := io.Copy(dst, file); err != nil {
				return http.StatusInternalServerError, err.Error()
			}
		}

		fmt.Printf("file_name: " + image.Name + "\n")          // now works
		fmt.Printf("version: " + image.Version + "\n")          // now works
		fmt.Printf("Album: " + image.Album + "\n")          // now works
		fmt.Printf("Mime: " + image.Mime + "\n")          // now works
		return 200, "ok"
	})
	m.Run()
}


func newPost(COUNTRY, PROVINCE, CITY, COUNTY,NAME,LEVEL,LABEL,PRICE string ,Status int) Ph_spot {
	return Ph_spot{
//		ID: strconv.FormatInt(time.Now().UnixNano(), 2),
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

//func loginpost(USER_NAME, PASSWORD, DESCRTPTION, QQ_NO,MOBILE,EMAIL,LEVEL,LOGO_PATH string) Ph_spot {
//	return Ph_spot{
//		//		ID: strconv.FormatInt(time.Now().UnixNano(), 2),
//		USER_NAME:  USER_NAME,
//		PASSWORD:   PASSWORD,
//		DESCRTPTION:    DESCRTPTION,
//		QQ_NO:QQ_NO,
//		MOBILE:MOBILE,
//		EMAIL:EMAIL,
//		LEVEL:LEVEL,
//		LOGO_PATH:LOGO_PATH,
//	}
//}

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
	//err = dbmap.CreateTablesIfNotExists()
	//checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}


