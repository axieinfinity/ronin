set -ex

case "$1" in
    "testnet")
    image_name="testnet"
    ;;

    "mainnet")
    image_name="mainnet"
    ;;

    *)
    echo "Please follow syntax: ./build.sh testnet|mainnet <build_tag>"
    exit 1
    ;;
esac

build_tag="$2"
if [ ${#build_tag} -lt 1 ]; then
  echo "Please specify a build tag!"
  exit
fi

commit_hash=$(git rev-parse HEAD | cut -c 1-6)
image_path=docker.axieinfinity.co/${image_name}:${build_tag}-${commit_hash}

docker build . -f docker/chainnode/Dockerfile -t ${image_path}
docker push ${image_path}

echo ${image_path}
