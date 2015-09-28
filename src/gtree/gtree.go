package gtree

import (
	"flag"
	"fmt"
	// "io"
	"bufio"
	"crypto/rand"
	// "io/ioutil"
	"os"
)

var (
	g_root  string
	g_size  int
	g_width int
	g_depth int
)

func init() {
	const (
		defaul_root  = ""
		usage_root   = "Root diractory path"
		defaul_width = 1
		usage_width  = "Diractory width"
		defaul_depth = 1
		usage_depth  = "Diractory tree depth"
		defaul_size  = 1024
		usage_size   = "Size of file inside every directory"
	)
	flag.StringVar(&g_root, "root", defaul_root, usage_root)
	flag.IntVar(&g_width, "width", defaul_width, usage_width)
	flag.IntVar(&g_depth, "depth", defaul_depth, usage_depth)
	flag.IntVar(&g_size, "size", defaul_size, usage_size)
}

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func writeRandomFile(path string, size int) {
	//http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file
	fo, err := os.Create(path)
	fmt.Println("Write File: ", path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(fo)
	buf := []byte(randString(size))
	if _, err := w.Write(buf); err != nil {
		panic(err)
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
}

func buildDirectoryTree(root string, w int, d int, size int) {
	if d > 0 {
		d--
		for i := 0; i < w; i++ {
			name := randString(5)
			name = root + "/" + name
			fmt.Println(name)
			os.Mkdir(name, 0755)
			writeRandomFile(name+"/"+randString(2)+".txt", size)
			// b := []byte(randString(size))
			// if err := ioutil.WriteFile(name+"/"+randString(2)+".txt", b, 0644); err != nil {
			// 	panic(err)
			// }
			go buildDirectoryTree(name, w, d, size)
		}
	}

}

func main() {
	flag.Parse()
	buildDirectoryTree(g_root, g_width, g_depth, g_size)
}
