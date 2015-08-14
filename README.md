# Name Search for OpenStreetMap (-like) Data

For the [2015 edition of the Chaos Communication Congress](https://events.ccc.de/camp/2015/wiki/Main_Page) some great people set up some OpenStreetMap software (web frontend, API, rendering etc) which is used to [display a map of the field](http://campmap.mazdermind.de). But there was no text search, so often it was difficult to find something on the camp ground.

## Setting up camp geocoder

Take all the data from the API as .osm file, load and save it in JOSM to sort the nodes in the correct order and convert it to .pbf (see `make pbf`). Take this pbf file and import it into a PostgreSQL database (don't forget to enable PostGIS, hstore and pg_trgm), see `make import`. Lastly build and start camp-geocoder (`go run *.go`). It will accept HTTP requests on port 51009.

## Notes

It is not super fast or super intelligent, in reality it allows a lot of typo errors and it should still find something.
