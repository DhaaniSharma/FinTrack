import DashboardForm from "./DashboardForm"
import "./dashboard.css"
import MLGraph from "../../components/MLGraph"

function DashboardPage() {
  return (
    <div className="dashboard-page">
      <DashboardForm />

      <div style={{marginTop: "30px" }}>
        <MLGraph />
      </div>
    </div>
  )
}

export default DashboardPage