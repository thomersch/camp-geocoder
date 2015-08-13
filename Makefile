import:
	imposm3 import 

build-osmconvert:
	wget -O - http://m.m.i24.cc/osmconvert.c | cc -x c - -lz -O3 -o osmconvert

pbf: osmconvert
	wget "http://maps.c3voc.de/api/0.6/map?bbox=13.2776,53.0222,13.345,53.0442" -O camp.osm
	./osmconvert camp.osm -o=camp.pbf

import:
	