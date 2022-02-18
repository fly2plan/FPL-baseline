let io;

const getIO = () => {
  return io;
}

const getSocket = (id) => {
  return io.sockets.sockets.get(id);
}

const socketConnection = (server) => {
  io = require('socket.io')(server, {cors: {origin: "*"}});

  io.on('connection', (socket) => {
    socket.emit('session:id', socket.id);
    console.log('client connected', socket.id);
  });
}

module.exports = {
  socketConnection,
  getIO,
  getSocket
}