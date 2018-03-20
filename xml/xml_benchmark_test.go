package xml

import (
	"testing"
	"io/ioutil"
	"fmt"
	"encoding/xml"
)

func BenchmarkUnmarshalConciliation(b *testing.B) {
	for n :=0; n < b.N; n++ {
		content, _ := ioutil.ReadFile(fmt.Sprint("testdata/conciliation1.xml"))
		var conciliation Conciliation
		xml.Unmarshal(content, &conciliation)
	}
}
