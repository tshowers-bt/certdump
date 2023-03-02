docker run -e GOOS=windows -e GOARCH=amd64 -v "$($PSScriptRoot):/src" -w/src golang:alpine go build certdump
Write-Output "###`n# Local`n###`n"
.\certdump.exe
Write-Output "`n`n###`n# In Docker`n###`n"
docker run -v "$($PSScriptRoot):/src" -w/src golang:alpine go run certdump
