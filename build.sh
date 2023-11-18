echo "Running go mod tidy..." && go mod tidy && echo "All dependencies installed";
cd cmd;
echo "Building Project..." && go build -o ../bin/goshowtree && echo "Build Complete";
echo "";
echo "You can find gobuildtree in the bin folder";
echo "Type ./gobuildtree <path> to view your tree";