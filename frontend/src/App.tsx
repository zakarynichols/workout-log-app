import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Home from "./pages/Home";
import Sessions from "./pages/Sessions";
import SessionDetails from "./pages/SessionDetailPage";

export default function App() {
  return (
    <Router>
        <nav>
        <Link to="/">Home</Link>
        <Link to="/sessions">Sessions</Link>
      </nav>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/sessions" element={<Sessions />} />
        <Route path="/sessions/:id" element={<SessionDetails />} />
      </Routes>
    </Router>
  );
}