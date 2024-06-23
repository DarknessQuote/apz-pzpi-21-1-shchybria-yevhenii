import { RouterProvider, createBrowserRouter } from "react-router-dom";
import AuthPage from "./routes/AuthPage";
import HomePage from "./routes/HomePage";
import RootLayout from "./routes/RootLayout";
import CompaniesPage from "./routes/CompaniesPage";
import ProjectsPage from "./routes/ProjectsPage";
import ProjectPage from "./routes/ProjectPage";

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
					element: <CompaniesPage />,
				},
				{
					path: "/projects",
					element: <ProjectsPage />,
				},
				{
					path: "/projects/:id",
					element: <ProjectPage />,
				},
			],
		},
	]);

	return <RouterProvider router={router} />;
};

export default App;
