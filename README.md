# Nexmo Go

A Golang client library for nexmo.

Currently this library only implements the verify, check, and verify-search endpoints.
The interface is still in heavy flux, but if you're interested, it looks a bit like this:

```go
import "github.com/judy2k/nexmo-go/nexmo"

// Create a client for making new requests:
client := nexmo.NewClient("api-id", "api-secret")

vResponse, err := client.Verify(nexmo.NewVerifyRequest("447712345678", "My App"))
if err != nil {
    log.Fatal(err)
}
// Get the requestID from the response:
requestID := vResponse.RequestID

// ...

// Check the code entered is correct
cResponse, err := client.Check(requestID, code)
if err != nil {
    log.Fatal(err)
}
log.Println(cResponse)
```


