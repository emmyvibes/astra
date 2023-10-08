module github.com/emmyvibes/paul/libpaul

go 1.20

require github.com/fjl/go-couchdb v0.1.0

//replace dario.cat/mergo => github.com/imdario/mergo latest
replace github.com/imdario/mergo => dario.cat/mergo v1.0.0
