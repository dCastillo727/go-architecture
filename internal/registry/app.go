package registry

type App struct {
	Services Services
}

func NewApp(services Services) *App {
	return &App{
		Services: services,
	}
}
