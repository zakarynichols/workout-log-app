import { useEffect, useState } from 'react'

function App() {
  const [sessionResponse, setSessionResponse] = useState("")
  const [healthResponse, setHealthResponse] = useState("")

  const [error, setError] = useState<string | null>(null)

  const backendUrl = "/api"

  useEffect(() => {
    fetch(`${backendUrl}/sessions`)
      .then((res) => {
        if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
        return res.text();
      })
      .then((text) => {
        try {
          const json = JSON.parse(text);
          setSessionResponse(JSON.stringify(json));
        } catch (e) {
          setError("Response is not valid JSON");
          console.error("JSON parse error:", e);
        }
      })
      .catch((err) => {
        setError(err.message);
      });

          fetch(`${backendUrl}/health`)
      .then((res) => {
        if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
        return res.text();
      })
      .then((text) => {
        try {
          const json = JSON.parse(text);
          setHealthResponse(JSON.stringify(json));
        } catch (e) {
          setError("Response is not valid JSON");
          console.error("JSON parse error:", e);
        }
      })
      .catch((err) => {
        setError(err.message);
      });

  }, [backendUrl])

  return (
    <div>
      {error ? (
        <p style={{ color: "red" }}>Error: {error}</p>
      ) : (
        <div>
          <p>{healthResponse}</p>
          <br/>
          <p>{sessionResponse}</p>
        </div>
      )}
    </div>
  )
}

export default App
