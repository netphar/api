# API specifications
## all GET requests
* all doses: `api.drugcomb.fimm.fi/doses`
* all combinations `api.drugcomb.fimm.fi/combinations`
* all conditions `api.drugcomb.fimm.fi/conditions`
* all drugs `api.drugcomb.fimm.fi/drugs`
* all cell_types as per cellosaurus `api.drugcomb.fimm.fi/cells`
* single dose by ID `api.drugcomb.fimm.fi/dose/{id:[0-9]}`
* single dose block by block_ID `api.drugcomb.fimm.fi/doses/{id:[0-9]}`
* combination by ID `api.drugcomb.fimm.fi/combination/{id:[0-9]}`
* healthcheck to test whether API is live `api.drugcomb.fimm.fi/healthcheck`

## POST and PUT requests used to create/update/delete new entries are not disclosed publicly, please contact repo owner for more details
