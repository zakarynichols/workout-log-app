import { useEffect, useState } from 'react'

function App() {
  const [response, setResponse] = useState("")
  const [error, setError] = useState<string | null>(null)

  const backendUrl = import.meta.env.VITE_BACKEND_URL

  console.log({ backendUrl })

  useEffect(() => {
    if (!backendUrl) {
      setError("Backend URL is not defined")
      return
    }
  fetch(backendUrl)
    .then((res) => {
      if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
      return res.text();  // get raw text first
    })
    .then((text) => {
      console.log("Raw response text:", text);
      // try parsing JSON now:
      try {
        const json = JSON.parse(text);
        setResponse(JSON.stringify(json, null, 2));
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
        <pre>Server response: {response}</pre>
      )}
    </div>
  )
}

export default App
