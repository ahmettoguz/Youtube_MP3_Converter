const WebSocket = require("ws");
const dotenv = require("dotenv");
const commonService = require("./commonService");

dotenv.config();

class ServerWebsocketService {
  constructor(port) {
    this.wsServer = new WebSocket.Server({ port });
    this.clients = new Map();
    this.currentClientId = commonService.generateRandomWord();

    this.wsServer.on("connection", (ws) => {
      this.clients.set(this.currentClientId, ws);

      console.info("Client connected to websocket");

      ws.send(
        JSON.stringify({
          status: "connected",
          message: "Connected to websocket",
          clientId: this.currentClientId,
        })
      );

      ws.on("message", (message) => {
        console.info(
          `Received message from client ${this.currentClientId}: ${message}`
        );
      });

      ws.on("close", () => {
        console.info(
          `Client ${this.currentClientId} disconnected from websocket`
        );
        this.clients.delete(this.currentClientId);
      });
    });
  }

  // getter for client id 
  get getCurrentClientId() {
    return this.currentClientId;
  }

  // Send message to every clients
  broadcast(message) {
    this.wsServer.clients.forEach((client) => {
      if (client.readyState === WebSocket.OPEN)
        client.send(JSON.stringify(message));
    });
  }

  // Send message to a specific client
  sendMessageToClient(clientId, message) {
    const client = this.clients.get(clientId);
    if (client && client.readyState === WebSocket.OPEN) {
      client.send(JSON.stringify(message));
    } else {
      console.error(`Client ${clientId} not found or not ready`);
    }
  }
}

const serverWebsocketService = new ServerWebsocketService(
  process.env.WEBSOCKET_PORT
);
module.exports = serverWebsocketService;
