# miro-whiteboard-grpc
A simple, online collaborative whiteboard platform that allows users to sketch, brainstorm, and share ideas without requiring sign-up
## Services
**1. User Service**
- Functionality: User authentication and management.
- Endpoints: Register, Login, GetUser.
- Tech Stack: Go, gRPC, PostgreSQL.

**2. Board Service**
- Functionality: Board creation, deletion, retrieval.
- Endpoints: CreateBoard, GetBoard, DeleteBoard.
- Tech Stack: Go, gRPC, MongoDB.

**3. Collaboration Service**
- Functionality: Real-time collaboration management.
- Endpoints: StartCollaboration, GetCollaborators.
- Tech Stack: Go, gRPC, Redis.

**4. Drawing Service**
- Functionality: Handle drawing actions and updates.
- Endpoints: AddDrawing, GetDrawings.
- Tech Stack: Go, gRPC, Redis.

**5. Gateway Service**
- Functionality: API gateway to route requests to appropriate services.
- Tech Stack: Go, gRPC, NGINX, WebSockets.
