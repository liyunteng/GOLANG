package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"html/template"
	"path"
	"strings"
	"runtime/debug"
)

const (
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./views"
	PUBLIC_DIR = "./public"
)
const (
	ListDir = 0x0001
)

var templates = make(map[string] *template.Template)
func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
	}

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template: ", templateName)

		t := template.Must(template.ParseFiles(templatePath))
		templateName = strings.TrimSuffix(templateName, ".html")
		templates[templateName] = t
	}
}

func uploadHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderHtml(w, "upload", nil);
		check(err)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()

		_, err = io.Copy(t, f)
		check(err)

		http.Redirect(w, r, "/view?id=" + filename,
			http.StatusFound)

	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExist(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images,fileInfo.Name())
	}
	locals["images"] = images

	err = renderHtml(w, "list", locals)
	check(err)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error {
	return templates[tmpl].Execute(w, locals)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(),
					http.StatusInternalServerError)
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				log.Printf("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}


func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {

	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		// log.Println(file)
		if (flags & ListDir) == 0 {
			if exists := isExist(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}
func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", PUBLIC_DIR, 0)
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/", safeHandler(listHandler))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
