package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"brian", "anashari", ""})
	_ = writer.Write([]string{"indah", "cahya", "nabilla"})
	_ = writer.Write([]string{"cornelia", "vanisa", ""})
	_ = writer.Write([]string{"gita", "sekar", "andarini"})

	writer.Flush()
}
