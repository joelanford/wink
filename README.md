# Wink: Golang API for Wink

Package `wink` is a [Go](https://golang.org/) client library for interacting with
the [Wink](http://www.wink.com) [API](http://docs.wink.apiary.io/). 


## Status

[![Travis Build Status](https://travis-ci.org/joelanford/wink.png)](https://travis-ci.org/joelanford/wink)

## Installation 

### Requirements

Wink requires Go 1.5 or later.


### Development

```
go get github.com/joelanford/wink
```

#### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/joelanford/wink"
)

func main() {
	c := wink.NewClient("<my_wink_client_id>", "<my_wink_client_secret>")
	err := c.Authenticate("<my_wink_email_address>", "<my_wink_password>")
	if err != nil {
		log.Fatalln(err)
	}

	devicesJSON, err := c.GetDevicesJSON()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(devicesJSON)
	}

	devices, err := c.GetDevices()
	if err != nil {
		log.Fatalln(err)
	} else {
		for _, device := range devices {
			fmt.Printf("%+v\n", device)
		}
	}
}
```


## Documentation

See [![GoDoc](http://godoc.org/github.com/joelanford/wink?status.png)](http://godoc.org/github.com/joelanford/wink)
for automatically generated API documentation.

This package's documentation is a work-in-progress. Pull requests are appreciated!

## Support

This is an unofficial client for the Wink API. Please do not contact Wink unless
you're sure the problem is not in this library. Please submit issues and/or pull
requests if you find issues with the package.


## Contributing

Contributions in the form of Pull Requests are gladly accepted.

If you have a Wink device that's not yet supported, please create a new issue
containing the full JSON device output or a pull request that adds support
for your device. 


## License

This is Free Software, released under the terms of the [GPL v3](http://www.gnu.org/licenses/gpl-3.0.en.html).
