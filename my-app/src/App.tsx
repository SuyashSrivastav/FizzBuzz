import React, { useState } from "react";
import "./App.css";

type FizzBuzzResponse = {
  value: string;
};
console.log("Entering index file...");

const App: React.FC = () => {
  console.log("Entering index func...");
  const [count, setCount] = useState(0);
  const [message, setMessage] = useState("");

  const handleButtonClick = async () => {
    try {
      setCount((count) => count + 1);
      const response = await fetch(`http://localhost:3000/fizzbuzz`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ "count": count+1 }),
      });
      console.log("response here", response)
      if (!response.ok) {
        throw new Error("Something went wrong");
      }
      const json = (await response.json()) as FizzBuzzResponse;
   
      setMessage(json.value);
    } catch (e) {
      console.error(e);
      setMessage("Error fetching data. Please try again later.");
    }
  };

  return (
    <div className="container">
      <div className="centered">
        <div className="count-container" style={{ textAlign: "center" }}>
          <p className="count-label">Your count</p>
          <p className="count">{count}</p>
        </div>
        <button className="blue-button" onClick={handleButtonClick}>
          Push me!
    </button>
        <p className="button-text">{message}</p>
      </div>
    </div>

  );
};

export default App;
