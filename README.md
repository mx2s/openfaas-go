# openfaas-go-httprouter

Go function for OpenFaas that uses [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter). Can be deployed to OpenFaas or can run locally as a webserver

## Usage instructions

Default webserver port (local): 8201

1. Install go modules dependencies
2. Run `scripts/build_binary.sh` - it will build go binary and will put it into `function/dist/handler`
3. You can run the binary itself and it will start local server. ; Or you can run `faas-cli build -f openfaas.yml` and it will build docker image that can be deployed to OpenFaas

## Additional notes

This approach is not tied to `golang` only, since it uses dockerfile and builds and image based on `alpine linux` you can cuztomize dockerfile and put any executable into it, just make sure that it will listen to a correct port

## Contribution

If there's any improvements etc. that you want to make - feel free to open up an Issue or a PR :)
Thanks
