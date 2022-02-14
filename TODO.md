### For cli.go

- Add Function recursion in `takeInput()`
- Add Type Conversions 

### For server.go

- Remove unnecessary `fmt.*` statements and Make it
customizable
like specifying port-no
- Turn it into a http and websocket server, use socket.io
in future
- Turn it into a service, so it can be accessed any time
from cli and client
- implement working with data types
- Enable authentication System where passwords are compared
using salt based algorithms like bcrypt
