import axios from "axios";

export async function getStatus() {
  const response = await axios.get("http://localhost:8080/status");
  return response.data.status;
}
