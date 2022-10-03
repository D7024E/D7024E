go test -coverprofile cover.out =./... ./...
echo "-------------------------------------------------------------"
go tool cover -func cover.out 
echo "-------------------------------------------------------------"
git clean -fxd --dry-run 
echo "-------------------------------------------------------------"
read -p "Do you want to clean [y/n]? " answer
go tool cover -html=cover.out
if [[ $answer = y ]] ; then
    git clean -fxd
fi