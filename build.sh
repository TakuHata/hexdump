APP_NAME="hexdump"

platforms=(
	"windows/amd64"
	"windows/386"
	"linux/amd64"
	"linux/arm64"
	"darwin/amd64"
	"darwin/arm64"
)

mkdir -p build

for platform in "${platforms[@]}"
do
	IFS="/" read -r -a platform_split <<< "$platform"
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}

	OUTPUT_NAME=$APP_NAME'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		OUTPUT_NAME+='.exe'
	fi

	echo "Сборка для $GOOS/$GOARCH..."

	env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s -w" -o build/$OUTPUT_NAME .

	if [ $? -ne 0 ]; then
			echo "Ошибка сборки для $GOOS/$GOARCH"
			exit 1
	fi
done

echo "Все билды успешно созданы в папке ./build/"
