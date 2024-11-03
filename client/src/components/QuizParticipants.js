import React, { useEffect, useState } from "react";
import Table from 'react-bootstrap/Table';

function QuizParticipants() {
  const [quizCode, setQuizCode] = useState("");
  const [ws, setWs] = useState(null);
  const [isConnected, setIsConnected] = useState(false);
  const [participants, setParticipants] = useState([]);
  const [userId, setUserId] = useState("");

  const handleJoinQuiz = () => {
    if (!quizCode || !userId) {
      alert("Please enter both Quiz Code and UserId.");
      return;
    }

    // TODO: use ws for local dev, will upgrade to wss in production
    const socket = new WebSocket(`ws://localhost:8080/api/v1/quizzes/${quizCode}`);

    socket.onopen = () => {
      console.log("Connected to WebSocket");
      setIsConnected(true);
      socket.send(
        JSON.stringify({ type: "join", data:{ user_id: userId }})
      );
    };

    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);

      if (message.type === "participants") {
        console.log(message.data);
        setParticipants(message.data.participants);
      }
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    socket.onclose = () => {
      console.log("WebSocket connection closed");
      setIsConnected(false);
    };

    setWs(socket);
  };

  useEffect(() => {
    return () => {
      if (ws) ws.close();
    };
  }, [ws]);

  return (
    <div className="quiz-participants-container">
      <h1 className="quiz-title">Join a Quiz</h1>

      {/* Quiz code and userId input */}
      {!isConnected && (
        <div className="join-quiz-form">
          <input
            type="text"
            placeholder="Enter Quiz Code"
            value={quizCode}
            onChange={(e) => setQuizCode(e.target.value)}
          />
          <input
            type="text"
            placeholder="Enter Your ID"
            value={userId}
            onChange={(e) => setUserId(e.target.value)}
          />
          <button onClick={handleJoinQuiz}>Join Quiz</button>
        </div>
      )}

      {/* Display participants if connected */}
      {isConnected && (
        <div className="participants-list">
          <h2>Participants in Quiz {quizCode}</h2>
          {participants.length > 0 ? (
            <Table striped bordered>
              <thead>
                <tr>
                  <th>UserID</th>
                </tr>
              </thead>
              <tbody>
                {participants.map((user) => (
                  <tr key={user.user_id}>
                    <td>{user.user_id}</td>
                  </tr>
                ))}
              </tbody>
            </Table>
          ) : (
            <p>No participants yet.</p>
          )}
        </div>
      )}
    </div>
  );
}

export default QuizParticipants;
