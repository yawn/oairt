package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/yawn/oairt"
	"github.com/yawn/oairt/types"
)

func _main(ctx context.Context) error {

	logOptions := &slog.HandlerOptions{}

	if os.Getenv("DEBUG") == "true" {
		logOptions.Level = slog.LevelDebug.Level()
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, logOptions)))

	ctx, _ = signal.NotifyContext(ctx, syscall.SIGTERM)

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY is not set")
	}

	ctx, cancel := context.WithCancelCause(ctx)

	client := oairt.New(apiKey)

	client.AddHandler(&oairt.Handler[*types.ServerError]{
		Type: types.TypeServerError,
		Handle: func(event *types.ServerError) (bool, error) {
			return false, event
		},
	})

	client.AddHandler(&oairt.Handler[*types.ServerSessionCreated]{
		Type: types.TypeServerSessionCreated,
		ID:   "session-management",
		Handle: func(event *types.ServerSessionCreated) (bool, error) {

			slog.Info("session created")

			if err := client.Send(ctx, &types.ClientResponseCreate{
				Response: &types.ClientResponse{
					Modalities:   []string{"text"},
					Instructions: "Please assist the user by telling them a very long joke.",
				},
			}); err != nil {
				return false, err
			}

			client.AddHandler(&oairt.Handler[*types.ServerResponseTextDelta]{
				Type: types.TypeServerResponseTextDelta,
				Handle: func(event *types.ServerResponseTextDelta) (bool, error) {
					fmt.Printf("%s", event.Delta)
					return true, nil

				},
			})

			client.AddHandler(&oairt.Handler[*types.ServerResponseTextDone]{
				Type: types.TypeServerResponseTextDone,
				Handle: func(event *types.ServerResponseTextDone) (bool, error) {

					fmt.Println("\n")

					cancel(fmt.Errorf("done"))

					return true, nil

				},
			})

			return false, nil

		},
	})

	if err := client.Start(ctx); err != nil {
		return err
	}

	return nil

}

func main() {

	if os.Getenv("DEBUG") != "true" {
		log.SetOutput(io.Discard)
	}

	if err := _main(context.Background()); err != nil {

		if err.Error() != "done" {
			panic(err)
		}

	}

}
