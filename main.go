package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	anime "github.com/ani-social/pipebomb/anime"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
	}))

	scraper := anime.NewAnimeScraper()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"intro":         "Welcome to pip-api. UWU",
			"documentation": "http://pipebomb.bytecats.codes:8000/docs",
		})
	})

	app.Post("/anime/search_anime", func(c *fiber.Ctx) error {
		var animeSearch anime.AnimeSearch
		if err := c.BodyParser(&animeSearch); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		query := strings.ReplaceAll(animeSearch.Query, " ", "_")
		cacheKey := "anime_search_" + query

		if cachedData, err := redisClient.Get(c.Context(), cacheKey).Result(); err == nil {
			var searchResults []anime.AnimeSearchResult
			if err := json.Unmarshal([]byte(cachedData), &searchResults); err == nil {
				return c.JSON(searchResults)
			}
		}

		searchResults, err := scraper.SearchAnime(animeSearch.Query)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		data, _ := json.Marshal(searchResults)
		redisClient.Set(c.Context(), cacheKey, data, 10*time.Minute)
		return c.JSON(searchResults)
	})

	// Implement the remaining endpoints similar to the above example, using the appropriate
	// functions and data structures from the anime, lightnovels, and movies packages.

	app.Listen(":8000")
}
