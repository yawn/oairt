// go:build example
package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gordonklaus/portaudio"
	"github.com/pkg/errors"
	"github.com/yawn/oairt"
	"github.com/yawn/oairt/types"
)

type Audio struct {
	lock    sync.Mutex
	OnInput func(data []byte)
	out     []float32
	stream  *portaudio.Stream
}

func (a *Audio) Start() error {

	if err := portaudio.Initialize(); err != nil {
		return errors.Wrapf(err, "failed to initialize portaudio")
	}

	stream, err := portaudio.OpenDefaultStream(1, 1, 24000, 1024, func(input, output []float32) {

		var out = make([]float32, 1024)

		a.lock.Lock()

		n := copy(out, a.out)

		a.out = a.out[n:]
		copy(output, out)

		var in []byte

		for _, e := range input {
			sample := int16(e * 32767)
			in = append(in, byte(sample), byte(sample>>8))
		}

		a.OnInput(in)

		a.lock.Unlock()

	})

	if err != nil {
		return errors.Wrapf(err, "failed to open default stream")
	}

	if err := stream.Start(); err != nil {
		return errors.Wrapf(err, "failed to start audio stream")
	}

	a.stream = stream

	return nil

}

func (a *Audio) Close() error {

	if err := a.stream.Close(); err != nil {
		return errors.Wrapf(err, "failed to close audio stream")
	}

	return portaudio.Terminate()

}

func (a *Audio) AddOutput(data []byte) {

	samples := make([]float32, len(data)/2)

	for i := 0; i < len(data); i += 2 {
		sample := int16(data[i]) | int16(data[i+1])<<8
		samples[i/2] = float32(sample) / 32767.0
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	a.out = append(a.out, samples...)

}

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

	client.AddHandler(&oairt.Handler[*types.ServerSessionCreated]{
		Type: types.TypeServerSessionCreated,
		ID:   "session-management",
		Handle: func(event *types.ServerSessionCreated) (bool, error) {

			slog.Info("session created")

			if err := client.Send(ctx, &types.ClientSessionUpdate{
				Session: types.ClientSession{
					Voice: types.String("shimmer"),
				},
			}); err != nil {
				return false, errors.Wrapf(err, "failed to send session update")
			}

			audio := new(Audio)

			if err := audio.Start(); err != nil {
				return false, errors.Wrapf(err, "failed to start audio")
			}

			audio.OnInput = func(data []byte) {

				if err := client.Send(ctx, &types.ClientInputAudioBufferAppend{
					Audio: base64.StdEncoding.EncodeToString(data),
				}); err != nil {
					cancel(errors.Wrapf(err, "failed to send audio buffer append"))
				}

			}

			client.AddHandler(&oairt.Handler[*types.ServerResponseAudioDelta]{
				Type: types.TypeServerResponseAudioDelta,
				ID:   "audio-decoder",
				Handle: func(event *types.ServerResponseAudioDelta) (bool, error) {

					data, err := base64.StdEncoding.DecodeString(event.Delta)

					if err != nil {
						return false, errors.Wrapf(err, "failed to decode audio delta")
					}

					audio.AddOutput(data)

					return false, nil

				},
			})

			slog.Info("audio started, handlers added")

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
