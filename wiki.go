package main

//An example of a basic wiki running on a server in go, Collaborative set of pages able to be edited and saved by different users

import(
	//"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"html/template" //will allow us to put the ugly looking HTML code in seperate files
	"regexp" //will be used for page validation 
)

//Page structures and methods
type Page struct {
	Title string
	Body []byte //byte instead of string because thats what i/o packages expect for a type
}

func (p *Page) save() error {
	filename := p.Title + ".txt" //makes the filename the title of the page plus txt 
	return ioutil.WriteFile(filename, p.Body, 0600) //will return Nil if all goes well error if not, 0600 indicates to make the file read only for the user
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) //read the file contents into the body member 
	if err != nil {
		return nil,err
	}
	return &Page{Title: title, Body: body}, nil // return page address with the filename as title and body as body
}


//Handlers to use http package to deal with web requests
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return //if you visit a page that does not exist then redirect to the edit page so they may enter content
	}
	formTemplate(w,"view",p) //use function we made to render the template
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p,err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	formTemplate(w,"edit",p) //reuse function, same as last handler
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body") //this returns a string we need it as a []byte
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save() //use our existing save page function
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //dealing with possible error when saving the page
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound) //save the page and redirect back to just a viewing page
}


var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
//instead of running parseFiles everytime we create or load a page, just call it once on the first run throughj to parse all templates into one(this lets us use ExecuteTemplate function for a specific one) 

func formTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//reused code in the handlers so put into its own function
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError) //in case there is an error rendering template
	}
}


var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
//will compile regex and will panic if expression fails (can use this to indicate error)

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path) //check the title of the page to make sure it is a valid path from our regex
		if m == nil {
			http.NotFound(w,r) //if the title is invalid then throw an error
			return
		}
		fn(w, r, m[2]) //if the title is valid then fn will create either a view,edit, or save handler 
	}
	//by taking in the titles here we have saved ourselves lots of needless code in the actual handlers
}


func main(){
	//original code to create, save, and load a page
	/*
	p1 := &Page{Title: "Test", Body: []byte("This is a test page")}
	p1.save()
	p2,_ := loadPage("Test")
	fmt.Println(string(p2.Body))
	*/

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080",nil)) //Opens the server on local port 8080, ListenAndServe maintains it and will log a fatal error if anything goes wrong
}
