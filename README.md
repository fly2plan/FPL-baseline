# FPL-baseline

A demonstration of an AO and NM enacting the synchronisation and subsequent acceptance of an FPL by 'baselining' (see https://github.com/eea-oasis/baseline).

Note that the 'session' key used in JSON requests of the postman collection is the id of a Socket.IO socket (see https://socket.io/docs/v4/server-socket-instance/) and the 'workflowId' used in JSON requests should be replaced with the one returned in the 'workgroup:update' websocket event

### Prerequisites
1. Install [nvm](https://github.com/nvm-sh/nvm) to install Node.js v16.13.2

### Installation
1. Clone the repo
   ```
   git clone https://github.com/fly2plan/FPL-baseline.git
   ```
2. Install npm packages
   ```
   pushd server && \
   nvm use 16.13.2 && npm install && \
   popd
   ```
### Demo
1. Start docker containers
   ```
   pushd ops && ./run.sh
   ```
2. Run postman collection
3. Stop docker containers
   ```
   ./stop.sh && popd
   ```

### Disclaimer
All requests of the postman collection work except for the generation of the proof of acknowledgement of the flight plan.
