#  redstone.go root directory structure
## jars/
	directory for minecraft server jars
## worlds/
	each world is stored in its own subdirectory inside worlds/
## redstone.json
	redstone.go's configuration file
```
.
|-- jars
|   |-- minecraft_server.<X>.<Y>.jar
|   `-- <...>
|-- worlds
|   |-- <some world's subdirectory>
|   `-- <...>
|-- <cert>.pem
|-- <key>.pem
`-- redstone.json
```
