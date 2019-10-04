package https

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TonyChouteau/multi-shot-server/manager"
)

func connect(c *gin.Context) {
	player := manager.CreatePlayer()

	c.JSON(200, player)
}

func refresh(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	x, _ := strconv.ParseFloat(c.Param("x"), 64)
	y, _ := strconv.ParseFloat(c.Param("y"), 64)

	others, count := manager.Refresh(id, x, y)

	c.JSON(200, struct {
		Others manager.PlayerList `json:"others"`
		Count  int                `json:"count"`
	}{
		others,
		count,
	})
}

/*
Serve function
*/
func Serve() {
	r := gin.Default()
	r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://www.tonychouteau.fr"},
	}))

	r.GET("/connect", connect)
	r.GET("/refresh/:id/:x/:y", refresh)

	err := http.ListenAndServe(":8084", r)
	//err := http.ListenAndServeTLS(":8084", "/etc/letsencrypt/live/www.domain.com/fullchain.pem", "/etc/letsencrypt/live/www.domain.com/privkey.pem", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
