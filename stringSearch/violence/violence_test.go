package violence

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println(SearchOne("ABC", "BCBABC"))
	fmt.Println(SearchAll("AB", "BCBABCABABCCC"))

}
