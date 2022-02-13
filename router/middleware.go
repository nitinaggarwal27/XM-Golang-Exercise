package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"nitinaggarwal27/XM-Golang-Exercise/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func contains(data []string, search string) bool {
	for _, value := range data {
		if value == search {
			return true
		}
	}
	return false
}

func checkLocation(c *gin.Context) {
	url := "https://ipapi.co/" + c.ClientIP() + "/json/"
	log.Println(url)
	method := "GET"

	payload := strings.NewReader(`{}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		c.Abort()
		c.JSON(500, gin.H{"error": true, "message": "Please try again."})
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		c.Abort()
		c.JSON(500, gin.H{"error": true, "message": "Please try again."})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		c.Abort()
		c.JSON(500, gin.H{"error": true, "message": "Please try again."})
		return
	}
	fmt.Println(string(body))

	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		c.Abort()
		c.JSON(500, gin.H{"error": true, "message": "Please try again."})
		return
	}
	reserved, ok := data["reserved"]
	if ok && reserved.(bool) {
		c.Next()
		return
	}
	country, ok := data["country_name"]
	if ok {
		if !contains(strings.Split(config.Conf.Service.ValidLocations, ","), country.(string)) {
			c.Abort()
			c.JSON(403, gin.H{"error": true, "message": "You are not allowed"})
			return
		}
	} else {
		c.Abort()
		c.JSON(500, gin.H{"error": true, "message": "Please try again"})
		return
	}
	c.Next()

}
