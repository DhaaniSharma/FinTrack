import LoginForm from "./LoginForm"
import ParticleUniverse from "../../components/animations/ParticleUniverse"
import "./login.css"

function LoginPage() {
  return (
    <div className="login-container">

      <div className="left-panel">
        <ParticleUniverse />

        <div className="branding">
          <div className="logo">🛡</div>
          <h1 className="brand-name">FinShield</h1>
          <p className="tagline">Secure. Smart. Financial Control.</p>
          <p className="subtext">
            AI-driven personal finance intelligence
          </p>
        </div>
      </div>

      <div className="right-panel">
        <LoginForm />
      </div>

    </div>
  )
}

export default LoginPage
