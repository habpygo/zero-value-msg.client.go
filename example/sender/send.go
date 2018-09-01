/*
MIT License
Copyright (c) 2017 Harry Boer, Jonah Polack

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"time"

	gmam "github.com/habpygo/zero-value-msg.client.go"
	"github.com/habpygo/zero-value-msg.client.go/metadata"
)

func main() {
	//"https://testnet140.tangle.works"
	//WARNING: The nodes have a nasty habit to go on/off line without warning or notice. If this happens try to find another one.
	//c, err := gmam.NewConnection("http://eugene.iota.community:14265", seed)
	c, err := gmam.NewConnection(metadata.Provider, metadata.Seed)
	if err != nil {
		panic(err)
	}

	msgTime := time.Now().UTC().String()
	message := "This message was generated on mam.client.go on: " + msgTime

	id, err := gmam.Send(metadata.Address, 0, message, c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sent Transaction: %v\n", id)
}
