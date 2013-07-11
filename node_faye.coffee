Faye = require("faye")
Faye.Logging.logLevel = 0
server = new Faye.NodeAdapter(mount: "/faye")
server.listen 4000

# client = new Faye.Client("http://192.168.9.41:3000/faye")

# client.subscribe "/foo", (data) ->
#   console.log ("Server Got a message: #{data} - #{data.message}")

# client.publish "/foo",
#   message: "Hello world"
