import ReactDOM from "react-dom";
import App from "./app/App";

// ========================================

const container = document.getElementById("root");
if (container === null) {
    throw new Error('Failed to find the root element');
}
const root = ReactDOM.createRoot(container);
root.render(<App />);
