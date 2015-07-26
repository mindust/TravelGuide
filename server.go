package main

import (
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"time"
	"os"
	"io"
	"strings"
)

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
	ID      string
	COUNTRY string `form:"COUNTRY" binding:"required"`
	PROVINCE   string `form:"PROVINCE" binding:"required"`
	CITY    string `form:"CITY" binding:"required"`
	COUNTY   string `form:"COUNTY" binding:"required"`
	NAME    string `form:"NAME" binding:"required"`
	LEVEL   string `form:"LEVEL" binding:"required"`
	GRADE  string  `form:"GRADE" binding:"required"`
    LABEL    string `form:"LABEL" binding:"required"`
	HIGHLIGHTS1    string `form:"HIGHLIGHTS1"`
	HIGHLIGHTS2    string `form:"HIGHLIGHTS2"`
	HIGHLIGHTS3    string `form:"HIGHLIGHTS3"`
	PRICE   string `form:"PRICE" binding:"required"`
	STATUS    string `form:"STATUS"`
}

type Ph_spot_with_image struct {
	ID				string
	NAME			string
	VIEW_SPOT_ID	string
	SOURCE_NAME		string
	FORMAT			string
	PATH			string
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
			r.HTML(200, "modify", newmap, render.HTMLOptions{
				Layout: "admin_layout",
			})
		}
	})

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
//
//	m.Get("/admin/upload",func(r render.Render){
//		r.HTML(200, "upload","",render.HTMLOptions{
//			Layout: "admin_layout",
//		})
//	})

	m.Post("/upload", binding.Bind(Ph_spot{}), func(spot Ph_spot, args martini.Params, w http.ResponseWriter, r *http.Request) (int, string){
		err := r.ParseMultipartForm(100000)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		p1 := newPost(uuid.NewV4().String(), spot.COUNTRY, spot.PROVINCE,  spot.CITY,  spot.COUNTY, spot.NAME, spot.LEVEL, spot.GRADE, spot.HIGHLIGHTS1, spot.HIGHLIGHTS2, spot.HIGHLIGHTS3, spot.LABEL, spot.PRICE, spot.STATUS)
		log.Println(p1)
		err = dbmap.Insert(&p1)
		checkErr(err, "Insert Spot failed")

		files := r.MultipartForm.File["IMAGE"]
		//newSpotImages := []Ph_spot_with_image
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			//获取扩展名
			sourceImageNameArr := strings.SplitAfter(files[i].Filename, ".")
			sourceImageNameExt := sourceImageNameArr[len(sourceImageNameArr)-1]
			u1 := uuid.NewV4().String()
			dstSource, err := os.Create("./uploads/source/" + u1+"_s."+sourceImageNameExt)
			dstAlbum, err := os.Create("./uploads/album/" + u1+"_a."+sourceImageNameExt)
			defer dstSource.Close()
			defer dstAlbum.Close()
			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			if _, err := io.Copy(dstSource, file); err != nil {
				return http.StatusInternalServerError, err.Error()
			}
			if _, err := io.Copy(dstAlbum, file); err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			image1 := newSpotImage(uuid.NewV4().String(), u1, p1.ID, files[i].Filename, sourceImageNameExt,"uploads")
			log.Println(image1)
			//保存图片信息到图片数据库
			err = dbmap.Insert(&image1)
			checkErr(err, "Insert iamge failed")
		}


		return 200, "ok"

	})


	m.Post("/update", binding.Bind(Ph_spot{}), func(spot Ph_spot, args martini.Params, w http.ResponseWriter, r *http.Request) (int, string){
		err := r.ParseMultipartForm(100000)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}

		p1 := newPost(uuid.NewV4().String(), spot.COUNTRY, spot.PROVINCE,  spot.CITY,  spot.COUNTY, spot.NAME, spot.LEVEL, spot.GRADE, spot.HIGHLIGHTS1, spot.HIGHLIGHTS2, spot.HIGHLIGHTS3, spot.LABEL, spot.PRICE, spot.STATUS)
		log.Println(p1)
		err = dbmap.Insert(&p1)
		checkErr(err, "Update Spot failed")

		files := r.MultipartForm.File["IMAGE"]
		//newSpotImages := []Ph_spot_with_image
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			//获取扩展名
			sourceImageNameArr := strings.SplitAfter(files[i].Filename, ".")
			sourceImageNameExt := sourceImageNameArr[len(sourceImageNameArr)-1]
			u1 := uuid.NewV4().String()
			dstSource, err := os.Create("./uploads/source/" + u1+"_s."+sourceImageNameExt)
			dstAlbum, err := os.Create("./uploads/album/" + u1+"_a."+sourceImageNameExt)
			defer dstSource.Close()
			defer dstAlbum.Close()
			if err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			if _, err := io.Copy(dstSource, file); err != nil {
				return http.StatusInternalServerError, err.Error()
			}
			if _, err := io.Copy(dstAlbum, file); err != nil {
				return http.StatusInternalServerError, err.Error()
			}

			image1 := newSpotImage(uuid.NewV4().String(), u1, p1.ID, files[i].Filename, sourceImageNameExt,"uploads")
			log.Println(image1)
			//保存图片信息到图片数据库
			err = dbmap.Insert(&image1)
			checkErr(err, "Insert iamge failed")
		}


		return 200, "ok"

	})

	m.Run()
}


func newPost(UUID, COUNTRY, PROVINCE, CITY, COUNTY,NAME,LEVEL,GRADE,HIGHLIGHTS1,HIGHLIGHTS2,HIGHLIGHTS3,LABEL,PRICE,STATUS string ) Ph_spot {
	return Ph_spot{
		ID: UUID,
		COUNTRY:  COUNTRY,
		PROVINCE:   PROVINCE,
		CITY:    CITY,
		COUNTY:COUNTY,
		NAME:NAME,
		LEVEL:LEVEL,
		GRADE:GRADE,
		HIGHLIGHTS1:HIGHLIGHTS1,
		HIGHLIGHTS2:HIGHLIGHTS2,
		HIGHLIGHTS3:HIGHLIGHTS3,
		LABEL:LABEL,
		PRICE:PRICE,
		STATUS:STATUS,
	}
}

func newSpotImage(UUID, NAME, VIEW_SPOT_ID, SOURCE_NAME, FORMAT, PATH string) Ph_spot_with_image{
	return Ph_spot_with_image{
		ID: UUID,
		NAME: NAME,
		VIEW_SPOT_ID: VIEW_SPOT_ID,
		SOURCE_NAME: SOURCE_NAME,
		FORMAT: FORMAT,
		PATH: PATH,
	}
}


func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/kanghui")
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Ph_spot{}, "ph_view_spots")
	dbmap.AddTableWithName(Ph_spot_with_image{}, "ph_view_spot_images")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}


