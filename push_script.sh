env GOOS=linux GOARCH=amd64 go build writer && ssh $1 'killall writer' && scp writer $1: && ssh $1 'nohup ./writer &> out.txt &'
