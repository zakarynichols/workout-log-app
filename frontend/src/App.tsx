import { useEffect, useState } from 'react'

function App() {
  const [response, setResponse] = useState("")

  const backendUrl = import.meta.env.VITE_BACKEND_URL;

  console.log({backendUrl})

  useEffect(() => {
    fetch(`${backendUrl}`)
      .then((res) => res.json())
      .then((json) => {
        setResponse(JSON.stringify(json))
      })
  }, []);

  return (
    <div>
      <p>Server response: {response}</p>
    </div>
  )
}


export default App
