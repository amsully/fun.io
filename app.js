var app = require('express')();
var http = require('http').Server(app);
var io = require('socket.io')(http);

app.get('/', function(req, res){
	res.sendFile( __dirname + '/index.html');
});

users=[]

io.on('connection', function(socket){
	console.log('A user connected');
	socket.on('setUsername', function(data) {
		if(users.indexOf(data) == -1 ) {
			users.push(data);
			socket.emit('userSet', {username: data});
		} else{
			socket.emit('userExists', data + ' usernmae is taken! Try something else')
		}
	})  // garbage nesting

	socket.on('msg', function(data) {
		// Send message to everyone
		io.sockets.emit('newmsg', data);
	});
});

http.listen(3000, function(){
	console.log('listening on *:3000');
});
