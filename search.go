package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
)

type Result struct {
	OSMID    int64
	Name     string
	Geometry GeoJSON
}

type GeoJSON interface{}

var db *sql.DB

func bootDB(connStr string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func Search(text string) ([]*Result, error) {
	rows, err := db.Query(`SELECT sub.osm_id, sub.name, sub.geom FROM
		(SELECT DISTINCT ON (osm_id) osm_id, name, ST_AsGeoJSON(geometry) AS geom, set_limit(0.1) FROM osm_all WHERE name % $1 LIMIT 10) AS sub
		ORDER BY similarity(sub.name, $1) DESC;`, text)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var resList []*Result
	for rows.Next() {
		var res Result
		var geojson []byte
		err := rows.Scan(&res.OSMID, &res.Name, &geojson)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		// unpack geojson
		err = json.Unmarshal(geojson, &res.Geometry)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		resList = append(resList, &res)
	}
	return resList, nil
}
