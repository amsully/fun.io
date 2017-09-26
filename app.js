var app = require('express')();
var http = require('http').Server(app);
var io = require('socket.io')(http);

app.get('/', function(req, res){
	res.sendFile( __dirname + '/index.html');
});

io.on('connection', function(socket){
	console.log('A user connected');

	// Send message after timeout of 4 seconds
	setTimeout(function(){
		// socket.send('Sent a message 4 seconds after connection');
		socket.emit('testerEvent', {description: 'A custom event named testerEvent'} );
	}, 4000);

	socket.on('clientEvent', function(data) {
		console.log(data);
	});

	// Whenever somebody disconnects
	socket.on('disconnect', function() {
		console.log('A user disconnected');
	});
});

http.listen(3000, function(){
	console.log('listening on *:3000');
});
