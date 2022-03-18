build: build-linux build-android

build-android: build-android-apk build-android-aar build-android-aab

build-linux:
	go build -tags=nowayland -o build/xxo .

build-android-apk:
	gogio -target=android -o build/xxo.apk -signkey=../keys/android/xxo.keystore -signpass=1234567890 .

build-android-aar:
	gogio -target=android -buildmode=archive -o build/xxo.aar -signkey=../keys/android/xxo.keystore -signpass=1234567890 .

build-android-aab:
	gogio -target=android -o build/xxo.aab -signkey=../keys/android/xxo.keystore -signpass=1234567890 .
