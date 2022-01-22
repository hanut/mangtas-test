package main

import (
	"bytes"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// Enable debug mode. For a more robust solution
	// this should be configured via the env and should be
	// gin.ReleaseMode for production.
	gin.SetMode(gin.DebugMode)

	// Setup the routes
	r := setupRoutes()

	// Start the server
	r.Run("localhost:3000")
}

// Setup routes and their handlers
func setupRoutes() *gin.Engine {
	r := gin.Default()

	// Don't trust any proxies, this is a simple service !
	r.SetTrustedProxies(nil)

	// Redirect all requests to the default '/' path
	r.POST("/", getTopTen)

	return r
}

// Struct defining the type of Top Ten Response objects
// in the response array from the service
type WordCountResult struct {
	Word  string
	Count uint16
}

// Validate and process new requests before returning the
// expected json response
func getTopTen(c *gin.Context) {
	// Read the bytes of the input into a buffer
	// and respond with an appropriate server error
	// incase the parsing fails.
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	// Remove all unnecessary characters like fullstops
	// commas etc except the hyphens
	reg, err := regexp.Compile("[.,:\"]+")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	// Split the string into words after converting the input to lowercase
	// and replacing all unnecessary characters with a space
	words := strings.Fields(reg.ReplaceAllString(strings.ToLower(buf.String()), " "))

	// Setup a map to track the individual word counts
	countMap := make(map[string]*WordCountResult)
	for _, word := range words {
		_, exists := countMap[word]
		if !exists {
			countMap[word] = &WordCountResult{
				Word:  word,
				Count: 1,
			}
			continue
		}
		countMap[word].Count++
	}
	// convert the map to array for sorting
	var countArr []*WordCountResult
	for _, item := range countMap {
		countArr = append(countArr, item)
	}
	sort.Sort(sort.Reverse(ByCount{countArr}))
	// replace old slice with a subslice of the first 10 items
	countArr = countArr[:10]
	c.JSON(200, countArr)
}
