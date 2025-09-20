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
        if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`)
        return res.json()
      })
      .then((json) => {
        setResponse(JSON.stringify(json, null, 2))
      })
      .catch((err) => {
        setError(err.message)
      })
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
