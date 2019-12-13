env GOOS=linux GOARCH=arm GOARM=7 go build
git commit -m "Update software" exec
git push origin master