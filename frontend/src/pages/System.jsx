import React, { useEffect, useState } from "react";

import { GetNetwork, SetNetwork } from "../../wailsjs/go/main/App";
import InputField from "../components/InputField";

export default function System() {
  const [ip, setIp] = useState("");
  const [subnet, setSubnet] = useState("");
  const [gateway, setGateway] = useState("");

  // ambil data awal
  useEffect(() => {
    loadNetwork();
  }, []);

  const loadNetwork = async () => {
    try {
      const res = await GetNetwork();

      setIp(res.ip?.trim());
      setSubnet(res.subnet?.trim());
      setGateway(res.gateway?.trim());
    } catch (err) {
      console.error(err);
    }
  };

  const save = async () => {
    try {
      const res = await SetNetwork(ip, subnet, gateway);
      alert(res);
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <h1>System Network</h1>

      <InputField
        label="IP Address"
        value={ip}
        onChange={setIp}
        placeholder="192.168.1.10"
        required
      />

      <InputField
        label="Subnet"
        value={subnet}
        onChange={setSubnet}
        placeholder="255.255.255.0"
      />

      <InputField
        label="Gateway"
        value={gateway}
        onChange={setGateway}
        placeholder="192.168.1.1"
      />

      <br />

      <button onClick={save}>
        Save Network
      </button>
    </div>
  );
}