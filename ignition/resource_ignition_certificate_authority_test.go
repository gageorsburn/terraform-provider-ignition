package ignition

var foo = `{
	"ignition": {
	  "version": "2.2.0-experimental",
	  "config": {
		"append": [{
		  "source": "https://s3.com/securely-fetched-config.ign"
		}]
	  },
	  "security": {
		"tls": {
		  "certificateAuthorities": [
			{
			  "source": "http://www.example.com/root.pem",
			  "verification": {
				"hash": "sha512-ab800f66a7544c0a8dbed0c57b38a3c1487c3369e2e9e90704d0c07743557ab2a28c528720566ffc64e3dfd5df1a557a4979b33009f5fd493fea02a7e30041d2"
			  }
			}
		  ]
		}
	  }
	}
  }`
