<p align="center">
  <a href="voiceai-go"><img src="voiceai.png" width="420" border="0" alt="voiceai-go"></a>
</p>
<p align="center">
  <a href="https://opensource.org/license/0bsd"><img src="https://img.shields.io/badge/License-0BSD-orange.svg?" alt="License"></a>
  <a href="https://github.com/0x1eef/voiceai-go/tags"><img src="https://img.shields.io/badge/version-0.1.0-green.svg?" alt="Version"></a>
</p>

## About


voiceai-go provides **unofficial** Go bindings for
[voice.ai](https://voice.ai)'s REST API &ndash; which supports
a lot of cool features like text-to-speech, voice replication and
much more. The README focuses on voice replication and text-to-speech.
It includes practical examples that can be run via the repository's
[cmd](cmd) directory.

### Quick start

#### voiceai.Client

The Client acts as the main gateway to the [voice.ai](https://voice.ai)
API, and its methods provide access to different REST endpoints. The `settings`
package is a recurring pattern in voiceai-go, and it provides a way to configure
the client and its different endpoints:

```go
package main

import (
	"fmt"

	"github.com/0x1eef/voiceai"
	"github.com/0x1eef/voiceai/settings"
)

func main() {
	client, err := voiceai.NewClient(
		settings.WithToken("yourtoken"),
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
	res, err := client.Voice().Clone(
		voice.WithName("Trebor"),
		voice.WithPath("share/inputs/trebor.mp3"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", res)
}
```

#### Speech.Create

Speech provides text-to-speech (TTS) capabilities. The `Create` method accepts
a string of text and [various other options](https://voice.ai/docs/api-reference/text-to-speech/generate-speech)
to configure the request. The following example produces [share/outputs/trebor.mp3](share/outputs/trebor.mp3)
by using the voice replicated in the previous example:

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
		settings.WithToken("yourtoken"),
	)
	if err != nil {
		panic(err)
	}
	res, err := client.Speech().Create(
		speech.WithText("Hello! My name is Trebor"),
		speech.WithVoiceID("yourvoiceid"),
		speech.WithFormat("mp3"),
	)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	out, _ := os.Create("share/outputs/trebor.mp3")
	defer out.Close()
	io.Copy(out, res.Body)
}
```

## API Index

The Quick Start section above provides examples for the most interesting methods
in the library, but there is support for more than that:

* [speech.go](speech.go)
	* Speech.Create

* [voice.go](voice.go)
	* Voice.All
	* Voice.Get
	* Voice.Delete
	* Voice.Update
	* Voice.Clone

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
