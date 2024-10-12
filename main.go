package main

import (
	"context"
	"github.com/robinbraemer/event"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"mcauth/codes"
	"strconv"
)

func main() {
	go setupHttpServer()

	proxy.Plugins = append(proxy.Plugins, proxy.Plugin{
		Name: "MC Auth",
		Init: func(ctx context.Context, proxy *proxy.Proxy) error {
			event.Subscribe(proxy.Event(), 0, onLogin)

			return nil
		},
	})

	gate.Execute()
}

func onLogin(e *proxy.PostLoginEvent) {
	code := codes.New(e.Player().ID())

	if code == -1 {
		e.Player().Disconnect(&component.Text{
			Content: "Failed to generate a verification code, try again later",
			S: component.Style{
				Color: color.Red,
			},
		})

		return
	}

	e.Player().Disconnect(&component.Text{
		Content: "Your one time verification code:\n\n",
		Extra: []component.Component{
			&component.Text{
				Content: strconv.Itoa(code),
				S: component.Style{
					Bold:       component.True,
					Underlined: component.True,
				},
			},
			&component.Text{
				Content: "\n\nThe code will be valid for the next 5 minutes",
				S: component.Style{
					Color: color.DarkGray,
				},
			},
		},
	})
}
