package main

type MyHandler struct{

}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello\n")
}

func main(){
  handler := MyHandler{}
  server := http.Server{
    Addr : "127.0.0.1:3013",
    Handler : &handler
  }

}
