# Sample MIDI

Sample MIDI project demonstrating MIDI operations in Go.

## Features

- **Send MIDI messages** in real-time to MIDI output ports
- **Generate MIDI files** (.mid format) with custom melodies
- **Play MIDI files** by sending them to MIDI output ports

## Prerequisites

### macOS: Enable IAC Driver

1. Open Audio MIDI Setup
2. Window → Show MIDI Studio
3. Double-click IAC Driver icon
4. Check "Device is online"
5. Create a port named "GoMIDI"

## Usage

### Send MIDI Messages

```bash
cd cmd/send_midi
make run
```

Options:
- `--port`: MIDI port name (default: "IAC Driver GoMIDI")

### Generate MIDI File

```bash
cd cmd/generate_midi_file
make run
```

Options:
- `--output`: Output file path (default: "output.mid")

### Send MIDI File

```bash
cd cmd/send_midi_file
make run --input output.mid
```

Options:
- `--port`: MIDI port name (default: "IAC Driver GoMIDI")
- `--input`: Input MIDI file path (required)

## Development

```bash
make precommit  # Run all checks
make test       # Run tests
make format     # Format code
```
