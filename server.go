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
    LABEL    string `form:"LABEL" binding:"required"`
	PRICE   string `form:"PRICE" binding:"required"`
	STATUS    string `form:"STATUS"`
	IMAGEPATH string
	IMAGENAME string
	IMAGEFORMAT string

}

type Spots_image struct  {
	imagepath string
	imagename string
	imageformat string
	name string
}

type Ph_spot_with_image struct {
	ID				string
	NAME			string
	VIEW_SPOT_ID	string
	SOURCE_NAME		string
	FORMAT			string
	PATH			string
}

type ph_travel_packages struct {
	ID string
	NAME string `form:"NAME" binding:"required"`
	DESCRIPTION string `form:"DESCRIPTION" binding:"required"`
	FEE int64 `form:"FEE" binding:"required"`
	START_DATE string `form:"START_DATE" binding:"required"`
	END_DATE string `form:"END_DATE" binding:"required"`
	DAYS int64 `form:"DAYS" binding:"required"`
	HOTELS string `form:"HOTELS" binding:"required"`
	TRANSPOT string `form:"TRANSPOT" binding:"required"`
	PERSON_NUM string `form:"PERSON_NUM" binding:"required"`
	TAGS string `form:"TAGS" binding:"required"`
	CONTENT string `form:"CONTENT" binding:"required"`
	ADVICE string `form:"ADVICE" binding:"required"`
	FEE_INCLUDE string `form:"FEE_INCLUDE" binding:"required"`
	FEE_NOT_INCLUDE string `form:"FEE_NOT_INCLUDE" binding:"required"`
	COLLECTION_ADDRESS string `form:"COLLECTION_ADDRESS" binding:"required"`
	HIGHLIGHTS string `form:"HIGHLIGHTS" binding:"required"`
	VIEW_SPOT_ID string `form:"VIEW_SPOT_ID" binding:"required"`
}


type ph_travel_package_tags struct  {
	ID string
	TRAVEL_PACKAGE_ID string
    CONTENT string
}

type ph_travel_package_highlights struct  {
	ID string
	TRAVEL_PACKAGE_ID string
	CONTENT string

}
type ph_travel_package_view_spots struct {
	ID string
	TRAVEL_PACKAGES_ID string
	VIEW_SPOT_ID string
	DESCRIPTION string
	DAYS int64
	PLAY_ITEMS string
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
		//Layout:    "layout",
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



	/**
	首页
	 */
	m.Get("/", func(r render.Render) {
		r.HTML(200, "main","")
	})

	m.Get("/main/hotSpots", func(r render.Render) {
		//fetch all spots
		var viewSpotList []Ph_spot
		_, err:= dbmap.Select(&viewSpotList, "select * from ph_view_spots")
		checkErr(err, "Select failed")
		newmap := map[string]interface{}{"metatitle": "this is my custom title", "posts": viewSpotList}
		r.HTML(200, "main_hot_spots", newmap)

	})

	m.Get("/main/hotCities", func(r render.Render) {
		var viewSpotList []Ph_spot
		_, err:= dbmap.Select(&viewSpotList, "select * from ph_view_spots")
		checkErr(err, "Select failed")
		newmap := map[string]interface{}{"metatitle": "this is my custom title", "posts": viewSpotList}
		r.HTML(200, "main_hot_cities",newmap)
	})

	m.Get("/spot/:id",func(args martini.Params, r render.Render) {
		var viewSpotList []Ph_spot
		_, err:= dbmap.Select(&viewSpotList, "select * from ph_view_spots where id=?", args["id"])
		checkErr(err, "Select failed")
		newmap := map[string]interface{}{"metatitle": "this is my custom title", "posts": viewSpotList}
		r.HTML(200, "spot_detail", newmap)
	})

	m.Get("/main/package/:id",func(args martini.Params, r render.Render) {
		var viewSpotList []Ph_spot
		_, err:= dbmap.Select(&viewSpotList, "select * from ph_view_spots where id=?", args["id"])
		checkErr(err, "Select failed")
		newmap := map[string]interface{}{"metatitle": "this is my custom title", "posts": viewSpotList}


////////////////////////////////////////
//		for i:=0;i < len(iarray);i++ {
//
//			p2 := travelpackagetags(uuid.NewV4().String(), P1.ID,iarray[i])
//			log.Println(p2)
//			err = dbmap.Insert(&p2)
//			checkErr(err, "Insert iamge failed")
//
//		}

			r.HTML(200, "package_detail", newmap)
	})
/***********************************************管理端功能************************************************/
	/**
	景点列表显示
	 */
	m.Get("/admin/viewSpots", func(r render.Render) {
		//fetch all rows
		var viewSpotList []Ph_spot
		_, err:= dbmap.Select(&viewSpotList, "select * from ph_view_spots")
		checkErr(err, "Select failed")
		newmap := map[string]interface{}{"metatitle": "this is my custom title", "posts": viewSpotList}
		r.HTML(200, "view_spot_list", newmap, render.HTMLOptions{
			Layout: "admin_layout",
		})
	})

	/**
	景点详情显示
 	*/
	m.Get("/admin/viewSpots/:id", func(args martini.Params, r render.Render) {
		var spotDetail Ph_spot
		err:= dbmap.SelectOne(&spotDetail, "select * from ph_view_spots where id=?", args["id"])


		if err != nil {
			newmap := map[string]interface{}{"metatitle": "404 Error", "message": "This is not found"}
			r.HTML(404, "error", newmap)
		} else {
			//var spotDetailImages Ph_spot_with_image
			//dbmap.Select(&spotDetailImages, "select * from ph_view_spot_iamges where view_spot_id=?", spotDetail.ID)
			//newImageMap := map[string]interface{}{"metatitle": spotDetail.NAME + " more custom", "spotImages": spotDetailImages}
			newmap := map[string]interface{}{"metatitle": spotDetail.NAME + " more custom", "post": spotDetail}
			r.HTML(200, "view_spot_update", newmap, render.HTMLOptions{
				Layout: "admin_layout",
			})
		}
	})

	/**
	景点修改内容获取
	 */
	m.Get("/admin/spots/:id/update", func(args martini.Params, r render.Render) {
		var spotDetail Ph_spot
		err:= dbmap.SelectOne(&spotDetail, "select * from ph_view_spots where id=?", args["id"])
		//simple error check
		if err != nil {
			newmap := map[string]interface{}{"metatitle": "404 Error", "message": "This is not found"}
			r.HTML(404, "error", newmap)
		} else {
			newmap := map[string]interface{}{"metatitle": spotDetail.NAME + " more custom", "post": spotDetail}
			r.HTML(200, "view_spot_update", newmap, render.HTMLOptions{
				Layout: "admin_layout",
			})
		}
	})

	/**
	景点修改
	 */
	m.Post("/admin/spots/:id/update", func(args martini.Params, r render.Render) {

	})


	/**
	景点新增界面
	 */
	m.Get("/admin/spots/create/",func(r render.Render){
		r.HTML(200, "view_spot_create","",render.HTMLOptions{
			Layout: "admin_layout",
		})
	})

	/**
	景点新增
	 */
	m.Post("/admin/spots/create", binding.Bind(Ph_spot{}), func(spot Ph_spot, args martini.Params, w http.ResponseWriter, r *http.Request) (int, string){
		err := r.ParseMultipartForm(100000)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}



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
			dstSource, err := os.Create("./public/uploads/source/" + u1+"_s."+sourceImageNameExt)
			dstAlbum, err := os.Create("./public/uploads/album/" + u1+"_a."+sourceImageNameExt)
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

			p1 := newViewSpot(uuid.NewV4().String(), spot.COUNTRY, spot.PROVINCE,  spot.CITY,  spot.COUNTY, spot.NAME, spot.LEVEL, spot.LABEL, spot.PRICE, spot.STATUS,"uploads",u1, sourceImageNameExt)
			log.Println(p1)
			err = dbmap.Insert(&p1)
			checkErr(err, "Insert Spot failed")

			image1 := newSpotImage(uuid.NewV4().String(), u1, p1.ID, files[i].Filename, sourceImageNameExt,"uploads")
			log.Println(image1)
			//保存图片信息到图片数据库
			err = dbmap.Insert(&image1)
			checkErr(err, "Insert iamge failed")
		}

		return 200, "ok"

	})

	/**
	套餐管理首页
	 */
	m.Get("/admin/travelPackagelist/", func(r render.Render) {
		//fetch all packages....
		var travelpackagelist []ph_travel_packages
		_, err:= dbmap.Select(&travelpackagelist, "SELECT * FROM ph_travel_packages")
		checkErr(err, "Select failed")
		newmap := map[string]interface{}{"metatitle": "this", "posts": travelpackagelist}
		r.HTML(200, "travel_package_list", newmap, render.HTMLOptions{
			Layout: "admin_layout",
		})

	})

	m.Get("/admin/travelPackagecreate/", func(r render.Render) {
		//fetch all packages....
		//		r.HTML(200, "travel_package_create", "", render.HTMLOptions{
		//			Layout: "admin_layout",
		//		})

		var viewSpotList []Ph_spot
		_, err:= dbmap.Select(&viewSpotList, "SELECT * FROM ph_view_spots")
		checkErr(err, "Select failed")
		newmap := map[string]interface{}{"metatitle": "this", "posts": viewSpotList}
		r.HTML(200, "travel_package_create", newmap, render.HTMLOptions{
			Layout: "admin_layout",
		})

	})

	m.Post("/admin/travelPackagecreate/", binding.Bind(ph_travel_packages{}), func(travel ph_travel_packages, r render.Render) {

		P1 := travelpackage(uuid.NewV4().String(),travel.NAME, travel.DESCRIPTION,  travel.FEE,  travel.START_DATE, travel.END_DATE, travel.DAYS, travel.HOTELS, travel.TRANSPOT, travel.PERSON_NUM, travel.TAGS,travel.CONTENT, travel.ADVICE,travel.FEE_INCLUDE,travel.FEE_NOT_INCLUDE,travel.COLLECTION_ADDRESS,travel.HIGHLIGHTS,travel.VIEW_SPOT_ID)
		log.Println(P1)
		err:= dbmap.Insert(&P1)
		checkErr(err, "Insert failed")
//		return 200, "ok"

//      保存标签
		iarray := strings.Split(P1.TAGS,"，")
		log.Println(iarray)

		for i:=0;i < len(iarray);i++ {

			p2 := travelpackagetags(uuid.NewV4().String(), P1.ID,iarray[i])
			log.Println(p2)
			err = dbmap.Insert(&p2)
			checkErr(err, "Insert iamge failed")

	    }
//      保存亮点
		highlights := strings.Split(P1.HIGHLIGHTS,",")
		log.Println(highlights)

		for i:=0;i < len(highlights);i++ {

			p2 := travelpackagehighlights(uuid.NewV4().String(), P1.ID,highlights[i])
			log.Println(p2)
			err = dbmap.Insert(&p2)
			checkErr(err, "Insert iamge failed")

		}
		//      保存景点
		package_spot_id := strings.Split(P1.VIEW_SPOT_ID,",")
		log.Println(package_spot_id)

		for i:=0;i < len(package_spot_id);i++ {
  			if package_spot_id[i]!=""{
				p2 := travelpackageviewspot(uuid.NewV4().String(), P1.ID, package_spot_id[i], P1.DESCRIPTION, P1.DAYS, "")
				log.Println(p2)
				err = dbmap.Insert(&p2)
				checkErr(err, "Insert iamge failed")
			}

		}


//		newmap := map[string]interface{}{"metatitle": "created post", "travel_package_create": p2}
		r.HTML(200, "main", "", render.HTMLOptions{
			Layout: "admin_layout",
		})
	})

	/**
	套餐详情显示
 	*/
	m.Get("/admin/travelPackagecreate/:id", func(args martini.Params, r render.Render) {
		var travelpackagedetail ph_travel_packages
		err:= dbmap.SelectOne(&travelpackagedetail, "select * from ph_travel_packages where id=?", args["id"])


		if err != nil {
			newmap := map[string]interface{}{"metatitle": "404 Error", "message": "This is not found"}
			r.HTML(404, "error", newmap)
		} else {
			//var spotDetailImages Ph_spot_with_image
			//dbmap.Select(&spotDetailImages, "select * from ph_view_spot_iamges where view_spot_id=?", spotDetail.ID)
			//newImageMap := map[string]interface{}{"metatitle": spotDetail.NAME + " more custom", "spotImages": spotDetailImages}
			newmap := map[string]interface{}{"metatitle": travelpackagedetail.NAME + " more custom", "post": travelpackagedetail}
			r.HTML(200, "travel_package_update", newmap, render.HTMLOptions{
				Layout: "admin_layout",
			})
		}
	})


	/**
	支付管理首页
 	*/
	m.Get("/admin/payments", func(r render.Render) {
		//fetch all payment history....
		r.HTML(200, "payment_list", "", render.HTMLOptions{
			Layout: "admin_layout",
		})
	})


	/*******************************************管理功能结束************************************************/

	//jump to login page
	m.Get("/user/login/",func(r render.Render){
		r.HTML(200, "login","")
//		r.HTML(200, "login","",render.HTMLOptions{
//			Layout: "admin_layout",
//
//		})
	})

	m.Run()
}


func newViewSpot(UUID, COUNTRY, PROVINCE, CITY, COUNTY,NAME,LEVEL,LABEL,PRICE,STATUS ,IMAGEPATH,IMAGENAME,IMAGEFORMAT string ) Ph_spot {
	return Ph_spot{
		ID: UUID,
		COUNTRY:  COUNTRY,
		PROVINCE:   PROVINCE,
		CITY:    CITY,
		COUNTY:COUNTY,
		NAME:NAME,
		LEVEL:LEVEL,
		LABEL:LABEL,
		PRICE:PRICE,
		STATUS:STATUS,
		IMAGEPATH:IMAGEPATH,
		IMAGENAME:IMAGENAME,
		IMAGEFORMAT:IMAGEFORMAT,
	}
}

func travelpackage(ID,NAME,DESCRIPTION string,FEE int64,START_DATE,END_DATE string, DAYS int64, HOTELS,TRANSPOT,PERSON_NUM,TAGS,CONTENT,ADVICE,FEE_INCLUDE,FEE_NOT_INCLUDE,COLLECTION_ADDRESS,HIGHLIGHTS,VIEW_SPOT_ID string ) ph_travel_packages{
	return ph_travel_packages{
		ID:ID,
		NAME:NAME,
		DESCRIPTION:DESCRIPTION,
		FEE:FEE,
		START_DATE:START_DATE,
		END_DATE:END_DATE,
		DAYS:DAYS,
		HOTELS:HOTELS,
		TRANSPOT:TRANSPOT,
		PERSON_NUM:PERSON_NUM,
		TAGS:TAGS,
		CONTENT:CONTENT,
		ADVICE:ADVICE,
		FEE_INCLUDE:FEE_INCLUDE,
		FEE_NOT_INCLUDE:FEE_NOT_INCLUDE,
		COLLECTION_ADDRESS:COLLECTION_ADDRESS,
		HIGHLIGHTS:HIGHLIGHTS,
		VIEW_SPOT_ID:VIEW_SPOT_ID,
	}
}
func travelpackagehighlights(ID,TRAVEL_PACKAGE_ID,CONTENT string) ph_travel_package_highlights{
	return ph_travel_package_highlights{
		ID:ID,
		TRAVEL_PACKAGE_ID:TRAVEL_PACKAGE_ID,
		CONTENT:CONTENT,
	}

}

func travelpackageviewspot(ID,TRAVEL_PACKAGE_ID,VIEW_SPOT_ID,DESCRIPTION string ,DAYS int64,PLAY_ITEMS string) ph_travel_package_view_spots{
	return ph_travel_package_view_spots{
		ID:ID,
		TRAVEL_PACKAGES_ID:TRAVEL_PACKAGE_ID,
		VIEW_SPOT_ID:VIEW_SPOT_ID,
		DESCRIPTION:DESCRIPTION,
		DAYS:DAYS,
		PLAY_ITEMS:PLAY_ITEMS,
	}
}
func travelpackagetags(ID, TRAVEL_PACKAGE_ID, CONTENT string) ph_travel_package_tags {
	return ph_travel_package_tags{
		ID:ID,
		TRAVEL_PACKAGE_ID:TRAVEL_PACKAGE_ID,
		CONTENT:CONTENT,

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
//	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/kanghui")
	db, err := sql.Open("mysql", "cdb_outerroot:cems@2015#@tcp(55c9b23419e13.sh.cdb.myqcloud.com:4592)/kanghui")

	checkErr(err, "sql.Open failed")


	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Ph_spot{}, "ph_view_spots")
	dbmap.AddTableWithName(Ph_spot_with_image{}, "ph_view_spot_images")
	dbmap.AddTableWithName(ph_travel_packages{}, "ph_travel_packages")
	dbmap.AddTableWithName(ph_travel_package_highlights{}, "ph_travel_package_highlights")
	dbmap.AddTableWithName(ph_travel_package_tags{}, "ph_travel_package_tags")
	dbmap.AddTableWithName(ph_travel_package_view_spots{}, "ph_travel_package_view_spots")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}


