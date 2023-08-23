https://www.youtube.com/watch?v=ASBUp7stqjo&ab_channel=AkhilSharma
Web Server using Golang

`http` package :

- `http`.FileServer(root FileSystem) returns a function that handle HTTP requests with content system of the file system
- `FileSystem` http.Dir(`path_to_root_dir_of_content`)
- `http.Handle(pattern, fileServer)` register that function that handle requests for a given pattern. `pattern` means PATH
- `http.HandleFunc(pattern, func)` register that handler function for a given pattern
- `http.ListenAndServe(":8080", nil)` start the server

`fmt` package :

- `fmt.Printf("%s", string_value)` print out something into console
