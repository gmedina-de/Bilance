build:
	GOARCH=arm go build main.go
	scp main pi@raspberrypi:/home/pi
	rm main