<p align="center">
  <a href="voiceai-go"><img src="voiceai.png" width="420" border="0" alt="voiceai-go"></a>
</p>
<p align="center">
  <a href="https://opensource.org/license/0bsd"><img src="https://img.shields.io/badge/License-0BSD-orange.svg?" alt="License"></a>
  <a href="https://github.com/0x1eef/voiceai-go/tags"><img src="https://img.shields.io/badge/version-0.1.0-green.svg?" alt="Version"></a>
</p>

## About


voiceai-go provides **unofficial** Go bindings for
[voice.ai](https://voice.ai)'s REST API &ndash; which supports features
such as text-to-speech, voice replication, and agent management. The README
focuses on voice replication and text-to-speech. It includes practical examples
that can be run via the repository's [cmd](cmd) directory.

## Quick start

#### voiceai.Client

The Client acts as the main gateway to the [voice.ai](https://voice.ai)
API, and its methods provide access to different REST endpoints. The `settings`
package is a recurring pattern in voiceai-go, and it provides a way to configure
the client and its different endpoints:

```go
package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
}
```

#### Voice.Clone

Voice provides voice replication capabilities. The `Clone` method accepts
the name of the voice to be created and a path to an audio file that contains
the voice to be replicated. The following example creates a voice named "Trebor"
and it uses the audio file located at [share/inputs/trebor.mp3](share/inputs/trebor.mp3).
The audio is computer-generated and intentionally not based on a real person:

```go
package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
	"github.com/0x1eef/voiceai/settings/voice"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	clonedVoice, err := client.Voice().Clone(
		voice.WithName("Trebor"),
		voice.WithPath("share/inputs/trebor.mp3"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID:          %s\n", clonedVoice.ID)
	fmt.Printf("Name:        %s\n", clonedVoice.Name)
	fmt.Printf("Status:      %s\n", clonedVoice.Status)
	fmt.Printf("Visibility:  %s\n", clonedVoice.Visibility)
}
```

#### Speech.Create

Speech provides text-to-speech (TTS) capabilities. This method accepts
a string of text and [various other options](https://voice.ai/docs/api-reference/text-to-speech/generate-speech).
It performs a blocking POST request and eventually returns an [io.ReadCloser](https://pkg.go.dev/io#ReadCloser)
that contains the full response body. The following example produces
[share/outputs/trebor.mp3](share/outputs/trebor.mp3) and it uses the voice
that was replicated in the previous example:

```go
package main

import (
	"io"
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
	"github.com/0x1eef/voiceai/settings/speech"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	audio, err := client.Speech().Create(
		speech.WithText("Hello! My name is Trebor"),
		speech.WithVoiceID("trebors_voice_id"),
		speech.WithFormat("mp3"),
	)
	if err != nil {
		panic(err)
	}
	defer audio.Close()
	out, _ := os.Create("share/outputs/trebor.mp3")
	defer out.Close()
	io.Copy(out, audio)
}
```

#### Speech.Stream

Speech supports streaming through the "Stream" method and is otherwise identical to
the "Create" method. It returns an [io.ReadCloser](https://pkg.go.dev/io#ReadCloser)
that can be read from as the audio is being produced by the server. This example
produces the same output as the previous example, and appears identical
but the audio is written to the file as it is streamed from the server:

```go
package main

import (
	"io"
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
	"github.com/0x1eef/voiceai/settings/speech"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	stream, err := client.Speech().Stream(
		speech.WithText("Hello! My name is Trebor"),
		speech.WithVoiceID("trebors_voice_id"),
		speech.WithFormat("mp3"),
	)
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	out, _ := os.Create("share/outputs/trebor.mp3")
	defer out.Close()
	io.Copy(out, stream)
}
```

#### Agent.{Create,Deploy}

Agent provides agent management capabilities. The "Create" method accepts the name of
the agent to be created and [various other options](https://voice.ai/docs/api-reference/agent-management/create-agent),
including a prompt that describes the agent's behavior:

```go
package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
	"github.com/0x1eef/voiceai/settings/agent"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken(os.Getenv("KEY")),
	)
	if err != nil {
		panic(err)
	}
	agent, err := client.Agent().Create(
		agent.WithName("Trebor"),
		agent.WithPrompt("Trebor is a helpful assistant"),
	)
	if err != nil {
		panic(err)
	}
	if agent, err := agent.Deploy(); err != nil {
		panic(err)
	} else {
		fmt.Printf("Agent deployed: %v\n", agent)
	}
}
```

## API Index

The Quick Start section above provides examples for the most interesting methods
in the library, but there is support for more than that:

* [speech.go](speech.go)
	* Speech.Create
	* Speech.Stream

* [voice.go](voice.go)
	* Voice.All
	* Voice.Get
	* Voice.Delete
	* Voice.Update
	* Voice.Clone

* [agent.go](agent.go)
	* Agent.All
	* Agent.Create
	* Agent.Deploy
	* Agent.Pause
	* Agent.Disable

## Resources

**REST API**

* [Documentation](https://voice.ai/docs/)
* [API Reference](https://voice.ai/docs/api-reference/)

**Library API**

This library implements patterns I have found and used in other libraries,
including my own, and I have documented some of them on my blog:

* [The functional options pattern in Go](https://0x1eef.github.io/posts/the-functional-options-pattern/)
* [How to make HTTP requests in Go](https://0x1eef.github.io/posts/how-to-make-http-requests-in-go)
* [How to stream a response body in Go](https://0x1eef.github.io/posts/how-to-stream-a-response-body-in-go)
* [How to parse JSON in Go](https://0x1eef.github.io/posts/how-to-parse-json-in-go)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
