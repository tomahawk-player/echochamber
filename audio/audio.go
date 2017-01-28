package audio

import "io"

// An Engine can be fed a multimedia file and it will read it, demux it, decode it and play it.
type Engine interface {
	Play(fileStream io.ReadCloser) error
	/* TODO: add more methods, such as volume and such */
}

// An Pipeline is a more tunable AudioEngine.
type Pipeline interface {
	Engine
	Output() Output
	SetOutput(Output) error
}

// An Output is an abstraction over an audio output device.
type Output interface {
	/* TODO: add methods */
}
