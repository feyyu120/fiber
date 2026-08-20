fmt.Println("server is running")
	app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true})