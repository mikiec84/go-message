package mail_test

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/emersion/go-messages/mail"
)

func ExampleReader() {
	// Let's assume r is an io.Reader that contains a mail.
	var r io.Reader

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	// Read each mail's part
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch p := p.(type) {
		case *mail.Text:
			b, _ := ioutil.ReadAll(p.Body)
			log.Println("Got text: %v", string(b))
		case *mail.Attachment:
			filename, _ := p.Header.Filename()
			log.Println("Got attachment: %v", filename)
		}
	}
}
