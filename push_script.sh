export GOPATH=$(pwd) && go get -u github.com/brotherlogic/writer && env GOOS=linux GOARCH=amd64 go build
ssh $1 'killall writer'
scp writer $1: && ssh $1 'nohup ./writer $2 $3 &> out.txt &'
