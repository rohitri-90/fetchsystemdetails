This is a simple utility written in golang to fetch the details about macbook using [macaddress.io](https://macaddress.io/)

### Building docker image

* clone the `repository`using git clone <URL>
* Navingate to the cloned directory and execute below commands :
* `docker build -t  fetchmacdetails:latest .`

### Running docker image to fetch the details

* To run the docker image and fetch the details macapi key and mac address of the system is required.
* macapi key can be generated by following the instructions [here](https://macaddress.io/api/documentation/making-requests)

* command to run the docker image :
   - `docker run -e MAC_ADDRESS=<mac address> MAC_API_TOKEN=<api token> fetchmacdetails:latest`