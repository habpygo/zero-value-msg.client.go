# mam.client.go

WARNING: Do not use this in a production environment! Under construction.
MAM specification is not fully implemented yet. Transactions are not secure as Merkle Tree authentication has not been implemented yet.

Small project to implement Masked Authenticated Messaging on the IOTA tangle with Golang.
This project is still under construction (see TODO) with the aim to get IoT sensors and devices to send MAMs.

Recently IOTA has introduced a pretty feisty test environment and test sites. You can find some sites in package `metadata`.

If you don't have a seed yet, follow the description here: https://iota.readme.io/docs/securely-generating-a-seed or, you
can use the `giotan` package at https://github.com/iotaledger/giotan. With this package you can easily generate a new seed by running `$ giotan new` then run `$ giotan addresses` and paste the generated address in the input field.


## Install

It is assumed that you have Golang installed. You also need to install the Go library API for IOTA which you can download at:

```javascript
go get -u github.com/iotaledger/giota
```

After that you can download the mamgoiota package.

```javascript
go get -u github.com/habpygo/mam.client.go
```

To be able to do testing and assertions you have to install the `stretchr` package

```javascript
go get -u github.com/stretchr/testify
```


## Sending MAMs to the IOTA tangle with Go

### API

#### Create a new Connection
```go
import gmam "github.com/habpygo/mam.client.go"

func main(){
    c, err := gmam.NewConnection("someNodeURL", "yourSeed")
    if c != nil && err == nil{
        fmt.Println("Connection is valid")
    }
}
```


#### Send a MAM to the IOTA tangle
```go
import gmam "github.com/habpygo/mam.client.go"

func main(){
    c, err := gmam.NewConnection("someNodeURL", "yourSeed")
    if err != nil{
        panic(err)
    }
    id, err := gmam.Send("the receiving address", 0, "your stringified message", c)
    if err != nil{
        panic(err)
    }
    fmt.Printf("Send to the Tangle. TransactionId: %v\n", id)
}
```
After sending, you find your transaction here https://testnet.thetangle.org/ giving the TransactionId


#### Read data from the IOTA tangle
Reading all transaction received by a certain adress:
```go
import gmam "github.com/habpygo/mam.client.go"

func main(){
    c, err := gmam.NewConnection("someNodeURL", "")
    if err != nil{
        panic(err)
    }

    ts, err := gmam.ReadTransactions("Receiving Address", c)
    if err != nil{
        panic(err)
    }
    for i, tr := range ts {
        t.Logf("%d. %v: %d IOTA, %v to %v\n", i+1, tr.Timestamp, tr.Value, tr.Message, tr.Recipient)
    }
}
```
The seed can be ommitted here, since reading does not require an account



Reading a special transaction by transactionID:
```go
import gmam "github.com/habpygo/mam.client.go"

func main(){
    c, err := gmam.NewConnection("someNodeURL", "")
    if err != nil{
        panic(err)
    }

    tx, err := gmam.ReadTransaction("Some transactionID", c)
    if err != nil{
        panic(err)
    }
    t.Logf("%v: %d IOTA, %v to %v\n", tx.Timestamp, tx.Value, tx.Message, tx.Recipient)
}
```


#### Examples
Check out our [example folder](/example) for a send and a receive example.

To run this, cd into the example folder and edit the `sender/send.go` and `receiver/receive.go` file, set the correct provider and address and you are ready to run.

Start the receiver first: `$ go run receiver/receive.go`. It will check for new messages every 5 seconds, until cancelled.

Then start the sender: `$ go run sender/send.go`.

You can also read all the past transactions, i.e. messages + value,  at the address: `go run history/history.go`.

If you pick up the transaction hash from the Terminal output and paste it into the input field on the site https://thetangle.org you find your transaction. 

If the Node is offline try another one, mentioned above.

TODO's are also pertinent to `webmamgiota`.
### TODOs
- [x] Implement MAM specs
- [ ] GoDoc
- [ ] Travis (This appeared to be a (solved) giota lib error)
- [ ] Make web-app (see webmamgiota). Under construction
- [ ] Read sensor data, e.g. RuuVi tag
- [ ] Make use of Merkle tree 
- [ ] More Read options
- [ ] Read by TransactionId





