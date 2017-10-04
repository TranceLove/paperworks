bash "Install godep" do
    code "#{node["golang"]["binpath"]}/go get github.com/tools/godep"
    cwd node["golang"]["default_gopath"]
    user node["golang"]["goget"]["runas"]
    environment ({
        "GOPATH" => node["golang"]["default_gopath"]
    })
end
