// Om Namah Shivay
package main

import (
	"fiberMongo/connection"
	"fmt"
)

func main() {
	connection.DbConnection()
	fmt.Println("This is for fiber api with mongo db..")
}
