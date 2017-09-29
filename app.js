var app = require('express')();
var http = require('http').Server(app);
var io = require('socket.io')(http);

app.get('/', function(req, res){
	res.sendFile( __dirname + '/index.html');
});

var roomno = 1;

io.on('connection', function(socket){
	// increase.. use def namespace
	if(io.nsps['/'].adapter.rooms["room-" + roomno] && io.nsps['/'].adapter.rooms["room-"+roomno].length > 1)
		roomno++;

	socket.join("room-" + roomno);

	// send events to people in the room
	io.sockets.in("room-" + roomno).emit('connectToRoom', "You are in room num: " + roomno);
});

http.listen(3000, function(){
	console.log('listening on *:3000');
});
