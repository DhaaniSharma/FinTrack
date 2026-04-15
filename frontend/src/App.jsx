/*import LoginPage from "./pages/Login/LoginPage"

function App(){
  return <LoginPage />
}

export default App*/

import { BrowserRouter, Routes, Route } from "react-router-dom"

import LoginPage from "./pages/Login/LoginPage"
import RegisterPage from "./pages/Register/RegisterPage"
import DashboardPage from "./pages/Dashboard/DashboardPage";


function App() {
  return (

    <BrowserRouter>

      <Routes>

        <Route path="/" element={<LoginPage />} />

        <Route path="/register" element={<RegisterPage />} />

        <Route path="/dashboard" element={<DashboardPage />} />

      </Routes>

    </BrowserRouter>

  )
}

export default App