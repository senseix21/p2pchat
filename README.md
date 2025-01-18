
```markdown
# P2P Chat Application

A simple peer-to-peer (P2P) chat application built with [libp2p](https://libp2p.io/). This application allows two peers to establish a direct connection and exchange messages using the `/chat/1.0.0` protocol.

## Features
- Peer-to-peer communication.
- Bidirectional chat with message streaming.
- Simple and customizable protocol (`/chat/1.0.0`).

## Requirements
- Go (1.19 or later)
- Basic understanding of terminal commands.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/p2p-chat.git
   cd p2p-chat
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o p2p-chat
   ```

## Usage

### Start a Listener
1. Run the listener on a specific port:
   ```bash
   ./p2p-chat -port=9000
   ```
   Example output:
   ```
   Host created with: QmUzp8XNr7aShjL8BqWf8DbKvkukTX3rK4fueHjd7u8CYq
   Listening on: /ip4/127.0.0.1/tcp/9000/p2p/QmUzp8XNr7aShjL8BqWf8DbKvkukTX3rK4fueHjd7u8CYq
   Listening on: /ip4/192.168.87.152/tcp/9000/p2p/QmUzp8XNr7aShjL8BqWf8DbKvkukTX3rK4fueHjd7u8CYq
   ```
   Share the multiaddress (e.g., `/ip4/127.0.0.1/tcp/9000/p2p/QmUzp8XNr7aShjL8BqWf8DbKvkukTX3rK4fueHjd7u8CYq`) with the peer who wants to connect.

### Start a Chat as Sender
1. Use the shared peer address to start a chat:
   ```bash
   ./p2p-chat -port=9001 -target=/ip4/127.0.0.1/tcp/9000/p2p/QmUzp8XNr7aShjL8BqWf8DbKvkukTX3rK4fueHjd7u8CYq
   ```
   Example output:
   ```
   Host created with: QmZhBmMnkU82xSb1YwYddnKvc1SxeVH4fYHn8eapQ1yVra
   Listening on: /ip4/127.0.0.1/tcp/9001/p2p/QmZhBmMnkU82xSb1YwYddnKvc1SxeVH4fYHn8eapQ1yVra
   Connecting to peer: QmUzp8XNr7aShjL8BqWf8DbKvkukTX3rK4fueHjd7u8CYq
   Successfully connected to the peer!
   Chat session started. Type your messages below:
   ```

### Example Chat Session
- **Listener Output**:
  ```
  2025/01/18 09:50:42 Incoming chat request
  2025/01/18 09:51:11 Peer: can you message me?
    You: Heelo how are you?
  2025/01/18 10:43:58 Peer: I am fine, what about you?
  ```

- **Sender Output**:
  ```
  2025/01/18 09:50:42 Connecting to peer: QmUzp8XNr7aShjL8BqWf8DbKvkukTX3rK4fueHjd7u8CYq
  Successfully connected to the peer!
  2025/01/18 09:50:42 Chat session started. Type your messages below:
  You: hello from the other side
    You: I am fine, what about you?
  ```

### Command-Line Options
- `-port=<port>`: Port to bind the application for incoming connections.
- `-target=<peer-address>`: Target peer address to connect.
- `-insecure`: Use unencrypted connections (for testing purposes).
- `-seed=<number>`: Set a custom seed for random ID generation.

## Code Overview
- **`StartChat`**: Establishes a connection to a peer and enables bidirectional communication.
- **`HandleIncomingStream`**: Listens for incoming connections and handles chat sessions.

## Contributing
Contributions are welcome! Feel free to submit issues or pull requests for enhancements and bug fixes.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments
- Built using [libp2p](https://github.com/libp2p/go-libp2p).
- Inspired by distributed and decentralized communication systems.
```

