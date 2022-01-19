OUTPUT=test/mocks
SRC_PKG=("wagers/internal/repository")
SRC_PKG+=("wagers/internal/service")

for src_pkg in ${SRC_PKG[@]}
do
    echo $src_pkg
    docker run -v "$PWD":/src -w /src vektra/mockery:v2.1 --all -r --case underscore --srcpkg $src_pkg --output $OUTPUT
done

sudo chown $USER:$USER test/mocks
sudo chown $USER:$USER test/mocks/*