package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi/config"
	"restapi/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Koneksi ke MySQL Berhasil")

	route := echo.New()

	// Member API Start
	// Saya sudah mencoba untuk get All data dengan query SELECT * FROM member tapi dari rules db.query echo tidak bisa untuk mendapatkan 2 values dalam 1 query
	// Dan ketika dicoba pun masih mendapatakan 1 values saja
	route.GET("member/:ID_MEMBER", func(c echo.Context) error {
		requested_id := c.Param("ID_MEMBER")
		fmt.Println(requested_id)
		var ID int
		var USERNAME string
		var GENDER string
		var SKINTYPE string
		var SKINCOLOR string

		queryText := db.QueryRow("SELECT ID_MEMBER,USERNAME, GENDER, SKINTYPE, SKINCOLOR FROM member WHERE ID_MEMBER = ?", requested_id).Scan(&ID, &USERNAME, &GENDER, &SKINTYPE, &SKINCOLOR)
		if queryText != nil {
			fmt.Println(queryText)
		}
		fmt.Println("Sukses mengambil data dari database")
		fmt.Println(models.Member{ID: ID, USERNAME: USERNAME, GENDER: GENDER, SKINTYPE: SKINTYPE, SKINCOLOR: SKINCOLOR})
		response := models.Member{ID: ID, USERNAME: USERNAME, GENDER: GENDER, SKINTYPE: SKINTYPE, SKINCOLOR: SKINCOLOR}
		return c.JSON(http.StatusOK, response)
	})

	route.POST("member/create_member", func(c echo.Context) error {
		user := new(models.Member)
		c.Bind(user)
		contentType := c.Request().Header.Get("Content-type")
		if contentType == "application/json" {
			fmt.Println("Request dari json")
			fmt.Println(*user)
			queryText := ("INSERT INTO member (USERNAME,GENDER,SKINTYPE,SKINCOLOR) values(?,?,?,?)")
			stmt, err := db.Prepare(queryText)
			if err != nil {
				fmt.Print(err.Error())
			}
			defer stmt.Close()
			result, err2 := stmt.Exec(user.USERNAME,
				user.GENDER,
				user.SKINTYPE,
				user.SKINCOLOR)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println(result.LastInsertId())
			fmt.Println("Sukses melakukan penambahan data")
			fmt.Println(*user)
		}
		response := struct {
			Message string
			Data    models.Member
		}{
			Message: "Sukses melakukan penambahan data",
			Data:    *user,
		}
		return c.JSON(http.StatusOK, response)
	})

	// Update data sudah bisa menerima input body dan response namun masih belum terupdate pada database nya
	route.PUT("member/update_member/:ID_MEMBER", func(c echo.Context) error {
		user := new(models.Member)
		c.Bind(user)
		requested_id := c.Param("ID_MEMBER")
		fmt.Println(requested_id)

		queryText := db.QueryRow("UPDATE member SET USERNAME = ?, GENDER = ?, SKINTYPE = ?, SKINCOLOR = ? WHERE member.ID_MEMBER = ?", requested_id, user.USERNAME, user.GENDER, user.SKINTYPE, user.SKINCOLOR)
		if queryText != nil {
			fmt.Println(queryText)
		}
		fmt.Println("Sukses melakukan perubahan data")
		fmt.Println(*user)
		response := struct {
			Message string
			Data    models.Member
		}{
			Message: "Sukses melakukan perubahan data",
			Data:    *user,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("member/delete_member/:ID_MEMBER", func(c echo.Context) error {
		requested_id := c.Param("ID_MEMBER")
		queryText := db.QueryRow("DELETE FROM member WHERE ID_MEMBER = ?", requested_id)
		if queryText != nil {
			fmt.Println(queryText)
		}
		fmt.Println("Sukses menghapus data dari database")
		response := "Sukses menghapus data dari database"
		return c.JSON(http.StatusOK, response)
	})
	// Member API End

	// Product API Start
	route.GET("product/:ID_PRODUCT", func(c echo.Context) error {
		requested_id := c.Param("ID_PRODUCT")
		fmt.Println(requested_id)
		var ID int
		var NAME_PRODUCT string
		var PRICE string

		queryText := db.QueryRow("SELECT ID_PRODUCT,NAME_PRODUCT, PRICE FROM product WHERE ID_PRODUCT = ?", requested_id).Scan(&ID, &NAME_PRODUCT, &PRICE)
		if queryText != nil {
			fmt.Println(queryText)
		}
		fmt.Println("Sukses mengambil data dari database")
		fmt.Println(models.Product{ID: ID, NAME_PRODUCT: NAME_PRODUCT, PRICE: PRICE})
		response := models.Product{ID: ID, NAME_PRODUCT: NAME_PRODUCT, PRICE: PRICE}
		return c.JSON(http.StatusOK, response)
	})
	// Product API End

	// Review Product API Start
	route.GET("review_product/:ID_REVIEW", func(c echo.Context) error {
		requested_id := c.Param("ID_REVIEW")
		fmt.Println(requested_id)
		var ID int
		var ID_PRODUCT int
		var ID_MEMBER int
		var DESC_REVIEW string
		var JUMLAH_LIKE_REVIEW int

		queryText := db.QueryRow("SELECT ID_REVIEW,ID_PRODUCT, ID_MEMBER, DESC_REVIEW, JUMLAH_LIKE_REVIEW FROM review_product WHERE ID_REVIEW = ?", requested_id).Scan(&ID, &ID_PRODUCT, &ID_MEMBER, &DESC_REVIEW, &JUMLAH_LIKE_REVIEW)
		if queryText != nil {
			fmt.Println(queryText)
		}
		fmt.Println("Sukses mengambil data dari database")
		fmt.Println(models.Review{ID: ID, ID_PRODUCT: ID_PRODUCT, ID_MEMBER: ID_MEMBER, DESC_REVIEW: DESC_REVIEW, JUMLAH_LIKE_REVIEW: JUMLAH_LIKE_REVIEW})
		response := models.Review{ID: ID, ID_PRODUCT: ID_PRODUCT, ID_MEMBER: ID_MEMBER, DESC_REVIEW: DESC_REVIEW, JUMLAH_LIKE_REVIEW: JUMLAH_LIKE_REVIEW}
		return c.JSON(http.StatusOK, response)
	})

	route.POST("review_product/create_review_product", func(c echo.Context) error {
		review := new(models.Review)
		c.Bind(review)
		contentType := c.Request().Header.Get("Content-type")
		if contentType == "application/json" {
			fmt.Println("Request dari json")
			fmt.Println(*review)
			queryText := ("INSERT INTO review_product (ID_PRODUCT,ID_MEMBER,DESC_REVIEW) values(?,?,?)")
			stmt, err := db.Prepare(queryText)
			if err != nil {
				fmt.Print(err.Error())
			}
			defer stmt.Close()
			result, err2 := stmt.Exec(review.ID_PRODUCT,
				review.ID_MEMBER,
				review.DESC_REVIEW)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println(result.LastInsertId())
			fmt.Println("Sukses melakukan penambahan data")
			fmt.Println(*review)
		}
		response := struct {
			Message string
			Data    models.Review
		}{
			Message: "Sukses melakukan penambahan data",
			Data:    *review,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.POST("review_product/like_product", func(c echo.Context) error {
		like := new(models.Like_Review)
		c.Bind(like)
		contentType := c.Request().Header.Get("Content-type")
		if contentType == "application/json" {
			fmt.Println("Request dari json")
			fmt.Println(*like)
			queryText := ("INSERT INTO like_review (ID_REVIEW,ID_MEMBER) values(?,?)")
			stmt, err := db.Prepare(queryText)
			if err != nil {
				fmt.Print(err.Error())
			}
			defer stmt.Close()
			result, err2 := stmt.Exec(like.ID_REVIEW,
				like.ID_MEMBER)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println(result.LastInsertId())
			fmt.Println("Sukses melakukan penambahan data")
			fmt.Println(*like)
		}
		response := struct {
			Message string
			Data    models.Like_Review
		}{
			Message: "Sukses melakukan penambahan data",
			Data:    *like,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("review_product/dislike_product/:ID_REVIEW", func(c echo.Context) error {
		requested_id := c.Param("ID_REVIEW")
		queryText := db.QueryRow("DELETE FROM like_review WHERE ID_REVIEW = ?", requested_id)
		if queryText != nil {
			fmt.Println(queryText)
		}
		fmt.Println("Sukses menghapus data dari database")
		response := "Sukses menghapus data dari database"
		return c.JSON(http.StatusOK, response)
	})
	// Review Product API End

	route.Start(":7000")
}
