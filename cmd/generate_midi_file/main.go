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
	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	app := &application{}
	os.Exit(service.Main(context.Background(), app, &app.SentryDSN, &app.SentryProxy))
}

type application struct {
	SentryDSN   string `required:"false" arg:"sentry-dsn"   env:"SENTRY_DSN"   usage:"SentryDSN"             display:"length"`
	SentryProxy string `required:"false" arg:"sentry-proxy" env:"SENTRY_PROXY" usage:"Sentry Proxy"`
	OutputFile  string `required:"true"  arg:"output"       env:"OUTPUT"       usage:"Output MIDI file path"                  default:"output.mid"`
}

func (a *application) Run(ctx context.Context, sentryClient libsentry.Client) error {
	// Create a new Standard MIDI File (SMF)
	s := smf.New()

	// Create a track
	var track smf.Track

	// Add a program change (set instrument to Acoustic Grand Piano)
	track.Add(0, midi.ProgramChange(0, 0)) // delta ticks 0, channel 0, program 0

	// Define a simple melody: C4, D4, E4, G4
	melody := []uint8{60, 62, 64, 67}
	// Define the duration for each note (in ticks)
	duration := uint32(480) // 480 ticks per quarter note

	// Add note on and note off events for each note in the melody
	for _, note := range melody {
		track.Add(0, midi.NoteOn(0, note, 100))    // delta 0, channel 0, velocity 100
		track.Add(duration, midi.NoteOff(0, note)) // delta duration, channel 0
	}

	// Close the track
	track.Close(0)

	// Add track to SMF
	if err := s.Add(track); err != nil {
		return errors.Wrapf(ctx, err, "add track to SMF failed")
	}

	// Write the SMF to file
	if err := s.WriteFile(a.OutputFile); err != nil {
		return errors.Wrapf(ctx, err, "write MIDI file %s failed", a.OutputFile)
	}

	return nil
}
