package config

/*
Server network settings:
*/
const PORT = ":8001"
const ADDRESS = "0.0.0.0"

/*
Server exterior request config
*/
const UseCors = true

/*
certificate for https:
*/
const UseTls = false
const TlsFileCert = "server.crt"
const TlsFileKey = "server.key"

/*
database connection
*/

/*
encryption settings
*/
const SHARED_KEY = "Hunter2"  /// for example: using shared key for JWT symetric signing
