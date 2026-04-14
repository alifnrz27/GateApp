import React, { useState } from "react";

export default function Profile() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const save = () => {
    console.log({ username, password });
    alert("Profile updated (dummy)");
  };

  return (
    <div>
      <h1>Profile</h1>

      <input
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />

      <br /><br />

      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      <br /><br />

      <button onClick={save}>Save</button>
    </div>
  );
}