import { RouterProvider, createBrowserRouter } from "react-router-dom";
import AuthPage from "./routes/AuthPage";
import HomePage from "./routes/HomePage";

const App = () => {
	const router = createBrowserRouter([
		{
			path: "/",
			element: <HomePage />,
		},
		{
			path: "/auth",
			element: <AuthPage />,
		},
	]);

	return <RouterProvider router={router} />;
};

export default App;
