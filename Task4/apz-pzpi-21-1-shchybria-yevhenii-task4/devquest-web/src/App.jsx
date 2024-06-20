import { RouterProvider, createBrowserRouter } from "react-router-dom";
import AuthPage from "./routes/AuthPage";
import HomePage from "./routes/HomePage";
import RootLayout from "./routes/RootLayout";

const App = () => {
	const router = createBrowserRouter([
		{
			path: "/",
			element: <RootLayout />,
			children: [
				{
					index: true,
					element: <HomePage />,
				},
				{
					path: "/auth",
					element: <AuthPage />,
				},
			],
		},
	]);

	return <RouterProvider router={router} />;
};

export default App;
