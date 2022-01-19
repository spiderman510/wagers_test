set -e
./scripts/init_test_env.sh

generate_report=${1:-"yes"}
echo "Start running tests..."
test_pkgs=(wagers/internal/service)
test_pkgs+=(wagers/internal/repository)

coverpkgs=""
for value in ${test_pkgs[@]}; do
    coverpkgs+="$value,"
done

go test -v -failfast -coverpkg=$coverpkgs -coverprofile=coverage.out ${test_pkgs[@]}
if [[ "yes" == "$generate_report" ]];
then
    go tool cover -html=coverage.out
fi