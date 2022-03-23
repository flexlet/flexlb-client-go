echo "building ..."

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/flexlb-cli ../cmd/flexlb-client/main.go


echo "copying ..."

DIR="/root/flexlb"
TARGET="root@<TARGET>"

scp /tmp/flexlb-cli ${TARGET}:${DIR}

#PROXY="root@<PROXY>"
#scp /tmp/flexlb-cli ${PROXY}:${DIR}
#ssh $PROXY "scp -r $DIR/flexlb-cli $TARGET:$DIR"