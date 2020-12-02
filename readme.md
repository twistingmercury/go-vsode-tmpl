# VSCode Go RESTful Service Template

This is a go RESTful API template demonstrating how to create a code template for other than .NET.  It doesn't necessarily represent the best way to setup a Go project, though it tries to present a solid workable base solution.

## Prequisites

* [Go](https://golang.org/doc/install)
* dotnet cli.  That means you'll need to have the [.NET Core SDK](https://dotnet.microsoft.com/download) installed.  There is a installer for Windows, OSX, and Linux.
* git
* (optional) [Visual Studio Code](https://code.visualstudio.com/)


## Template Installation

1. Navigate to the `template` directory: `cd ~/go/src/github.com/twistingmercury/go-vscode-tmpl/template`
2. Execute dotnet install: `dotnet new --install .`
3. You should this entry below within list of templates currently installed:

```
Templates                                         Short Name               Language          Tags                                 
----------------------------------------------------------------------------------------------------------------------------------
[omitted for brevity] 
...                      
Go Rest                                           gorest                   [Go]              Web/RESTful/Go                       
...
```

## Use the template

1. Navigate to the directory to where you want to create the project.  For example, I create mine in `~/go/src/github.com/twistingmercury` since that's the path of my projects, you know, so that I can use `go get`.
2. Run the `dotnet new` command: `dotnet new gorest -o sample -n sample`. You should see the message `The template "Go Rest" was created successfully.`
3. Navigate to the directory where your new project was created: `cd sample`.
4. You may need to run `go mod tidy`. Just run it to be safe.
5. Next, initialize git: `git init`
6. Optionally, initialize your git workflow.  I'm partial to git flow: `git flow init`.
7. After that, run `make` and you should see a successful build.