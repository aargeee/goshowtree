echo "Running go mod tidy..." && go mod tidy && echo "All dependencies installed";
echo "Running tests...";
go test -v;