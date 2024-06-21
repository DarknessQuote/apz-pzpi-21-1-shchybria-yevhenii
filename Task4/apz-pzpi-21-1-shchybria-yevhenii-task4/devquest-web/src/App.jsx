import { RouterProvider, createBrowserRouter } from "react-router-dom";
import AuthPage from "./routes/AuthPage";
import HomePage from "./routes/HomePage";
import RootLayout from "./routes/RootLayout";
import CompaniesPage from "./routes/CompaniesPage";

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
				{
					path: "/companies",
					element: <CompaniesPage />
				}
			],
		},
	]);

	return <RouterProvider router={router} />;
};

export default App;
