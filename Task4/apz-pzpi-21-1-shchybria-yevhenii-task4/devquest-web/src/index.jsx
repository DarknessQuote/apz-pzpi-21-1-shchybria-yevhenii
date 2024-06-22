import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import {
	CssBaseline,
	StyledEngineProvider,
	ThemeProvider,
	createTheme,
} from "@mui/material";
import { AuthProvider } from "./context/AuthContext";
import "./adapting/i18n";

const rootElement = document.getElementById("root");
const root = ReactDOM.createRoot(rootElement);

const theme = createTheme({
	components: {
		MuiPopover: {
			defaultProps: {
				container: rootElement,
			},
		},
		MuiPopper: {
			defaultProps: {
				container: rootElement,
			},
		},
		MuiDialog: {
			defaultProps: {
				container: rootElement,
			},
		},
	},
});

root.render(
	<StyledEngineProvider injectFirst>
		<ThemeProvider theme={theme}>
			<AuthProvider>
				<CssBaseline />
				<App />
			</AuthProvider>
		</ThemeProvider>
	</StyledEngineProvider>
);
