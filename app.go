package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// // inject event emitter ke handler
	// handler.EmitEvent = func(name string, data interface{}) {
	// 	runtime.EventsEmit(a.ctx, name, data)
	// }

	// go func() {
	// 	http.HandleFunc("/api/gate/", handler.GateHandler)

	// 	log.Println("HTTP Server running at :8000")
	// 	http.ListenAndServe(":8000", nil)
	// }()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) TriggerGate(data string) string {
	// call API internal
	return "ok"
}

func (a *App) Hello(name string) string {
	return "Hello " + name
}

// func (a *App) GetNetwork() (service.NetworkInfo, error) {
// 	return service.GetNetworkInfo()
// }

// func (a *App) SetNetwork(ip, subnet, gateway string) string {
// 	err := service.SetNetworkInfo(ip, subnet, gateway)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return "OK"
// }
