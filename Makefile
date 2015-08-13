PG_DATABASE		= campcoder
OSM_FILE		= camp.osm
PBF_FILE		= camp.pbf
MAPPING_FILE	= imposm_mapping.json
IMPOSM_PATH		= "$(GOPATH)/bin/imposm3"

build-osmconvert:
	curl  http://m.m.i24.cc/osmconvert.c | cc -x c - -lz -O3 -o osmconvert

pbf: osmconvert
	# wget "http://maps.c3voc.de/api/0.6/map?bbox=13.2776,53.0222,13.345,53.0442" -O $(OSM_FILE)
	./osmconvert $(OSM_FILE) -o=$(PBF_FILE)

import:
	$(IMPOSM_PATH) import -read $(PBF_FILE) -mapping=$(MAPPING_FILE) -write -connection $(PG_CONNSTR) -overwritecache -deployproduction -srid 4326
