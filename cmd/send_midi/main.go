// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"time"

	"github.com/bborbe/errors"
	libsentry "github.com/bborbe/sentry"
	"github.com/bborbe/service"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

func main() {
	app := &application{}
	os.Exit(service.Main(context.Background(), app, &app.SentryDSN, &app.SentryProxy))
}

type application struct {
	SentryDSN   string `required:"false" arg:"sentry-dsn"   env:"SENTRY_DSN"   usage:"SentryDSN"      display:"length"`
	SentryProxy string `required:"false" arg:"sentry-proxy" env:"SENTRY_PROXY" usage:"Sentry Proxy"`
	Port        string `required:"true"  arg:"port"         env:"PORT"         usage:"MIDI port name"                  default:"IAC Driver GoMIDI"`
}

func (a *application) Run(ctx context.Context, sentryClient libsentry.Client) error {
	defer midi.CloseDriver()

	out, err := midi.FindOutPort(a.Port)
	if err != nil {
		return errors.Wrapf(ctx, err, "find MIDI output port %s failed", a.Port)
	}

	send, err := midi.SendTo(out)
	if err != nil {
		return errors.Wrapf(ctx, err, "create MIDI sender failed")
	}

	// Send some notes
	if err := send(midi.NoteOn(0, 60, 100)); err != nil {
		return errors.Wrapf(ctx, err, "send note on failed")
	}
	time.Sleep(500 * time.Millisecond)
	if err := send(midi.NoteOff(0, 60)); err != nil {
		return errors.Wrapf(ctx, err, "send note off failed")
	}

	if err := send(midi.NoteOn(0, 64, 100)); err != nil {
		return errors.Wrapf(ctx, err, "send note on failed")
	}
	time.Sleep(500 * time.Millisecond)
	if err := send(midi.NoteOff(0, 64)); err != nil {
		return errors.Wrapf(ctx, err, "send note off failed")
	}

	return nil
}
