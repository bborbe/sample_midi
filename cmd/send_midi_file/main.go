// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"

	"github.com/bborbe/errors"
	libsentry "github.com/bborbe/sentry"
	"github.com/bborbe/service"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	app := &application{}
	os.Exit(service.Main(context.Background(), app, &app.SentryDSN, &app.SentryProxy))
}

type application struct {
	SentryDSN   string `required:"false" arg:"sentry-dsn"   env:"SENTRY_DSN"   usage:"SentryDSN"         display:"length"`
	SentryProxy string `required:"false" arg:"sentry-proxy" env:"SENTRY_PROXY" usage:"Sentry Proxy"`
	Port        string `required:"true"  arg:"port"         env:"PORT"         usage:"MIDI port name"                     default:"IAC Driver GoMIDI"`
	FilePath    string `required:"true"  arg:"file"        env:"FILE"        usage:"Input MIDI file path"`
}

func (a *application) Run(ctx context.Context, sentryClient libsentry.Client) error {
	defer midi.CloseDriver()

	// Find output port
	out, err := midi.FindOutPort(a.Port)
	if err != nil {
		return errors.Wrapf(ctx, err, "find MIDI output port %s failed", a.Port)
	}

	// Read and play MIDI file with proper timing
	rd := smf.ReadTracks(a.FilePath)
	if err := rd.Play(out); err != nil {
		return errors.Wrapf(ctx, err, "play MIDI file failed")
	}

	return nil
}
