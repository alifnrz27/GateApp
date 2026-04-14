import React, { useEffect, useState } from "react";

// Wails event
import { EventsOn } from "../../wailsjs/runtime/runtime";

export default function Dashboard() {
  const [logs, setLogs] = useState([]);

  useEffect(() => {

    // listen event dari Go
    EventsOn("gate-log", (data) => {
      console.log("EVENT:", data);

      setLogs((prev) => [
        {
          time: new Date().toLocaleString(),
          ...data,
        },
        ...prev,
      ]);
    });

  }, []);

  return (
    <div>
      <h1>Dashboard (Realtime)</h1>

      <table border="1" cellPadding="10">
        <thead>
          <tr>
            <th color="black">Time</th>
            <th color="black">Username</th>
            <th color="black">Gate ID</th>
            <th color="black">Trigger</th>
            <th color="black">Raw</th>
          </tr>
        </thead>

        <tbody>
          {logs.map((log, i) => (
            <tr key={i}>
              <td color="black">{log.time}</td>
              <td color="black">{log.username}</td>
              <td color="black">{log.id_gate}</td>
              <td color="black">{log.trigger}</td>
              <td color="black"><pre>{JSON.stringify(log, null, 2)}</pre></td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}