import React, { useState } from "react";

function App() {
  const [name, setName] = useState("");
  const [greeting, setGreeting] = useState("");

  const handleNameChange = (e) => {
    setName(e.target.value);
  };

  const fetchGreeting = async () => {
    try {
      // Replace <backend-vm-ip> with your backend VM's IP address
      const response = await fetch(`http://localhost:8080/greet/${name}`);
      const data = await response.text();
      setGreeting(data);
    } catch (error) {
      console.error("Error fetching greeting:", error);
      setGreeting("Failed to fetch greeting.");
    }
  };

  return (
    <div className="App">
      <h1>React Greeting App</h1>
      <input
        type="text"
        value={name}
        onChange={handleNameChange}
        placeholder="Enter your name"
      />
      <button onClick={fetchGreeting}>Greet</button>
      {greeting && <p>{greeting}</p>}
    </div>
  );
}

export default App;
