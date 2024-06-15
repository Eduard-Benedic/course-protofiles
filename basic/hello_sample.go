package basic

import (
	"log"

	"github.com/Eduard-Benedic/course-protofiles/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Clark Kent",
	}

	log.Println(&h)
}
