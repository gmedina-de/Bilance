GOARCH=arm go build main.go
ssh -t pi@raspberrypi 'sudo killall main'
scp -r main database.db view pi@raspberrypi:/home/pi
ssh -t pi@raspberrypi '/home/pi/main'