import React, { useState } from "react";

// nanti ganti dengan OpenGate dari Go
import { Hello } from "../../wailsjs/go/main/App";

export default function Gate() {
  const [gateId, setGateId] = useState("");

  const openGate = async () => {
    try {
      // nanti ganti ke OpenGate(gateId)
      const res = await Hello(gateId);
      alert(res);
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <h1>Gate Control</h1>

      <input
        placeholder="Gate UUID"
        value={gateId}
        onChange={(e) => setGateId(e.target.value)}
      />

      <br /><br />

      <button onClick={openGate}>
        Open Gate
      </button>
    </div>
  );
}