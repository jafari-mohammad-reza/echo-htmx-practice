package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func GetBlogs(c echo.Context) error {
	var blogs []Blog
	rows, err := db.Query(`SELECT * FROM Blog`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()
	for rows.Next() {
		var b Blog
		err := rows.Scan(&b.Id, &b.Title)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		blogs = append(blogs, b)
	}
	if err = rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string][]Blog{"Blogs": blogs})
}

func GetBlog(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"Path": "get blog",
	})
}
func CreateBlog(c echo.Context) error {
	return c.JSON(201, map[string]string{
		"Path": "Create blog",
	})
}
func DeleteBlog(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"Path": "delete blog",
	})
}
