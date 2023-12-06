package Gee

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	r := newRouter()
	ok := reflect.DeepEqual(r.parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(r.parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(r.parsePattern("/p/*name/*"), []string{"p", "*name"})

	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func newTestRouter() *router {
	r := newRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "assets/*filepath", nil)
	return r
}

func TestGetRouter(t *testing.T) {
	r := newTestRouter()

	node, params := r.getRouter("GET", "/hello/tom")
	if node == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if node.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if params["name"] != "tom" {
		t.Fatal("name should be equal to 'tom'")
	}

	fmt.Printf("matched path : %s, params['name'] = %s\n", node.pattern, params["name"])

}
