# FPL-baseline

A demonstration of an AO and NM enacting the synchronisation and subsequent acceptance of an FPL by 'baselining' (see https://github.com/eea-oasis/baseline).

Note that the 'session' key used in JSON requests of the postman collection is the id of a Socket.IO socket (see https://socket.io/docs/v4/server-socket-instance/)

### Prerequisites
1. Install [nvm](https://github.com/nvm-sh/nvm) to install Node.js v16.13.2
2. Install [ZoKrates](https://github.com/Zokrates/ZoKrates)

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
3. Generate circuit artifacts
   ```
   pushd server/baseline/privacy/circuits && \
   mkdir -p artifacts/{fpl,ack} && \
   cd ./artifacts/fpl && \
   zokrates compile -i ../../fpl_alt.zok && \
   zokrates setup && \
   cd ../ack && \
   zokrates compile -i ../../ack.zok && \
   zokrates setup && \
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
This repository is not yet fit for demonstration, but functional, due to the time needed to generate a witness and its accompanying ZKP.

Specifically using ZoKrates the time taken to compute a witness and generate a proof that testifies the validity of an edDSA signature can be seen below
```
time zokrates compute-witness -a 11235262813416749352946671613624768990483867072409116430966727976005139139840 511414244548838903499384480114964687808135619297892176794780594630704013108 9420157548029651208855531045168918868055444360621540115883508079663646576131 1033368034010549136882133934610770450228332062380241485674379400956488976079 9538509273393731932753091461103589732208060879506454810948329017717247018741 201669551810930794136676105378284253445 249706279142178819082221996773784642474 && time zokrates generate-proof
Computing witness...
Witness file written to 'witness'

real    0m2.106s
user    0m2.011s
sys     0m0.096s
Generating proof...
WARNING: You are using the G16 scheme which is subject to malleability. See zokrates.github.io/toolbox/proving_schemes.html#g16-malleability for implications.
Proof written to 'proof.json'

real    0m7.377s
user    0m7.261s
sys     0m0.116s
```
To resolve this the number of constraints of the witness should be reduced, which can be done by switching from SHA256 to a more skSNARK friendly hash function such as MiMC.
