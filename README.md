# infobip

Infobip API client library in Go

[![Build Status](https://travis-ci.org/nuveo/infobip.svg?branch=master)](https://travis-ci.org/nuveo/infobip)
[![GoDoc](https://godoc.org/github.com/nuveo/infobip?status.png)](https://godoc.org/github.com/nuveo/infobip)
[![Go Report Card](https://goreportcard.com/badge/github.com/nuveo/infobip)](https://goreportcard.com/report/github.com/nuveo/infobip)

## Usage

To initiate a client, you should use the `infobip.ClientWithBasicAuth` func. This will returns a pointer to `infobip.Client` and allows to you use features from service.

### Sending a single message

The func needs a `infobip.Message` struct. That struct consists of the following attributes:

| Attribute | Type | Description |
|-----------|------|-------------|
| From | string | Represents a sender ID which can be alphanumeric or numeric |
| To | string | Message destination address |
| Text | string | Text of the message that will be sent |

It has a func to validate the `From` and `To` attributes, according to Infobip docs, and it is used into all funcs that make a request to the service. The following code is a basic example of the validate func:

```go
package main

import (
    "log"

    "github.com/nuveo/infobip"
)

func main() {
    m := infobip.Message{
        From: "Company", // or company number
        To:   "41793026727",
        Text: "This is an example of the body of the SMS message",
    }
    err := m.Validate()
    if err != nil {
        log.Fatalf("Infobip message error: %v", err)
    }
}
```

Finally, the following code is a full example to send a single message to Infobip service:

```go
package main

import (
    "fmt"
    "log"

    "github.com/nuveo/infobip"
)

func main() {
    client := infobip.ClientWithBasicAuth("foo", "bar")
    r, err := client.SingleMessage(m) // "m" refers to the variable from the previous example
    if err != nil {
        log.Fatalf("Infobip error: %v", err)
    }
    fmt.Printf("Infobip response: %v", r)
}
```

### Sending a advanced message

The func needs a `infobip.BulkMessage` struct. That struct consists of the following attributes:

| Attribute | Type | Description |
|-----------|------|-------------|
| ID | string | The ID which uniquely identifies the request |
| Messages | slice of Message | Message values |

And the `infobip.Message` struct consists in the following attributes:

| Attribute | Type | Description |
|-----------|------|-------------|
| From | string | Represents a sender ID which can be alphanumeric or numeric |
| Destinations | slice of Destination | Destination values |
| Text | string | Text of the message that will be sent |
| Transliteration | string | Conversion of a message text from one script to another |
| LanguageCode | string | Code for language character set of a message text |

And finally the `infobip.Destination` struct consists in the following attributes:

| Attribute | Type | Description |
|-----------|------|-------------|
| ID | string | The ID that uniquely identifies the message sent |
| To | string | Message destination address |

The following code is a basic example of the validate func:

```go
package main

import (
    "log"

    "github.com/nuveo/infobip"
)

func main() {
    m := infobip.BulkMessage{
        Messages: []infobip.Message{
            infobip.Message{
                From: "Company", // or company number
                Destinations: []infobip.Destination{
                    infobip.Destination{
                        To: "41793026727",
                    },
                },
                Text:            "This is an example of the body of the SMS message",
                Transliteration: "PORTUGUESE",
                LanguageCode:    "PT",
            },
        },
    }
    err := m.Validate()
    if err != nil {
        log.Fatalf("Infobip message error: %v", err)
    }
}
```

Finally, the following code is a full example to send an advanced message to Infobip service:

```go
package main

import (
    "fmt"
    "log"

    "github.com/nuveo/infobip"
)

func main() {
    client := infobip.ClientWithBasicAuth("foo", "bar")
    r, err := client.AdvancedMessage(m) // "m" refers to the variable from the previous example
    if err != nil {
        log.Fatalf("Infobip error: %v", err)
    }
    fmt.Printf("Infobip response: %v", r)
}
```
