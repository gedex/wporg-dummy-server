package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	// matches content in [xxx], where xxx is a digits.
	headrevContent = `
#
# ChangeLog for /
#
# Generated by Trac 1.0.1
# 02/16/15 18:17:51

Mon, 16 Feb 2015 18:16:33 GMT 1 [1091620]
`
	// matches content inside <a> as plugin name.
	allpluginsContent = `
<html><head><title> - Revision 1091621: /</title></head>
<body>
<h2> - Revision 1091621: /</h2>
<ul>
<li><a href="1/">1/</a></li>
<li><a href="2/">2/</a></li>
</ul>
<hr noshade><em>Powered by <a href="http://subversion.apache.org/">Apache Subversion</a> version 1.7.18 (r1615261).</em>
</body></html>
`

	// matches content in modified files (after '*' and before first '/').
	updatedpluginsContent = `
#
# ChangeLog for /
#
# Generated by Trac 1.0.1
# 02/16/15 18:25:50

Mon, 16 Feb 2015 18:16:33 GMT 1 [1091620]
	* 1/tags/1.0/readme.txt (modified)
	* 1/tags/1.0/1.php (modified)

	commit message


Mon, 16 Feb 2015 18:15:56 GMT 2 [1091619]
	* 2/tags/0.5.0/readme.txt (modified)

	commit message

`
)

var (
	addr  = flag.String("addr", "localhost:8081", "Listen address")
	file1 = flag.String("file1", "./1.zip", "File 1.zip")
	file2 = flag.String("file2", "./2.zip", "File 2.zip")
)

func main() {
	flag.Parse()

	// Checks if file1 and file2 exist.
	if _, err := os.Stat(*file1); os.IsNotExist(err) {
		log.Fatalf("File %s does not exists", *file1)
		os.Exit(2)
	}
	if _, err := os.Stat(*file2); os.IsNotExist(err) {
		log.Fatalf("File %s does not exists", *file2)
		os.Exit(2)
	}

	http.HandleFunc("/headrev", headrev)
	http.HandleFunc("/allplugins", allplugins)
	http.HandleFunc("/updatedplugins", updatedplugins)
	http.HandleFunc("/download/", download)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func headrev(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, headrevContent)
}

func allplugins(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, allpluginsContent)
}

func updatedplugins(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, updatedpluginsContent)
}

func download(w http.ResponseWriter, r *http.Request) {
	zip := r.URL.Path[len("/download/"):]
	switch zip {
	case "1.zip":
		http.ServeFile(w, r, *file1)
	case "2.zip":
		http.ServeFile(w, r, *file2)
	default:
		http.NotFound(w, r)
	}
}
