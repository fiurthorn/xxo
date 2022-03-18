export PASS=1234567890

build: build-linux build-android

build-android: build-android-apk build-android-aar build-android-aab

build-linux: export GOOS=linux
build-linux:
	@echo build linux
	@go build -tags=nowayland -o build/xxo .

build-windows:
	@echo build windows.exe
	@gogio -target=windows -o build/xxo.exe -minsdk=7 .

build-android-apk:
	@echo build android.apk
	@gogio -target=android -o build/xxo.apk -signkey=../keys/android/xxo.keystore -signpass=${PASS} .

build-android-aar:
	@echo build aar
	@gogio -target=android -buildmode=archive -o build/xxo.aar -signkey=../keys/android/xxo.keystore -signpass=${PASS} .

build-android-aab:
	@echo build aab
	@gogio -target=android -o build/xxo.aab -signkey=../keys/android/xxo.keystore -signpass=${PASS} .
