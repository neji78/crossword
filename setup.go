func main() {
	loadConfig()
	initDB()
	initPuzzleTable()

	app := fiber.New()

	setupRoutes(app)

	// Run server
	app.Listen(":" + viper.GetString("server.port"))
}
