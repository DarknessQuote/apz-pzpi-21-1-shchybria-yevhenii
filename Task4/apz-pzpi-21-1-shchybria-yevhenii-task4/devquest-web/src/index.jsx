import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import { CssBaseline, StyledEngineProvider } from "@mui/material";
import { AuthProvider } from "./context/AuthContext";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
	<StyledEngineProvider injectFirst>
		<AuthProvider>
			<CssBaseline />
			<App />
		</AuthProvider>
	</StyledEngineProvider>
);
