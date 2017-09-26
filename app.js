var app = require('express')();
var http = require('http').Server(app);
var io = require('socket.io')(http);

app.get('/', function(req, res){
	res.sendFile( __dirname + '/index.html');
});

var clients = 0

io.on('connection', function(socket){
	console.log('A user connected');
	clients++;

	socket.emit('newClientConnect', {description: "Hey!"})
	socket.broadcast.emit('newClientConnect', {description:clients + ' clients connected'});
	// io.sockets.emit('broadcast', {description: clients + ' clients connected!'})

	socket.on('disconnect', function(){
		clients--;
		io.sockets.emit('broadcast', {description: clients + 'clients connected!'});
	});
});

http.listen(3000, function(){
	console.log('listening on *:3000');
});
