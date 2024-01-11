package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/api/content", apiContentHandler)

	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Your Website Title",
	}

	tmpl, err := template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://unpkg.com/htmx.org@1.7.2/dist/htmx.min.js"></script>
    <script src="https://cdn.tailwindcss.com"></script>
    <link rel="stylesheet" href="/static/styles.css">
    <title>{{.Title}}</title>
</head>
<body>
<div class="flex flex-col min-h-screen">
  <header class="fixed top-0 w-full h-16 px-4 bg-white dark:bg-gray-900 flex items-center justify-between">
    <div class="font-bold text-lg">Your Name</div>
    <nav class="space-x-4">
      <a class="hover:underline" href="#">
        Home
      </a>
      <a class="hover:underline" href="#">
        About
      </a>
      <a class="hover:underline" href="#">
        Projects
      </a>
      <a class="hover:underline" href="#">
        Contact
      </a>
    </nav>
  </header>
  <main class="flex-1 mt-16">
    <section
      id="home"
      class="h-screen bg-gray-200 dark:bg-gray-800 flex items-center justify-center text-center p-4"
    >
      <div>
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Hello, I'm Your Name</h1>
        <p class="text-xl md:text-2xl mb-4">I'm a full-stack developer</p>
        <button class="inline-flex items-center justify-center text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 px-6 py-2 rounded bg-blue-500 text-white">
          View My Work
        </button>
      </div>
    </section>
    <section id="about" class="py-16 px-4 bg-white dark:bg-gray-900">
      <div class="max-w-2xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">About Me</h2>
        <div class="flex items-center justify-center space-x-4">
          <img
            src="/placeholder.svg"
            alt="Profile Picture"
            class="w-48 h-48 rounded-full"
            width="200"
            height="200"
            style="aspect-ratio: 200 / 200; object-fit: cover;"
          />
          <div class="text-left">
            <p class="mb-2">👋 Hi, I'm a full-stack developer with over 5 years of experience.</p>
            <p class="mb-2">🚀 I love creating high-performance applications using Go and JavaScript.</p>
            <p class="mb-2">💼 Currently working at Awesome Company.</p>
          </div>
        </div>
      </div>
    </section>
    <section id="projects" class="py-16 px-4 bg-gray-100 dark:bg-gray-800">
      <h2 class="text-3xl font-bold text-center mb-8">Projects</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div class="rounded-lg border bg-card text-card-foreground shadow-sm" data-v0-t="card">
          <div class="flex flex-col space-y-1.5 p-6">
            <h3 class="text-2xl font-semibold leading-none tracking-tight">Project 1</h3>
          </div>
          <div class="p-6">
            <img
              src="/placeholder.svg"
              alt="Project 1"
              class="w-full h-48 object-cover"
              width="200"
              height="200"
              style="aspect-ratio: 200 / 200; object-fit: cover;"
            />
            <p class="mt-4">A brief description of the project.</p>
          </div>
          <div class="flex items-center p-6">
            <button class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 w-full">
              View Details
            </button>
          </div>
        </div>
        <div class="rounded-lg border bg-card text-card-foreground shadow-sm" data-v0-t="card">
          <div class="flex flex-col space-y-1.5 p-6">
            <h3 class="text-2xl font-semibold leading-none tracking-tight">Project 2</h3>
          </div>
          <div class="p-6">
            <img
              src="/placeholder.svg"
              alt="Project 2"
              class="w-full h-48 object-cover"
              width="200"
              height="200"
              style="aspect-ratio: 200 / 200; object-fit: cover;"
            />
            <p class="mt-4">A brief description of the project.</p>
          </div>
          <div class="flex items-center p-6">
            <button class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 w-full">
              View Details
            </button>
          </div>
        </div>
        <div class="rounded-lg border bg-card text-card-foreground shadow-sm" data-v0-t="card">
          <div class="flex flex-col space-y-1.5 p-6">
            <h3 class="text-2xl font-semibold leading-none tracking-tight">Project 3</h3>
          </div>
          <div class="p-6">
            <img
              src="/placeholder.svg"
              alt="Project 3"
              class="w-full h-48 object-cover"
              width="200"
              height="200"
              style="aspect-ratio: 200 / 200; object-fit: cover;"
            />
            <p class="mt-4">A brief description of the project.</p>
          </div>
          <div class="flex items-center p-6">
            <button class="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 w-full">
              View Details
            </button>
          </div>
        </div>
      </div>
    </section>
    <section id="contact" class="py-16 px-4 bg-white dark:bg-gray-900">
      <div class="max-w-2xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-8">Contact Me</h2>
        <form class="space-y-4">
          <input
            class="flex h-10 rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 w-full"
            placeholder="Name"
          />
          <input
            class="flex h-10 rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 w-full"
            placeholder="Email"
            type="email"
          />
          <input
            class="flex h-10 rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 w-full"
            placeholder="Subject"
          />
          <textarea
            class="flex min-h-[80px] rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 w-full h-32"
            placeholder="Message"
          ></textarea>
          <button
            class="inline-flex items-center justify-center text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-primary/90 h-10 w-full px-6 py-2 rounded bg-blue-500 text-white"
            type="submit"
          >
            Send Message
          </button>
        </form>
      </div>
    </section>
  </main>
  <footer class="h-16 px-4 bg-white dark:bg-gray-900 flex items-center justify-center">
    <p class="text-sm text-gray-500 dark:text-gray-400">© Your Name</p>
  </footer>
</div>
</body>
</html>
`)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func apiContentHandler(w http.ResponseWriter, r *http.Request) {
	content := "This is dynamic content from the backend!"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"content": content})
}

