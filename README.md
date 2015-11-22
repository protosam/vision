# vision
A template parser for web design to be easily updated without need to know more than basic HTML. 

#### Steps to setup
Make sure the GOPATH is setup
```
echo $GOPATH
```
If it is not, refer to this documentation: https://golang.org/doc/code.html  
  
  
Now we will need to put our package in it's place for reference.
```
mkdir -p $GOPATH/src/github.com/protosam
cd $_
git clone https://github.com/protosam/vision.git
```
  
  
Test it out:
```
cd vision/example
go run example.go
```
  
  
You can use the `example.go` code to begin your boiler place. Vision is extremely simple. The entire point is to completely separate your HTML from the code so that templates can be easily made by non-programmers. 
