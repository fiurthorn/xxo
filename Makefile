build: build-linux build-android

build-linux:
	go build -tags=nowayland -o build/xxo .

build-android:
	gogio -target=android -o build/xxo.apk .
